package actions

import (
	"path/filepath"

	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine
var Profiles = filepath.Join(".", "profiles")

func init() {
	r = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}
