package grbac

import (
	"github.com/mikespook/gorbac"
)

func main() {
	rbac := gorbac.New()

	ra := gorbac.NewStdRole("admin")

	pa := gorbac.NewStdPermission("get")

	ra.Assign(pa)

	rbac.Add(ra)

	rbac.IsGranted("admin", pa, nil)

}
