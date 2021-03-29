package mail

import (
	"godeliver/pkg/logging"
	"gopkg.in/gomail.v2"
)

//todo 分为pdf和mail两种形式。
func SendMail() {
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", "search@jareee.com")
	// 设置收件人
	//m.SetHeader("To", receiver)

	// 抄送
	//m.SetAddressHeader("")
	// 邮件标题
	m.SetHeader("Subject", "search订阅中有新内容")
	// 发送邮件
	// todo 生成配置模版
	//m.SetBody("text/html", mailBody)
	// 发送pdf
	//m.Attach("pdf")

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "search@jareee.com", "13858441809zjJ")

	if err := d.DialAndSend(m); err != nil {
		logging.Fatal("sendmail,err: ", err)
		panic(err)
	}
}

func SendMailPDF() {

}
