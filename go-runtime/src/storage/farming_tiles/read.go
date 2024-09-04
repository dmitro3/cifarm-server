package farming_tiles

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadFarmingTileObjectById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	id string,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_FARMING_TILES
	query := fmt.Sprintf(`+value.id:%s`, id)
	order := []string{
		"-create_time",
	}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}
	var object = objects.Objects[0]
	return object, nil
}

func ToFarmingTile(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.FarmingTile, error) {
	if object == nil {
		return nil, nil
	}
	var farmingTile *_collections.FarmingTile
	err := json.Unmarshal([]byte(object.Value), &farmingTile)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return farmingTile, nil
}
