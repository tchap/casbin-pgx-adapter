package pgxadapter

import (
	"context"

	"github.com/casbin/casbin/v2/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Adapter implements persist.Adapter and persist.BatchAdapter.
//
// persist.UpdatableAdapter and persist.FilteredAdapter interfaces are currently not supported.
type Adapter struct {
	pool      *pgxpool.Pool
	tableName string
}

func New(pool *pgxpool.Pool, tableName string) *Adapter {
	return &Adapter{
		pool:      pool,
		tableName: tableName,
	}
}

// LoadPolicy loads all policy rules from the storage.
func (adapter *Adapter) LoadPolicy(model model.Model) error {
	rules, err := listRules(adapter.pool, adapter.tableName)
	if err != nil {
		return err
	}
	for i := range rules {
		loadPolicyLine(&rules[i], model)
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (adapter *Adapter) SavePolicy(model model.Model) error {}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicy(section string, ptype string, rule []string) error {
	return insertRule(adapter.pool, adapter.tableName, ruleFromTokens(ptype, rule))
}

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicies(section string, ptype string, rules [][]string) error {
	return adapter.withTx(func(tx pgx.Tx) error {
		for _, tokens := range rules {
			if err := insertRule(tx, adapter.tableName, ruleFromTokens(ptype, tokens)); err != nil {
				return err
			}
		}
		return nil
	})
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicy(section string, ptype string, rule []string) error {
}

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicies(section string, ptype string, rules [][]string) error {}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemoveFilteredPolicy(section string, ptype string, fieldIndex int, fieldValues ...string) error {
}

func (adapter *Adapter) withTx(fnc func(pgx.Tx) error) error {
	ctx := context.Background()
	tx, err := adapter.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to begin a transaction")
	}

	if err := fnc(tx); err != nil {
		tx.Rollback(ctx) //nolint:errcheck
		return err
	}

	return errors.Wrap(tx.Commit(ctx), "failed to commit a transaction")
}
