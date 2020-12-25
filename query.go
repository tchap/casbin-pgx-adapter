package pgxadapter

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type queryExecutor interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type scanner interface {
	Scan(dest ...interface{}) error
}

type rule struct {
	PType pgtype.Text
	V0    pgtype.Text
	V1    pgtype.Text
	V2    pgtype.Text
	V3    pgtype.Text
	V4    pgtype.Text
	V5    pgtype.Text
}

func listRules(qe queryExecutor) ([]rule, error) {
	panic("not implemented")
}
