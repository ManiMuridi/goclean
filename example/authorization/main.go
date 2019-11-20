package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jBytes, _ := json.Marshal(struct{ Name string }{"Bob"})
	fmt.Println(jBytes)
	fmt.Println(string(jBytes))
	//config.Load()
	//policy := &security.Policy{
	//	Subject: "sales",
	//	Object:  "employee",
	//	Action:  "read",
	//}
	//
	//mongoUrl := fmt.Sprintf("%s:%s", config.GetString("mongodb.host"), config.GetString("mongodb.port"))
	//
	//a := mongodbadapter.NewAdapter(mongoUrl)
	//security.Configure("./rbac.conf", a)

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

	//authorized, err := security.AuthorizePolicy(policy)
	//
	//fmt.Println(authorized, err)
}
