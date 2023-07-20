package tests

import (
	"context"
	"math"
	"sort"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/prysmaticlabs/prysm/crypto/bls"

	"github.com/bnb-chain/greenfield/e2e/core"
	"github.com/bnb-chain/greenfield/sdk/keys"
	"github.com/bnb-chain/greenfield/sdk/types"
	storagetestutils "github.com/bnb-chain/greenfield/testutil/storage"
	paymenttypes "github.com/bnb-chain/greenfield/x/payment/types"
	sptypes "github.com/bnb-chain/greenfield/x/sp/types"
	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"
	virtualgrouptypes "github.com/bnb-chain/greenfield/x/virtualgroup/types"
)

func (s *PaymentTestSuite) copyObject(user keys.KeyManager, sp *core.StorageProvider, bucketName, objectName, dstBucketName, dstObjectName string) (*storagetypes.ObjectInfo, error) {
	msgCopyObject := storagetypes.NewMsgCopyObject(user.GetAddr(), bucketName, dstBucketName, objectName, dstObjectName, math.MaxUint, nil)
	msgCopyObject.DstPrimarySpApproval.Sig, _ = sp.ApprovalKey.Sign(msgCopyObject.GetApprovalBytes())

	s.SendTxBlock(user, msgCopyObject)
	// HeadObject
	queryHeadObjectRequest := storagetypes.QueryHeadObjectRequest{
		BucketName: dstBucketName,
		ObjectName: dstObjectName,
	}
	headObjectResponse, err := s.Client.HeadObject(context.Background(), &queryHeadObjectRequest)
	return headObjectResponse.ObjectInfo, err

}

func (s *PaymentTestSuite) updateBucket(user keys.KeyManager, bucketName string, paymentAddress string, chargedReadQuota uint64) (*storagetypes.BucketInfo, error) {

	msgUpdateBucketInfo := storagetypes.NewMsgUpdateBucketInfo(
		user.GetAddr(), bucketName, &chargedReadQuota, user.GetAddr(), storagetypes.VISIBILITY_TYPE_PRIVATE)
	if paymentAddress != "" {
		msgUpdateBucketInfo.PaymentAddress = paymentAddress
	}
	s.SendTxBlock(user, msgUpdateBucketInfo)

	queryHeadObjectRequest := storagetypes.QueryHeadBucketRequest{
		BucketName: bucketName,
	}
	headObjectResponse, err := s.Client.HeadBucket(context.Background(), &queryHeadObjectRequest)
	return headObjectResponse.BucketInfo, err

}

// TestStorageBill_CopyObject_WithoutPriceChange
func (s *PaymentTestSuite) TestStorageBill_CopyObject_WithoutPriceChange() {
	var err error
	ctx := context.Background()
	sp := s.PickStorageProvider()
	gvg, found := sp.GetFirstGlobalVirtualGroup()
	s.Require().True(found)
	queryFamilyResponse, err := s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: gvg.FamilyId,
	})
	s.Require().NoError(err)
	family := queryFamilyResponse.GlobalVirtualGroupFamily
	user0 := s.GenAndChargeAccounts(1, 1000000)[0]

	streamAddresses := []string{
		user0.GetAddr().String(),
		family.VirtualPaymentAddress,
		gvg.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}

	paymentParams, err := s.Client.PaymentQueryClient.Params(ctx, &paymenttypes.QueryParamsRequest{})
	s.T().Logf("paymentParams %s, err: %v", paymentParams, err)
	s.Require().NoError(err)

	bucketName := s.createBucket(sp, user0, 0)

	// create object with none zero payload size
	streamRecordsBefore := s.getStreamRecords(streamAddresses)
	_, _, objectName, objectId, checksums, payloadSize := s.createObject(user0, bucketName, false)

	// assertions
	streamRecordsAfter := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	lockFee := s.calculateLockFee(sp, bucketName, objectName, payloadSize)
	s.Require().Equal(streamRecordsAfter.User.LockBalance.Sub(streamRecordsBefore.User.LockBalance), lockFee)
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate).Int64(), int64(0))

	// case: seal object without price change
	s.sealObject(sp, gvg, bucketName, objectName, objectId, checksums)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.LockBalance, sdkmath.ZeroInt())
	gvgFamilyRate, gvgRate, taxRate, userTotalRate := s.calculateStorageRates(sp, bucketName, objectName, payloadSize, 0)
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), gvgFamilyRate)
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate), gvgRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	distBucketName := bucketName
	distObjectName := storagetestutils.GenRandomObjectName()

	objectIfo, err := s.copyObject(user0, sp, bucketName, objectName, distBucketName, distObjectName)
	s.Require().NoError(err)
	s.sealObject(sp, gvg, distBucketName, distObjectName, objectIfo.Id, objectIfo.Checksums)
	// assertions
	streamRecordsAfterCopy := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfterCopy.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfterCopy.User.LockBalance, sdkmath.ZeroInt())
	//gvgFamilyRate1, gvgRate1, taxRate1, userTotalRate1 := s.calculateStorageRates(sp,distBucketName, distObjectName, payloadSize)
	s.Require().Equal(streamRecordsAfterCopy.User.NetflowRate.Sub(streamRecordsAfter.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfterCopy.GVGFamily.NetflowRate.Sub(streamRecordsAfter.GVGFamily.NetflowRate), gvgFamilyRate)
	s.Require().Equal(streamRecordsAfterCopy.GVG.NetflowRate.Sub(streamRecordsAfter.GVG.NetflowRate), gvgRate)
	s.Require().Equal(streamRecordsAfterCopy.Tax.NetflowRate.Sub(streamRecordsAfter.Tax.NetflowRate), taxRate)
}

