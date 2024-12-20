package mailer

import "embed"

const (
	FromName           = "social"
	maxRetires         = 3
	UsrWelcomeTemplate = "user_invitation.tmpl"
)

var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
