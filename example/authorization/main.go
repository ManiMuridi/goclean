package main

import (
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
)

type Policy struct {
	Subject string
	Object  string
	Action  string
}

func main() {
	p := &Policy{
		Subject: "sales",
		Object:  "employee",
		Action:  "read",
	}

	fmt.Println(Enforce(p))
}

func Enforce(policy *Policy) bool {
	e, err := casbin.NewEnforcer("./rbac.conf", "./policy.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if allowed, err := e.Enforce(policy.Subject, policy.Object, policy.Action); allowed && err == nil {
		// permit alice to read data1
		fmt.Println("Permitted")
		return true
	} else if err != nil {
		printErr(err)
		return false
	}

	fmt.Println("Not Permitted")

	return false
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
