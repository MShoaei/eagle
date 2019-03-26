package utils

import (
	"os"
	"path"

	"github.com/spf13/afero"
)

var Fs = afero.NewOsFs()
var cwd, _ = os.Getwd()
var ProfilesDir = path.Join(cwd, "profiles")
