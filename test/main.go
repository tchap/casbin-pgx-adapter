package main

import (
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e.EnableLog(true)
	m := e.GetModel()
	m.PrintModel()
}
