package adapter

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type rule struct {
	PType string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

// Adapter implements persist.Adapter, persist.BatchAdapter and persist.UpdatableAdapter.
//
// persist.FilteredAdapter interface is currently not implemented.
type Adapter struct {
	pool *pgxpool.Pool
}

func NewAdapter(pool *pgxpool.Pool) *Adapter {
	return &Adapter{pool: pool}
}

// LoadPolicy loads all policy rules from the storage.
func (adapter *Adapter) LoadPolicy(model model.Model) error {
}

// SavePolicy saves all policy rules to the storage.
func (adapter *Adapter) SavePolicy(model model.Model) error {}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
}

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
}

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
}

// UpdatePolicy updates a policy rule from storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {}
