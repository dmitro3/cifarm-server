package placed_items

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

func WritePlacedItemObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	placedItem _collections.PlacedItem,
	key string,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	value, err := json.Marshal(placedItem)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_key := key
	if key == "" {
		_key = uuid.NewString()
	}

	write := &runtime.StorageWrite{
		UserID:          userId,
		Key:             _key,
		Collection:      _constants.COLLECTION_PLACED_ITEMS,
		Value:           string(value),
		PermissionRead:  2,
		PermissionWrite: 0,
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		write,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func WritePlacedItemObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	placedItems []_collections.PlacedItem,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	var writes []*runtime.StorageWrite

	for _, placedItem := range placedItems {
		value, err := json.Marshal(placedItem)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		write := &runtime.StorageWrite{
			UserID:          userId,
			Key:             uuid.NewString(),
			Collection:      _constants.COLLECTION_PLACED_ITEMS,
			Value:           string(value),
			PermissionRead:  1,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
