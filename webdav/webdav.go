package webdav

import (
	Config "Webdav-File-Revisioner/config"
	Emailer "Webdav-File-Revisioner/emailer"
	"fmt"
	"path/filepath"
	"time"

	"github.com/studio-b12/gowebdav"
)

func CreateWebdavClient(config Config.Config, emailer Emailer.Emailer) Webdav {
	goWebDavClient := gowebdav.NewClient(
		config.WebdavUrl,
		config.WebdavUser,
		config.WebdavPassword,
	)

	return Webdav{
		webDavClient: goWebDavClient,
		emailer:      emailer,
	}
}

func (this *Webdav) CopyFile(source, destination string) {
	fileName := filepath.Base(source)
	now := time.Now()
	formattedDate := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	newFileName := fileName + "." + formattedDate
	destination = destination + newFileName

	err := this.webDavClient.Copy(source, destination, true)
	if err != nil {
		this.emailer.SendErrorEmail(err)
		panic(err)
	}
}
