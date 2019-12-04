package getwild

import (
	"os"
	"path/filepath"
)

func init() {
	args := os.Args
	for i, arg := range args {
		match, err := filepath.Glob(arg)
		if err == nil && len(match) > 0 {
			args[i] = match[0]
		}
	}
	os.Args = args
}
