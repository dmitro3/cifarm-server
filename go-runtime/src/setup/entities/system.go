package entities

import (
	_system "cifarm-server/src/storage/system"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSystem(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	users := _collections.Users{
		UserIds: []string{},
	}

	err := _system.WriteSystemUsersObject(ctx, logger, db, nk, users)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}