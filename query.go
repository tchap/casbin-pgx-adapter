package pgxadapter

import (
	"context"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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

func ruleFromTokens(ptype string, ruleTokens []string) *rule {
	switch len(ruleTokens) {
	case 2:
		return &rule{
			PType: pgtype.Text{Status: pgtype.Present, String: ptype},
			V0:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[0]},
			V1:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[1]},
			V2:    pgtype.Text{Status: pgtype.Null},
			V3:    pgtype.Text{Status: pgtype.Null},
			V4:    pgtype.Text{Status: pgtype.Null},
			V5:    pgtype.Text{Status: pgtype.Null},
		}

	case 3:
		return &rule{
			PType: pgtype.Text{Status: pgtype.Present, String: ptype},
			V0:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[0]},
			V1:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[1]},
			V2:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[2]},
			V3:    pgtype.Text{Status: pgtype.Null},
			V4:    pgtype.Text{Status: pgtype.Null},
			V5:    pgtype.Text{Status: pgtype.Null},
		}

	case 4:
		return &rule{
			PType: pgtype.Text{Status: pgtype.Present, String: ptype},
			V0:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[0]},
			V1:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[1]},
			V2:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[2]},
			V3:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[3]},
			V4:    pgtype.Text{Status: pgtype.Null},
			V5:    pgtype.Text{Status: pgtype.Null},
		}

	case 5:
		return &rule{
			PType: pgtype.Text{Status: pgtype.Present, String: ptype},
			V0:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[0]},
			V1:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[1]},
			V2:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[2]},
			V3:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[3]},
			V4:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[4]},
			V5:    pgtype.Text{Status: pgtype.Null},
		}

	default:
		return &rule{
			PType: pgtype.Text{Status: pgtype.Present, String: ptype},
			V0:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[0]},
			V1:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[1]},
			V2:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[2]},
			V3:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[3]},
			V4:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[4]},
			V5:    pgtype.Text{Status: pgtype.Present, String: ruleTokens[5]},
		}
	}
}

func listRules(qe queryExecutor, tableName string) ([]rule, error) {
	var query strings.Builder
	query.WriteString("SELECT ptype, v0, v1, v2, v3, v4, v5 FROM ")
	query.WriteString(tableName)

	rows, err := qe.Query(context.Background(), query.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to list rules")
	}

	var rules []rule
	for rows.Next() {
		var r rule
		if err := scanRule(rows, &r); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to scan rows")
	}
	return rules, nil
}

func insertRule(qe queryExecutor, tableName string, r *rule) error {
	var query strings.Builder
	query.WriteString("INSERT INTO ")
	query.WriteString(tableName)
	query.WriteString(" (ptype, v0, v1, v2, v3, v4, v5) VALUES ($1, $2, $3, $4, $5, $6, $7")

	_, err := qe.Exec(context.Background(), query.String(), r.PType, r.V0, r.V1, r.V2, r.V3, r.V4, r.V5)
	return errors.Wrap(err, "failed to insert a rule")
}

func scanRule(s scanner, dst *rule) error {
	return errors.Wrap(s.Scan(
		&dst.PType,
		&dst.V0,
		&dst.V1,
		&dst.V2,
		&dst.V3,
		&dst.V4,
		&dst.V5,
	), "failed to scan a rule")
}
