package setup_entities

import (
	collections_seeds "cifarm-server/src/collections/seeds"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSeeds(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	seeds := []collections_seeds.Seed{
		{
			Key:                       collections_seeds.KEY_CARROT,
			Price:                     50,
			GrowthStageDuration:       60 * 60, //1 hours
			GrowthStages:              5,
			Premium:                   false,
			Perennial:                 false,
			MinHarvestQuantity:        14,
			MaxHarvestQuantity:        20,
			BasicHarvestExperiences:   12,
			PremiumHarvestExperiences: 60,
		},
		{
			Key:                         collections_seeds.KEY_POTATO,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
		},
		{
			Key:                         collections_seeds.KEY_CUCUMBER,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
		},
		{
			Key:                         collections_seeds.KEY_PINEAPPLE,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
		},
		{
			Key:                         collections_seeds.KEY_WATERMELON,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
		},
		{
			Key:                         collections_seeds.KEY_PINEAPPLE,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
		},
	}

	err := collections_seeds.WriteMany(ctx, logger, db, nk, collections_seeds.WriteManyParams{
		Seeds: seeds,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
