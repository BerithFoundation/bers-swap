// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package mariadb

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateBersSwapHistory(ctx context.Context, arg CreateBersSwapHistoryParams) (sql.Result, error)
	GetBersSwapHistory(ctx context.Context, senderTxHash string) (BersSwapHist, error)
}

var _ Querier = (*Queries)(nil)
