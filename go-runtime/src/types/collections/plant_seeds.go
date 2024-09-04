package collections

type PlantSeed struct {
	Id                          string `json:"id"`
	GrowthStageDuration         int64  `json:"growthStageDuration"`
	GrowthStages                int    `json:"growthStages"`
	SeedPrice                   int64  `json:"price"`
	Premium                     bool   `json:"premium"`
	Perennial                   bool   `json:"perennial"`
	NextGrowthStageAfterHarvest int    `json:"nextGrowthStageAfterHarvest"`
	MinHarvestQuantity          int    `json:"minHarvestQuantity"`
	MaxHarvestQuantity          int    `json:"maxHarvestQuantity"`
}