// TestStorageBill_CopyObject_WithoutPriceChange
func (s *PaymentTestSuite) TestStorageBill_CopyObject_WithPriceChange() {
	var err error
	ctx := context.Background()
	sp := s.PickStorageProvider()
	gvg, found := sp.GetFirstGlobalVirtualGroup()
	s.Require().True(found)
	queryFamilyResponse, err := s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: gvg.FamilyId,
	})
	s.Require().NoError(err)
	family := queryFamilyResponse.GlobalVirtualGroupFamily
	user0 := s.GenAndChargeAccounts(1, 1000000)[0]

	streamAddresses := []string{
		user0.GetAddr().String(),
		family.VirtualPaymentAddress,
		gvg.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}

	paymentParams, err := s.Client.PaymentQueryClient.Params(ctx, &paymenttypes.QueryParamsRequest{})
	s.T().Logf("paymentParams %s, err: %v", paymentParams, err)
	s.Require().NoError(err)

	bucketName := s.createBucket(sp, user0, 0)

	// create object with none zero payload size
	streamRecordsBefore := s.getStreamRecords(streamAddresses)
	_, _, objectName, objectId, checksums, payloadSize := s.createObject(user0, bucketName, false)

	// assertions
	streamRecordsAfter := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	lockFee := s.calculateLockFee(sp, bucketName, objectName, payloadSize)
	s.Require().Equal(streamRecordsAfter.User.LockBalance.Sub(streamRecordsBefore.User.LockBalance), lockFee)
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate).Int64(), int64(0))

	// case: seal object without price change
	s.sealObject(sp, gvg, bucketName, objectName, objectId, checksums)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.LockBalance, sdkmath.ZeroInt())
	gvgFamilyRate, gvgRate, taxRate, userTotalRate := s.calculateStorageRates(sp, bucketName, objectName, payloadSize, 0)
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), gvgFamilyRate)
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate), gvgRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	priceRes, err := s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: 0,
	})
	s.Require().NoError(err)
	// update new price
	msgUpdatePrice := &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     priceRes.SpStoragePrice.ReadPrice,
		FreeReadQuota: priceRes.SpStoragePrice.FreeReadQuota,
		StorePrice:    priceRes.SpStoragePrice.StorePrice.MulInt64(1000),
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)

	distBucketName := s.createBucket(sp, user0, 0)
	distObjectName := storagetestutils.GenRandomObjectName()
	objectIfo, err := s.copyObject(user0, sp, bucketName, objectName, distBucketName, distObjectName)
	s.Require().NoError(err)
	s.sealObject(sp, gvg, distBucketName, distObjectName, objectIfo.Id, objectIfo.Checksums)
	// assertions
	streamRecordsAfterCopy := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfterCopy.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfterCopy.User.LockBalance, sdkmath.ZeroInt())
	gvgFamilyRate1, gvgRate1, taxRate1, userTotalRate1 := s.calculateStorageRates(sp, distBucketName, distObjectName, payloadSize, 0)
	s.Require().Equal(streamRecordsAfterCopy.GVGFamily.NetflowRate.Sub(streamRecordsAfter.GVGFamily.NetflowRate), gvgFamilyRate1)
	s.Require().Equal(streamRecordsAfterCopy.GVG.NetflowRate.Sub(streamRecordsAfter.GVG.NetflowRate), gvgRate1)
	s.Require().Equal(streamRecordsAfterCopy.Tax.NetflowRate.Sub(streamRecordsAfter.Tax.NetflowRate), taxRate1)
	s.Require().Equal(streamRecordsAfterCopy.User.NetflowRate.Sub(streamRecordsAfter.User.NetflowRate).BigInt().String(), userTotalRate1.Neg().BigInt().String())

}

