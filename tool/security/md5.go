package security

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func ToMd5String(planText string) string {
	if len(planText) == 0 {
		return planText
	}

	md5Bytes := md5.Sum([]byte(planText))
	return strings.ToUpper(fmt.Sprintf("%x", md5Bytes))
}
