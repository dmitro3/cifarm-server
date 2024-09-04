package config

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadConfigPlayerMetdataObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*api.StorageObject, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	name := _constants.STORAGE_INDEX_CONFIG_PLAYER_METADATA
	query := ""
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}

	object := objects.Objects[0]
	return object, nil
}

func ToConfigPlayerMetdata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.PlayerMetadata, error) {
	if object == nil {
		return nil, nil
	}
	var playerMetdata *_collections.PlayerMetadata
	err := json.Unmarshal([]byte(object.Value), &playerMetdata)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return playerMetdata, nil
}