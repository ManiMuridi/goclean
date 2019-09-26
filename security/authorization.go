package security

import (
	"github.com/casbin/casbin/v2"
)

var (
	enforcer *casbin.Enforcer
)

// Configure loads and bootstraps the authorization enforcer.
func Configure(confPath string, data interface{}) {
	var err error
	if enforcer, err = casbin.NewEnforcer(confPath, data); err != nil {
		panic(err)
	}
}
