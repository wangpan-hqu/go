package gomail

import (
	"gopkg.in/gomail.v2"
)

//https://github.com/go-gomail/gomail
func Sendmail(from string, to string, subject string, body string, host string, port int, username string, password string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	//m.SetHeader("To", "qiujiahongde@163.com", "mail12@163.com")  //发送多个人
	m.SetHeader("To", to) //主送int
	//	m.SetHeader("Cc", "qiujiahongde@163.com") //抄送
	//	m.SetHeader("Bcc", "309284701@qq.com")  // 密送
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	//发送html格式邮件。
	m.SetBody("text/html", body)
	//	m.Attach("/home/Alex/lolcat.jpg")  //添加附件
	d := gomail.NewDialer(host, port, username, password)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
