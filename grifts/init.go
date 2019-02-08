package grifts

import (
	"github.com/MShoaei/command_control/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
