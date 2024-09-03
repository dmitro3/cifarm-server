package rpcs

import (
	_daily_rewards "cifarm-server/src/rpcs/daily_rewards"
	_shop "cifarm-server/src/rpcs/shop"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeRpcs(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _daily_rewards.InitializeDailyReward(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = _shop.InitializeShop(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
