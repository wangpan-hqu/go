package gomail

import (
	"testing"
)

func TestSendmail(t *testing.T) {

	Sendmail("wangpan_hqu@163.com", "1768044647@qq.com", "Hello!", "您的邮箱验证码: <p style='color:red'>red </p>", "smtp.163.com", 25, "wangpan_hqu@163.com", "OEKTKKOHFPKOOTHS")

}
