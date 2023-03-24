package gomail

import (
	"gopkg.in/gomail.v2"
)

//https://github.com/go-gomail/gomail
func gomail_use() {
	m := gomail.NewMessage()
	m.SetHeader("From", "wangpan_hqu@163.com")
	//m.SetHeader("To", "qiujiahongde@163.com", "mail12@163.com")  //发送多个人
	m.SetHeader("To", "1768044647@qq.com") //主送
	//	m.SetHeader("Cc", "qiujiahongde@163.com") //抄送
	//	m.SetHeader("Bcc", "309284701@qq.com")  // 密送
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	//发送html格式邮件。
	m.SetBody("text/html", "您的邮箱验证码: <p style='color:red'>red </p>")
	//	m.Attach("/home/Alex/lolcat.jpg")  //添加附件
	d := gomail.NewDialer("smtp.163.com", 25, "wangpan_hqu@163.com", "OEKTKKOHFPKOOTHS")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
