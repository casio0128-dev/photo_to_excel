package commons

import (
	"os"
	"strings"
)

func CreateFilePath(filenames ...string) string {
	return strings.Join(filenames, string(os.PathSeparator))
}
