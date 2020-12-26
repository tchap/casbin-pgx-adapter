package pgxadapter

import (
	"strings"

	"github.com/casbin/casbin/v2/model"
)

// loadPolicyLine is based on https://pkg.go.dev/github.com/casbin/casbin/v2/persist#LoadPolicyLine
func loadPolicyLine(r *rule, m model.Model) {
	tokens := r.Tokens()
	key := tokens[0]
	sec := key[:1]
	m[sec][key].Policy = append(m[sec][key].Policy, tokens[1:])
	m[sec][key].PolicyMap[strings.Join(tokens[1:], model.DefaultSep)] = len(m[sec][key].Policy) - 1
}
