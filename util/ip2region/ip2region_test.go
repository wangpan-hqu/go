package ip2region

import (
	"fmt"
	"testing"
)

func TestGetRegion(t *testing.T) {
	fmt.Printf("region: %s", GetRegion("27.154.111.137"))
}
