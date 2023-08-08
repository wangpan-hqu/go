package security

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func sign_use() {
	stringList := []string{"accessToken", "timestamp", "nonce"}
	sort.Strings(stringList)
	signOrigin := strings.Replace(strings.Trim(fmt.Sprint(stringList), "[]"), " ", "", -1)
	h := sha1.New()
	if _, err := h.Write([]byte(signOrigin)); err != nil {
		fmt.Println(err)
	}
	bs := h.Sum([]byte("a"))
	fmt.Println(bs)
	fmt.Println(hex.EncodeToString(bs))

}

func GenerateSign(timestamp string, nonce string, accessToken string) (string, error) {
	stringList := []string{accessToken, timestamp, nonce}
	sort.Strings(stringList)
	signOrigin := strings.Replace(strings.Trim(fmt.Sprint(stringList), "[]"), " ", "", -1)
	h := sha1.New()
	if _, err := h.Write([]byte(signOrigin)); err != nil {
		return "", err
	}

	bs := h.Sum(nil)
	signEncoded := hex.EncodeToString(bs)
	return signEncoded, nil
}
