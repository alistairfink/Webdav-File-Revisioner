package webdav

import (
	Emailer "Webdav-File-Revisioner/emailer"

	"github.com/studio-b12/gowebdav"
)

type Webdav struct {
	webDavClient *gowebdav.Client
	emailer      Emailer.Emailer
}
