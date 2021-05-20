package config

type Config struct {
	FilePath          string
	DestinationPath   string
	WebdavUrl         string
	WebdavUser        string
	WebdavPassword    string
	UseEmailer        bool
	FromEmail         string
	FromEmailPassword string
	ToEmail           string
	SmtpSever         string
	SmtpAuthServer    string
}

// TODO: Add emailer info
