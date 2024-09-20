package rpcs_farming

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_tiles "cifarm-server/src/collections/tiles"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HarvestCropRpcParams struct {
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

type HarvestCropRpcResponse struct {
	HarvestedCropInventoryKey string `json:"harvestedCropInventoryKey"`
}

func HarvestCropRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *HarvestCropRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemTileKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "tile not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	tile, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value1, err := json.Marshal(tile)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	logger.Info(string(value1))

	if !tile.IsPlanted {
		errMsg := "tile is not being planted"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !tile.FullyMatured {
		errMsg := "plant not fully matured"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	isPremium := tile.ReferenceKey == collections_tiles.KEY_PREMIUM
	//write to inventories the havested items
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: tile.SeedGrowthInfo.Crop.Key,
			Type:         collections_inventories.TYPE_HARVESTED_PLANT,
			Quantity:     tile.SeedGrowthInfo.HarvestQuantityRemaining,
			IsPremium:    isPremium,
			Deliverable:  true,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var experiences int64
	if isPremium {
		experiences = tile.SeedGrowthInfo.Crop.PremiumHarvestExperiences
	} else {
		experiences = tile.SeedGrowthInfo.Crop.BasicHarvestExperiences
	}

	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		Amount: experiences,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update tile status
	tile.FullyMatured = false
	tile.IsPlanted = false
	tile.SeedGrowthInfo = collections_placed_items.SeedGrowthInfo{}

	//update the tile
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *tile,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(HarvestCropRpcResponse{
		HarvestedCropInventoryKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}