// TestStorageBill_UpdateBucketQuota
func (s *PaymentTestSuite) TestStorageBill_UpdateBucketQuota() {
	var err error
	ctx := context.Background()
	sp := s.PickStorageProvider()
	// recover price
	defer s.RecoverSPPrice(sp)
	gvg, found := sp.GetFirstGlobalVirtualGroup()
	s.Require().True(found)
	queryFamilyResponse, err := s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: gvg.FamilyId,
	})
	s.Require().NoError(err)
	family := queryFamilyResponse.GlobalVirtualGroupFamily
	user := s.GenAndChargeAccounts(1, 10)[0]

	streamAddresses := []string{
		user.GetAddr().String(),
		family.VirtualPaymentAddress,
		gvg.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}
	streamRecordsBefore := s.getStreamRecords(streamAddresses)

	paymentParams, err := s.Client.PaymentQueryClient.Params(ctx, &paymenttypes.QueryParamsRequest{})
	s.T().Logf("paymentParams %s, err: %v", paymentParams, err)
	s.Require().NoError(err)

	// case: create bucket with zero read quota
	bucketName := s.createBucket(sp, user, 0)

	// bucket created
	queryHeadBucketRequest := storagetypes.QueryHeadBucketRequest{
		BucketName: bucketName,
	}
	_, err = s.Client.HeadBucket(ctx, &queryHeadBucketRequest)
	s.Require().NoError(err)

	// assertions
	streamRecordsAfter := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate, streamRecordsBefore.User.NetflowRate)
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate, streamRecordsBefore.GVGFamily.NetflowRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate, streamRecordsBefore.Tax.NetflowRate)

	readQuota := uint64(1024 * 1024 * 100)
	// case: update bucket read quota
	bucketInfo, err := s.updateBucket(user, bucketName, "", readQuota)
	s.Require().NoError(err)

	// check price and rate calculation
	queryGetSpStoragePriceByTimeResp, err := s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.T().Logf("queryGetSpStoragePriceByTimeResp %s, err: %v", queryGetSpStoragePriceByTimeResp, err)
	s.Require().NoError(err)

	readPrice := queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice
	readChargeRate := readPrice.MulInt(sdk.NewIntFromUint64(bucketInfo.ChargedReadQuota)).TruncateInt()
	s.T().Logf("readPrice: %s, readChargeRate: %s", readPrice, readChargeRate)
	taxRate := paymentParams.Params.VersionedParams.ValidatorTaxRate.MulInt(readChargeRate).TruncateInt()
	userTotalRate := readChargeRate.Add(taxRate)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), readChargeRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	expectedOutFlows := []paymenttypes.OutFlow{
		{ToAddress: family.VirtualPaymentAddress, Rate: readChargeRate},
		{ToAddress: paymenttypes.ValidatorTaxPoolAddress.String(), Rate: taxRate},
	}
	userOutFlowsResponse, err := s.Client.OutFlows(ctx, &paymenttypes.QueryOutFlowsRequest{Account: user.GetAddr().String()})
	s.Require().NoError(err)
	sort.Slice(userOutFlowsResponse.OutFlows, func(i, j int) bool {
		return userOutFlowsResponse.OutFlows[i].ToAddress < userOutFlowsResponse.OutFlows[j].ToAddress
	})
	sort.Slice(expectedOutFlows, func(i, j int) bool {
		return expectedOutFlows[i].ToAddress < expectedOutFlows[j].ToAddress
	})
	s.Require().Equal(expectedOutFlows, userOutFlowsResponse.OutFlows)
	// update new price
	msgUpdatePrice := &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice.MulInt64(100),
		FreeReadQuota: queryGetSpStoragePriceByTimeResp.SpStoragePrice.FreeReadQuota,
		StorePrice:    queryGetSpStoragePriceByTimeResp.SpStoragePrice.StorePrice,
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)

	// case: update bucket read quota
	bucketInfo, err = s.updateBucket(user, bucketName, "", readQuota*2)
	s.Require().NoError(err)

	// check price and rate calculation
	queryGetSpStoragePriceByTimeResp, err = s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.T().Logf("queryGetSpStoragePriceByTimeResp %s, err: %v", queryGetSpStoragePriceByTimeResp, err)
	s.Require().NoError(err)

	readPrice = queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice
	readChargeRate = readPrice.MulInt(sdk.NewIntFromUint64(bucketInfo.ChargedReadQuota)).TruncateInt()
	s.T().Logf("readPrice: %s, readChargeRate: %s", readPrice, readChargeRate)
	taxRate = paymentParams.Params.VersionedParams.ValidatorTaxRate.MulInt(readChargeRate).TruncateInt()
	userTotalRate = readChargeRate.Add(taxRate)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), readChargeRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	expectedOutFlows = []paymenttypes.OutFlow{
		{ToAddress: family.VirtualPaymentAddress, Rate: readChargeRate},
		{ToAddress: paymenttypes.ValidatorTaxPoolAddress.String(), Rate: taxRate},
	}
	userOutFlowsResponse, err = s.Client.OutFlows(ctx, &paymenttypes.QueryOutFlowsRequest{Account: user.GetAddr().String()})
	s.Require().NoError(err)
	sort.Slice(userOutFlowsResponse.OutFlows, func(i, j int) bool {
		return userOutFlowsResponse.OutFlows[i].ToAddress < userOutFlowsResponse.OutFlows[j].ToAddress
	})
	sort.Slice(expectedOutFlows, func(i, j int) bool {
		return expectedOutFlows[i].ToAddress < expectedOutFlows[j].ToAddress
	})
	s.Require().Equal(expectedOutFlows, userOutFlowsResponse.OutFlows)
	// set big read price
	msgUpdatePrice = &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice.MulInt64(1024 * 1024 * 1024),
		FreeReadQuota: queryGetSpStoragePriceByTimeResp.SpStoragePrice.FreeReadQuota,
		StorePrice:    queryGetSpStoragePriceByTimeResp.SpStoragePrice.StorePrice,
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)

	chargedReadQuota := readQuota * 1024 * 1024
	msgUpdateBucketInfo := storagetypes.NewMsgUpdateBucketInfo(
		user.GetAddr(), bucketName, &chargedReadQuota, user.GetAddr(), storagetypes.VISIBILITY_TYPE_PRIVATE)

	s.SendTxBlockWithExpectErrorString(msgUpdateBucketInfo, user, "apply user flows list failed")

}

