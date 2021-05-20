package emailer

import "net/smtp"

type Emailer struct {
	useEmailer bool
	fromEmail  string
	fromAuth   smtp.Auth
	toEmail    string
	subject    string
	smtpServer string
}
