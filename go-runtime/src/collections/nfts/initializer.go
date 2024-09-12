package collections_nfts

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := RegisterByReferenceKey(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = RegisterByTokenId(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
