package pgxadapter

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type queryExecutor interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
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

func (r *rule) Tokens() []string {
	// V0 and V1 are always present, right?
	switch {
	case r.V2.Status == pgtype.Null:
		return []string{
			r.PType.String,
			r.V0.String,
			r.V1.String,
		}

	case r.V3.Status == pgtype.Null:
		return []string{
			r.PType.String,
			r.V0.String,
			r.V1.String,
			r.V2.String,
		}

	case r.V4.Status == pgtype.Null:
		return []string{
			r.PType.String,
			r.V0.String,
			r.V1.String,
			r.V2.String,
			r.V3.String,
		}

	case r.V5.Status == pgtype.Null:
		return []string{
			r.PType.String,
			r.V0.String,
			r.V1.String,
			r.V2.String,
			r.V3.String,
			r.V4.String,
		}

	default:
		return []string{
			r.PType.String,
			r.V0.String,
			r.V1.String,
			r.V2.String,
			r.V3.String,
			r.V4.String,
			r.V5.String,
		}
	}
}

func listRules(qe queryExecutor) ([]rule, error) {
	panic("not implemented")
}
