package adapter

import (
	"github.com/casbin/casbin/v2/model"
)

type Adapter struct{}

// LoadPolicy loads all policy rules from the storage.
func (adapter *Adapter) LoadPolicy(model model.Model) error {
}

// SavePolicy saves all policy rules to the storage.
func (adapter *Adapter) SavePolicy(model model.Model) error {}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
}

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {}

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (adapter *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (adapter *Adapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {}

// IsFiltered returns true if the loaded policy has been filtered.
func (adapter *Adapter) IsFiltered() bool {}
