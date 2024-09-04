package farming_tools

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteFarmingToolObjectsParams struct {
	FarmingTools []_collections.FarmingTool
}

func WriteFarmingToolObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteFarmingToolObjectsParams,
) error {
	var writes []*runtime.StorageWrite
	for _, farmingTool := range params.FarmingTools {
		value, err := json.Marshal(farmingTool)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_FARMING_TOOLS,
			Key:             farmingTool.Id,
			Value:           string(value),
			PermissionRead:  2,
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