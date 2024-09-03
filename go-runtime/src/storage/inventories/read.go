package inventories

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadInventoryObjectParams struct {
	Id string `json:"id"`
}

func ReadInventoryObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadInventoryObjectParams,
) (*api.StorageObject, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	name := _constants.STORAGE_INDEX_INVENTORY_OBJECTS
	query := fmt.Sprintf("+value.id:%s +user_id:%s", params.Id, userId)
	order := []string{}

	inventories, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(inventories.Objects) == 0 {
		return nil, nil
	}
	var inventory = inventories.Objects[0]
	return inventory, nil
}

func ToReadInventoryObjectValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.Inventory, error) {
	if object == nil {
		return nil, nil
	}

	var inventory *_collections.Inventory
	err := json.Unmarshal([]byte(object.Value), &inventory)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return inventory, nil
}
