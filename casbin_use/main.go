package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	e, _ := casbin.NewSyncedEnforcer("acl_simple_model.conf", "acl_simple_policy.csv")
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	if passed, _ := e.Enforce(sub, obj, act); passed {
		// permit alice to read data1
		fmt.Println("Enforce policy passed.")
	} else {
		// deny the request, show an error
		fmt.Println("Enforce policy denied.")
	}
}