// TestStorageBill_UpdatePaymentAddress
func (s *PaymentTestSuite) TestStorageBill_UpdatePaymentAddress() {
	var err error
	ctx := context.Background()
	sp := s.PickStorageProvider()
	defer s.RecoverSPPrice(sp)
	gvg, found := sp.GetFirstGlobalVirtualGroup()
	s.Require().True(found)
	queryFamilyResponse, err := s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: gvg.FamilyId,
	})
	s.Require().NoError(err)
	family := queryFamilyResponse.GlobalVirtualGroupFamily
	user := s.GenAndChargeAccounts(1, 100)[0]

	paymentAccountAddr := s.CreatePaymentAccount(user, 1, 17)
	paymentAcc := sdk.MustAccAddressFromHex(paymentAccountAddr)
	streamAddresses := []string{
		user.GetAddr().String(),
		family.VirtualPaymentAddress,
		gvg.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}
	streamRecordsBefore := s.getStreamRecords(streamAddresses)

	paymentParams, err := s.Client.PaymentQueryClient.Params(ctx, &paymenttypes.QueryParamsRequest{})
	s.T().Logf("paymentParams %s, err: %v", paymentParams, err)
	s.Require().NoError(err)

	// case: create bucket with zero read quota
	bucketName := s.createBucket(sp, user, 0)

	// bucket created
	queryHeadBucketRequest := storagetypes.QueryHeadBucketRequest{
		BucketName: bucketName,
	}
	_, err = s.Client.HeadBucket(ctx, &queryHeadBucketRequest)
	s.Require().NoError(err)

	// assertions
	streamRecordsAfter := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate, streamRecordsBefore.User.NetflowRate)
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate, streamRecordsBefore.GVGFamily.NetflowRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate, streamRecordsBefore.Tax.NetflowRate)

	readQuota := uint64(1024 * 100)
	// case: update bucket read quota
	bucketInfo, err := s.updateBucket(user, bucketName, "", readQuota)
	s.Require().NoError(err)

	// check price and rate calculation
	queryGetSpStoragePriceByTimeResp, err := s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.T().Logf("queryGetSpStoragePriceByTimeResp %s, err: %v", queryGetSpStoragePriceByTimeResp, err)
	s.Require().NoError(err)

	readPrice := queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice
	readChargeRate := readPrice.MulInt(sdk.NewIntFromUint64(bucketInfo.ChargedReadQuota)).TruncateInt()
	s.T().Logf("readPrice: %s, readChargeRate: %s", readPrice, readChargeRate)
	taxRate := paymentParams.Params.VersionedParams.ValidatorTaxRate.MulInt(readChargeRate).TruncateInt()
	userTotalRate := readChargeRate.Add(taxRate)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), readChargeRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	expectedOutFlows := []paymenttypes.OutFlow{
		{ToAddress: family.VirtualPaymentAddress, Rate: readChargeRate},
		{ToAddress: paymenttypes.ValidatorTaxPoolAddress.String(), Rate: taxRate},
	}
	userOutFlowsResponse, err := s.Client.OutFlows(ctx, &paymenttypes.QueryOutFlowsRequest{Account: user.GetAddr().String()})
	s.Require().NoError(err)
	sort.Slice(userOutFlowsResponse.OutFlows, func(i, j int) bool {
		return userOutFlowsResponse.OutFlows[i].ToAddress < userOutFlowsResponse.OutFlows[j].ToAddress
	})
	sort.Slice(expectedOutFlows, func(i, j int) bool {
		return expectedOutFlows[i].ToAddress < expectedOutFlows[j].ToAddress
	})
	s.Require().Equal(expectedOutFlows, userOutFlowsResponse.OutFlows)
	// update new price
	msgUpdatePrice := &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice.MulInt64(100),
		FreeReadQuota: queryGetSpStoragePriceByTimeResp.SpStoragePrice.FreeReadQuota,
		StorePrice:    queryGetSpStoragePriceByTimeResp.SpStoragePrice.StorePrice,
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)

	// case: update bucket paymentAccountAddr
	bucketInfo, err = s.updateBucket(user, bucketName, paymentAccountAddr, readQuota)
	s.Require().NoError(err)

	// check price and rate calculation
	queryGetSpStoragePriceByTimeResp, err = s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.T().Logf("queryGetSpStoragePriceByTimeResp %s, err: %v", queryGetSpStoragePriceByTimeResp, err)
	s.Require().NoError(err)

	readPrice = queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice
	readChargeRate = readPrice.MulInt(sdk.NewIntFromUint64(bucketInfo.ChargedReadQuota)).TruncateInt()
	s.T().Logf("readPrice: %s, readChargeRate: %s", readPrice, readChargeRate)
	taxRate = paymentParams.Params.VersionedParams.ValidatorTaxRate.MulInt(readChargeRate).TruncateInt()
	userTotalRate = readChargeRate.Add(taxRate)

	// assertions
	streamAddresses[0] = paymentAccountAddr
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), readChargeRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)

	expectedOutFlows = []paymenttypes.OutFlow{
		{ToAddress: family.VirtualPaymentAddress, Rate: readChargeRate},
		{ToAddress: paymenttypes.ValidatorTaxPoolAddress.String(), Rate: taxRate},
	}
	userOutFlowsResponse, err = s.Client.OutFlows(ctx, &paymenttypes.QueryOutFlowsRequest{Account: paymentAccountAddr})
	s.Require().NoError(err)
	sort.Slice(userOutFlowsResponse.OutFlows, func(i, j int) bool {
		return userOutFlowsResponse.OutFlows[i].ToAddress < userOutFlowsResponse.OutFlows[j].ToAddress
	})
	sort.Slice(expectedOutFlows, func(i, j int) bool {
		return expectedOutFlows[i].ToAddress < expectedOutFlows[j].ToAddress
	})
	s.Require().Equal(expectedOutFlows, userOutFlowsResponse.OutFlows)

	// set big read price
	msgUpdatePrice = &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice.MulInt64(1024 * 1024 * 1024),
		FreeReadQuota: queryGetSpStoragePriceByTimeResp.SpStoragePrice.FreeReadQuota,
		StorePrice:    queryGetSpStoragePriceByTimeResp.SpStoragePrice.StorePrice,
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)

	chargedReadQuota := readQuota * 1024 * 1024 * 1024 * 1024
	msgUpdateBucketInfo := storagetypes.NewMsgUpdateBucketInfo(
		user.GetAddr(), bucketName, &chargedReadQuota, paymentAcc, storagetypes.VISIBILITY_TYPE_PRIVATE)
	s.SendTxBlockWithExpectErrorString(msgUpdateBucketInfo, user, "apply user flows list failed")
	// new payment account balance not enough
	paymentAccountAddr = s.CreatePaymentAccount(user, 1, 13)
	paymentAcc = sdk.MustAccAddressFromHex(paymentAccountAddr)
	msgUpdateBucketInfo = storagetypes.NewMsgUpdateBucketInfo(
		user.GetAddr(), bucketName, &readQuota, paymentAcc, storagetypes.VISIBILITY_TYPE_PRIVATE)

	s.SendTxBlockWithExpectErrorString(msgUpdateBucketInfo, user, "apply user flows list failed")

}

