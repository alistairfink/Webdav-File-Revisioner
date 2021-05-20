package main

import (
	Config "Webdav-File-Revisioner/config"
	Emailer "Webdav-File-Revisioner/emailer"
	"Webdav-File-Revisioner/webdav"
)

func main() {
	config := Config.GetConfig()
	emailer := Emailer.CreateEmailer(config)
	webDavClient := webdav.CreateWebdavClient(config, emailer)
	webDavClient.CopyFile(config.FilePath, config.DestinationPath)
}
