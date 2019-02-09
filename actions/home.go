package actions

import (
	"github.com/MShoaei/command_control/models"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("index.html"))
}
