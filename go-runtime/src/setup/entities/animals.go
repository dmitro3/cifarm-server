package setup_entities

import (
	collections_animals "cifarm-server/src/collections/animals"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupAnimals(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	animals := []collections_animals.Animal{
		{
			Key:                       collections_animals.KEY_CHICKEN,
			OffspringPrice:            1000,
			IsNFT:                     false,
			GrowthTime:                60 * 60 * 7, //7 days
			YieldTime:                 60 * 60,     //1 days
			AvailableInShop:           true,
			HungerTime:                60 * 12, //12 hours
			MinHarvestQuantity:        14,
			MaxHarvestQuantity:        20,
			BasicHarvestExperiences:   32,
			PremiumHarvestExperiences: 96,
		},
		{
			Key:                       collections_animals.KEY_COW,
			IsNFT:                     true,
			GrowthTime:                60 * 60 * 14, //14 days
			YieldTime:                 60 * 60 * 2,  //2 days
			AvailableInShop:           false,
			HungerTime:                60 * 12, //12 hours
			MinHarvestQuantity:        14,
			MaxHarvestQuantity:        20,
			BasicHarvestExperiences:   32,
			PremiumHarvestExperiences: 96,
		},
	}

	err := collections_animals.WriteMany(ctx, logger, db, nk, collections_animals.WriteManyParams{
		Animals: animals,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
