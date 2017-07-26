package common

// MailHandler 处理器
type MailHandler interface {
	PostMail(mailList []string, subject, content, attachment string)
}
