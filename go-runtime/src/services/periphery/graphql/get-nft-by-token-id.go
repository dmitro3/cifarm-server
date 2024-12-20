package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftByTokenIdInput struct {
	TokenId          string `json:"tokenId,omitempty"`
	Network          string `json:"network,omitempty"`
	NftCollectionKey string `json:"nftCollectionKey,omitempty"`
	ChainKey         string `json:"chainKey,omitempty"`
}

type GetNftByTokenIdArgs struct {
	Input GetNftByTokenIdInput `json:"input,omitempty"`
}

func GetNftByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	args GetNftByTokenIdArgs,
) (*NftDataResponse, error) {
	url, err := config.CifarmPeripheryGraphqlUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	client := graphql.NewClient(url, nil)
	logger.Info("%v", args.Input.TokenId)
	query := `query Query($args: GetNftByTokenIdArgs!) {
  nftByTokenId(args: $args) {
    ownerAddress,
    tokenId,
    metadata {
        image,
        properties
    }
  }
}`
	variables := map[string]interface{}{
		"args": args,
	}
	result := struct {
		NftByTokenId NftDataResponse `json:"nftByTokenId,omitempty"`
	}{}

	err = client.WithDebug(true).Exec(context.Background(),
		query,
		&result,
		variables,
	)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info(result.NftByTokenId.OwnerAddress)
	return &result.NftByTokenId, nil
}
