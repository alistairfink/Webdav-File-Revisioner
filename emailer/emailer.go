package emailer

import (
	Config "Webdav-File-Revisioner/config"
	"net/smtp"
)

func CreateEmailer(config Config.Config) Emailer {
	if !config.UseEmailer {
		return Emailer{
			useEmailer: false,
		}
	}

	return Emailer{
		useEmailer: true,
		fromEmail:  config.FromEmail,
		fromAuth:   smtp.PlainAuth("", config.FromEmail, config.FromEmailPassword, config.SmtpAuthServer),
		toEmail:    config.ToEmail,
		subject:    "An Error Occurred Backing Up: " + config.FilePath,
		smtpServer: config.SmtpSever,
	}
}

func (this *Emailer) SendErrorEmail(err error) {
	if !this.useEmailer {
		return
	}

	message := "Subject: " + this.subject + "\n\n" + err.Error()
	err = smtp.SendMail(
		this.smtpServer,
		this.fromAuth,
		this.fromEmail,
		[]string{this.toEmail},
		[]byte(message),
	)
	if err != nil {
		panic("This is bad: " + err.Error())
	}
}
