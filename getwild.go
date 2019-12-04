package getwild

import (
	"os"
	"path/filepath"
)

func init() {
	os.Args = getwild(os.Args)
}

func getwild(args []string) []string {
	result := make([]string, len(args))
	for _, arg := range args {
		match, err := filepath.Glob(arg)
		if err == nil && len(match) > 0 {
			result = append(result, match...)
		} else {
			result = append(result, arg)
		}
	}
	return result
}