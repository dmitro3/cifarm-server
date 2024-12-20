package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_player "cifarm-server/src/collections/player"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UnfollowRpcParams struct {
	UserId string `json:"userId,omitempty"`
}

func UnfollowRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	// simply add to the user's followers list
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *FollowRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_player.ReadFollowings(ctx, logger, db, nk, collections_player.ReadFollowingsParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "followings not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	followings, err := collections_common.ToValue[collections_player.Followings](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	delete(followings.FollowedUsers, params.UserId)

	err = collections_player.WriteFollowings(ctx, logger, db, nk, collections_player.WriteFollowingsParams{
		UserId:     userId,
		Followings: *followings,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
