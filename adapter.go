package pgxadapter

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Adapter implements persist.Adapter and persist.BatchAdapter.
//
// persist.UpdatableAdapter and persist.FilteredAdapter interfaces are currently not supported.
type Adapter struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Adapter {
	return &Adapter{pool: pool}
}

// LoadPolicy loads all policy rules from the storage.
func (adapter *Adapter) LoadPolicy(model model.Model) error {
	rules, err := listRules(adapter.pool)
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
}

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicies(section string, ptype string, rules [][]string) error {}

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
