package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func dingtalk_use() {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) // 毫秒时间戳
	state := "my info to decision"
	signature := EncodeSHA256(timestamp, "otESK06u4P8EetuTqqKgGjU8nZYI0yHNtl1lfRFob38ePCFEEYJODAepESBWP1fj") // 加密签名  加密算法见我另一个函数
	url2 := fmt.Sprintf(
		"https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s",
		"dingfv1zcj2uozqabeow", timestamp, signature)

	p := struct {
		Tmp_auth_code string `json:"tmp_auth_code"`
		state         string `json:"state"`
	}{"513fbefe737a3321b920bb0e3bb227cb", state} // post数据
	p1, _ := json.Marshal(p)
	p2 := string(p1)
	p3 := strings.NewReader(p2) //构建post数据

	resp, err := http.Post(url2, "application/json;charset=UTF-8", p3)
	//fmt.Println(1, resp, err)
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(2, string(body), err)
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i) ///返回的数据给i
	fmt.Println(1111, string(body), err)

}
func EncodeSHA256(message, secret string) string {
	// 钉钉签名算法实现
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sum := h.Sum(nil) // 二进制流
	message1 := base64.StdEncoding.EncodeToString(sum)

	uv := url.Values{}
	uv.Add("0", message1)
	message2 := uv.Encode()[2:]
	return message2
}