func (s *PaymentTestSuite) TestStorageBill_MigrationBucket() {
	var err error
	ctx := context.Background()
	primarySP := s.PickStorageProvider()
	s.RecoverSPPrice(primarySP)

	gvg, found := primarySP.GetFirstGlobalVirtualGroup()
	s.Require().True(found)
	queryFamilyResponse, err := s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: gvg.FamilyId,
	})
	s.Require().NoError(err)
	family := queryFamilyResponse.GlobalVirtualGroupFamily
	user0 := s.GenAndChargeAccounts(1, 1000000)[0]

	streamAddresses := []string{
		user0.GetAddr().String(),
		family.VirtualPaymentAddress,
		gvg.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}

	paymentParams, err := s.Client.PaymentQueryClient.Params(ctx, &paymenttypes.QueryParamsRequest{})
	s.T().Logf("paymentParams %s, err: %v", paymentParams, err)
	s.Require().NoError(err)

	bucketName := s.createBucket(primarySP, user0, 0)
	bucketInfo, err := s.Client.HeadBucket(context.Background(), &storagetypes.QueryHeadBucketRequest{
		BucketName: bucketName,
	})
	s.Require().NoError(err)

	// create object with none zero payload size
	streamRecordsBefore := s.getStreamRecords(streamAddresses)
	_, _, objectName, objectId, checksums, payloadSize := s.createObject(user0, bucketName, false)

	// assertions
	streamRecordsAfter := s.getStreamRecords(streamAddresses)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	lockFee := s.calculateLockFee(primarySP, bucketName, objectName, payloadSize)
	s.Require().Equal(streamRecordsAfter.User.LockBalance.Sub(streamRecordsBefore.User.LockBalance), lockFee)
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate).Int64(), int64(0))
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate).Int64(), int64(0))

	// case: seal object without price change
	s.sealObject(primarySP, gvg, bucketName, objectName, objectId, checksums)

	// assertions
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	gvgFamilyRate, gvgRate, taxRate, userTotalRate := s.calculateStorageRates(primarySP, bucketName, objectName, payloadSize, 0)
	s.T().Logf("gvgFamilyRate: %v, gvgRate: %v, taxRate: %v, userTotalRate: %v", gvgFamilyRate, gvgRate, taxRate, userTotalRate)
	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.LockBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Sub(streamRecordsBefore.User.NetflowRate), userTotalRate.Neg())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), gvgFamilyRate)
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate), gvgRate)
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate)
	taxRate0 := taxRate
	dstPrimarySP := s.CreateNewStorageProvider()

	s.RecoverSPPrice(dstPrimarySP)

	// new sp CreateGlobalVirtualGroup
	var secondarySPIDs []uint32
	var secondarySPs []*core.StorageProvider

	for _, ssp := range s.StorageProviders {
		if ssp.Info.Id != primarySP.Info.Id {
			secondarySPIDs = append(secondarySPIDs, ssp.Info.Id)
			secondarySPs = append(secondarySPs, ssp)
		}
		if len(secondarySPIDs) == 6 {
			break
		}
	}
	gvgID, _ := s.BaseSuite.CreateGlobalVirtualGroup(dstPrimarySP, 0, secondarySPIDs, 1)
	gvgResp, err := s.Client.VirtualGroupQueryClient.GlobalVirtualGroup(context.Background(), &virtualgrouptypes.QueryGlobalVirtualGroupRequest{
		GlobalVirtualGroupId: gvgID,
	})
	s.Require().NoError(err)
	dstGVG := gvgResp.GlobalVirtualGroup
	s.Require().True(found)

	queryFamilyResponse, err = s.Client.GlobalVirtualGroupFamily(ctx, &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: dstGVG.FamilyId,
	})
	s.Require().NoError(err)
	family = queryFamilyResponse.GlobalVirtualGroupFamily
	streamAddresses = []string{
		user0.GetAddr().String(),
		family.VirtualPaymentAddress,
		dstGVG.VirtualPaymentAddress,
		paymenttypes.ValidatorTaxPoolAddress.String(),
	}
	fundAddress := primarySP.FundingKey.GetAddr()
	queryBalanceRequest := banktypes.QueryBalanceRequest{Denom: s.Config.Denom, Address: fundAddress.String()}

	fundBalanceBefore, err := s.Client.BankQueryClient.Balance(context.Background(), &queryBalanceRequest)
	s.Require().NoError(err)

	streamRecordsBefore = s.getStreamRecords(streamAddresses)
	// construct the signatures
	var gvgMappings []*storagetypes.GVGMapping
	gvgMappings = append(gvgMappings, &storagetypes.GVGMapping{SrcGlobalVirtualGroupId: gvg.Id, DstGlobalVirtualGroupId: dstGVG.Id})
	for _, gvgMapping := range gvgMappings {
		migrationBucketSignHash := storagetypes.NewSecondarySpMigrationBucketSignDoc(s.GetChainID(), bucketInfo.BucketInfo.Id, dstPrimarySP.Info.Id, gvgMapping.SrcGlobalVirtualGroupId, gvgMapping.DstGlobalVirtualGroupId).GetBlsSignHash()
		secondarySigs := make([][]byte, 0)
		secondarySPBlsPubKeys := make([]bls.PublicKey, 0)
		for _, ssp := range secondarySPs {
			sig, err := core.BlsSignAndVerify(ssp, migrationBucketSignHash)
			s.Require().NoError(err)
			secondarySigs = append(secondarySigs, sig)
			pk, err := bls.PublicKeyFromBytes(ssp.BlsKey.PubKey().Bytes())
			s.Require().NoError(err)
			secondarySPBlsPubKeys = append(secondarySPBlsPubKeys, pk)
		}
		aggBlsSig, err := core.BlsAggregateAndVerify(secondarySPBlsPubKeys, migrationBucketSignHash, secondarySigs)
		s.Require().NoError(err)
		gvgMapping.SecondarySpBlsSignature = aggBlsSig
	}

	// send msgMigrationBucket
	msgMigrationBucket := storagetypes.NewMsgMigrateBucket(user0.GetAddr(), bucketName, dstPrimarySP.Info.Id)
	msgMigrationBucket.DstPrimarySpApproval.ExpiredHeight = math.MaxInt
	msgMigrationBucket.DstPrimarySpApproval.Sig, err = dstPrimarySP.ApprovalKey.Sign(msgMigrationBucket.GetApprovalBytes())
	s.SendTxBlock(user0, msgMigrationBucket)
	s.Require().NoError(err)

	// complete MigrationBucket
	msgCompleteMigrationBucket := storagetypes.NewMsgCompleteMigrateBucket(dstPrimarySP.OperatorKey.GetAddr(), bucketName, dstGVG.FamilyId, gvgMappings)
	s.SendTxBlock(dstPrimarySP.OperatorKey, msgCompleteMigrationBucket)
	streamRecordsAfter = s.getStreamRecords(streamAddresses)
	fundBalanceAfter, err := s.Client.BankQueryClient.Balance(context.Background(), &queryBalanceRequest)
	s.Require().NoError(err)
	s.T().Logf("fundBalanceBefore: %v, fundBalanceAfter: %v, diff: %v", fundBalanceBefore, fundBalanceAfter, fundBalanceAfter.Balance.Amount.Sub(fundBalanceBefore.Balance.Amount))
	s.Require().True(fundBalanceAfter.Balance.Amount.Sub(fundBalanceBefore.Balance.Amount).GT(sdkmath.NewInt(0)), "migrate sp fund address need settle")
	gvgFamilyRate, gvgRate, taxRate, userTotalRate = s.calculateStorageRates(dstPrimarySP, bucketName, objectName, payloadSize, time.Now().Unix())
	s.T().Logf("gvgFamilyRate: %v, gvgRate: %v, taxRate: %v, userTotalRate: %v", gvgFamilyRate, gvgRate, taxRate, userTotalRate)
	s.T().Logf("NetflowRate: %v, userTotalRate: %v, actual taxRate diff: %v, expect taxRate diff: %v", streamRecordsAfter.User.NetflowRate.Neg(), userTotalRate.Neg(), streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate.Sub(taxRate0))

	s.Require().Equal(streamRecordsAfter.User.StaticBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.User.LockBalance, sdkmath.ZeroInt())
	s.Require().Equal(streamRecordsAfter.GVGFamily.NetflowRate.Sub(streamRecordsBefore.GVGFamily.NetflowRate), gvgFamilyRate)
	s.Require().Equal(streamRecordsAfter.GVG.NetflowRate.Sub(streamRecordsBefore.GVG.NetflowRate), gvgRate)
	// tax rate diff
	s.Require().Equal(streamRecordsAfter.Tax.NetflowRate.Sub(streamRecordsBefore.Tax.NetflowRate), taxRate.Sub(taxRate0))
	s.Require().Equal(streamRecordsAfter.User.NetflowRate.Neg(), userTotalRate.Abs())

}

