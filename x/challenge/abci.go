package challenge

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	k "github.com/bnb-chain/greenfield/x/challenge/keeper"
	"github.com/bnb-chain/greenfield/x/challenge/types"
	sptypes "github.com/bnb-chain/greenfield/x/sp/types"
	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"
)

func BeginBlocker(ctx sdk.Context, keeper k.Keeper) {
	// reset count of challenge in current block to zero
	keeper.ResetChallengeCountCurrentBlock(ctx)

	// delete too old slashes at this height
	coolingOff := keeper.SlashCoolingOffPeriod(ctx)
	blockHeight := uint64(ctx.BlockHeight())
	if blockHeight <= coolingOff {
		return
	}

	height := blockHeight - coolingOff
	keeper.RemoveSlashUntil(ctx, height)
}

func EndBlocker(ctx sdk.Context, keeper k.Keeper) {
	count := keeper.GetChallengeCountCurrentBlock(ctx)
	needed := keeper.ChallengeCountPerBlock(ctx)
	if count >= needed {
		return
	}

	events := make([]proto.Message, 0)                      // for events
	objectMap := make(map[string]struct{})                  // for de-duplication
	iteration, maxIteration := uint64(0), 10*(needed-count) // to prevent endless loop
	for count < needed && iteration < maxIteration {
		iteration++
		seed := k.SeedFromRandaoMix(ctx.BlockHeader().RandaoMix, iteration)

		// random object info
		objectKey := k.RandomObjectKey(seed)
		objectInfo, found := keeper.StorageKeeper.GetObjectAfterKey(ctx, objectKey)
		if !found { // there is no object info yet, cannot generate challenges
			return
		}

		if objectInfo.ObjectStatus != storagetypes.OBJECT_STATUS_IN_SERVICE {
			continue
		}

		// random redundancy index (sp address)
		var spOperatorAddress string
		secondarySpAddresses := objectInfo.SecondarySpAddresses

		bucket, found := keeper.StorageKeeper.GetBucket(ctx, objectInfo.BucketName)
		if !found {
			continue
		}

		redundancyIndex := k.RandomRedundancyIndex(seed, uint64(len(secondarySpAddresses)+1))
		redundancyIndex--

		if redundancyIndex == types.RedundancyIndexPrimary { // primary sp
			spOperatorAddress = bucket.PrimarySpAddress
		} else {
			spOperatorAddress = objectInfo.SecondarySpAddresses[redundancyIndex]
		}

		addr, err := sdk.AccAddressFromHexUnsafe(spOperatorAddress)
		if err != nil {
			continue
		}
		sp, found := keeper.SpKeeper.GetStorageProvider(ctx, addr)
		if !found || sp.Status != sptypes.STATUS_IN_SERVICE {
			continue
		}

		mapKey := fmt.Sprintf("%s-%d", spOperatorAddress, objectInfo.Id)
		if _, ok := objectMap[mapKey]; ok { // already generated for this pair
			continue
		}

		// check recent slash
		if keeper.ExistsSlash(ctx, strings.ToLower(spOperatorAddress), objectInfo.Id.Uint64()) {
			continue
		}

		// random segment/piece index
		segments := k.CalculateSegments(objectInfo.PayloadSize, keeper.StorageKeeper.MaxSegmentSize(ctx))
		segmentIndex := k.RandomSegmentIndex(seed, segments)

		objectMap[mapKey] = struct{}{}
		challengeId := keeper.GetOngoingChallengeId(ctx)
		keeper.SetOngoingChallengeId(ctx, challengeId+1)
		events = append(events, &types.EventStartChallenge{
			ChallengeId:       challengeId,
			ObjectId:          objectInfo.Id.Uint64(),
			SegmentIndex:      segmentIndex,
			SpOperatorAddress: spOperatorAddress,
			RedundancyIndex:   redundancyIndex,
		})

		count++
	}
	_ = ctx.EventManager().EmitTypedEvents(events...)
}
