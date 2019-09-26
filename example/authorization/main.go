package main

import (
	"fmt"

	mongodbadapter "github.com/casbin/mongodb-adapter/v2"

	"github.com/ManiMuridi/goclean/config"
	"github.com/ManiMuridi/goclean/security"
)

func main() {
	config.Load()
	policy := &security.Policy{
		Subject: "sales",
		Object:  "employee",
		Action:  "read",
	}

	mongoUrl := fmt.Sprintf("%s:%s", config.GetString("mongodb.host"), config.GetString("mongodb.port"))

	a := mongodbadapter.NewAdapter(mongoUrl)
	security.Configure("./rbac.conf", a)

	//security.AddPolicy(&security.RbacPolicy{
	//	Subject: "sales",
	//	Object:  "employee",
	//	Action:  "read",
	//	Effect:  "allow",
	//})

	//security.RemovePolicy(&security.RbacPolicy{
	//	Subject: "sales",
	//	Object:  "employee",
	//	Action:  "read",
	//	Effect:  "allow",
	//})

	authorized, err := security.AuthorizePolicy(policy)

	fmt.Println(authorized, err)
}