func (s *PaymentTestSuite) RecoverSPPrice(sp *core.StorageProvider) {
	ctx := context.Background()

	queryGetSpStoragePriceByTimeResp, err := s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.Require().NoError(err)
	recoverReadPrice, err := sdk.NewDecFromStr("100.0087")
	s.Require().NoError(err)

	msgUpdatePrice := &sptypes.MsgUpdateSpStoragePrice{
		SpAddress:     sp.OperatorKey.GetAddr().String(),
		ReadPrice:     recoverReadPrice,
		FreeReadQuota: queryGetSpStoragePriceByTimeResp.SpStoragePrice.FreeReadQuota,
		StorePrice:    queryGetSpStoragePriceByTimeResp.SpStoragePrice.StorePrice,
	}
	s.SendTxBlock(sp.OperatorKey, msgUpdatePrice)
	queryGetSpStoragePriceByTimeResp, err = s.Client.QueryGetSpStoragePriceByTime(ctx, &sptypes.QueryGetSpStoragePriceByTimeRequest{
		SpAddr:    sp.OperatorKey.GetAddr().String(),
		Timestamp: time.Now().Unix(),
	})
	s.Require().NoError(err)
	s.T().Logf("queryGetSpStoragePriceByTimeResp read price: %s", queryGetSpStoragePriceByTimeResp.SpStoragePrice.ReadPrice)
}

