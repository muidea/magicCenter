package net

import (
	"net/smtp"
	"strings"
)

// SendMail 发送邮件
// user : example@example.com login smtp server user
// password: xxxxx login smtp server password
// host: smtp.example.com:port   smtp.163.com:25
// to: example@example.com;example1@163.com;example2@sina.com.cn;...
// subject:The subject of mail
// body: The content of mail
// mailtyoe: mail type html or text
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
