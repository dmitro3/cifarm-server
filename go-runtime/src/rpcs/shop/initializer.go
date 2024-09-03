package shop

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeShop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("go_buy_plant_seed", BuyPlantSeedRpc)
	if err != nil {
		return err
	}

	return nil
}