// CreatePaymentAccount create new payment account and return latest payment account
func (s *PaymentTestSuite) CreatePaymentAccount(user keys.KeyManager, amount, decimal int64) string {
	ctx := context.Background()
	// create a new payment account
	msgCreatePaymentAccount := &paymenttypes.MsgCreatePaymentAccount{
		Creator: user.GetAddr().String(),
	}
	_ = s.SendTxBlock(user, msgCreatePaymentAccount)
	// query user's payment accounts
	queryGetPaymentAccountsByOwnerRequest := paymenttypes.QueryGetPaymentAccountsByOwnerRequest{
		Owner: user.GetAddr().String(),
	}
	paymentAccounts, err := s.Client.GetPaymentAccountsByOwner(ctx, &queryGetPaymentAccountsByOwnerRequest)
	s.Require().NoError(err)
	s.T().Log(paymentAccounts)
	paymentAccountAddr := paymentAccounts.PaymentAccounts[len(paymentAccounts.PaymentAccounts)-1]
	// charge payment account
	paymentAcc := sdk.MustAccAddressFromHex(paymentAccountAddr)
	msgSend := banktypes.NewMsgSend(user.GetAddr(), paymentAcc, []sdk.Coin{{Denom: "BNB", Amount: types.NewIntFromInt64WithDecimal(amount, decimal)}})
	s.SendTxBlock(user, msgSend)

	return paymentAccountAddr
}
