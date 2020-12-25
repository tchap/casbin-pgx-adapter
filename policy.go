package pgxadapter

import (
	"strings"

	"github.com/casbin/casbin/v2/model"
)

func loadPolicyLine(r *rule, m model.Model) {
	key := r.PType.String
	sec := key[:1]
	m[sec][key].Policy = append(m[sec][key].Policy, tokens[1:])
	m[sec][key].PolicyMap[strings.Join(tokens[1:], model.DefaultSep)] = len(m[sec][key].Policy) - 1
}
