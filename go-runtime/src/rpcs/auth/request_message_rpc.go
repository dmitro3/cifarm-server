package rpcs_auth

import (
	services_cibase_authenticator_api "cifarm-server/src/services/cibase/api/authenticator"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type RequestMessageRpcResponse struct {
	Message string `json:"message"`
}

func RequestMessageRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	response, err := services_cibase_authenticator_api.RequestMessage(ctx, logger)
	if err != nil {
		return "", err
	}

	_response := &RequestMessageRpcResponse{Message: response.Message}

	out, err := json.Marshal(_response)
	if err != nil {
		return "", nil
	}

	return string(out), nil
}
