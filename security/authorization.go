package security

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
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

func ConfigureModel(m model.Model, data interface{}) {
	var err error
	if enforcer, err = casbin.NewEnforcer(m, data); err != nil {
		panic(err)
	}
}
