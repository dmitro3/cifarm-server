package setup_entities

import (
	collections_buildings "cifarm-server/src/collections/buildings"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupBuildings(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	buildings := []collections_buildings.Building{
		{
			Key:             collections_buildings.KEY_COOP,
			Price:           1000,
			AvailableInShop: true,
		},
		{
			Key:             collections_buildings.KEY_PASTURE,
			Price:           2500,
			AvailableInShop: true,
		},
		{
			Key: collections_buildings.KEY_HOME,
		},
	}

	err := collections_buildings.WriteMany(ctx, logger, db, nk, collections_buildings.WriteManyParams{
		Buildings: buildings,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
