package f

import (
	"fmt"
	"strings"
)

func TidyImports(m *Term, a *Args) error {
	if a.Ext != ".go" {
		return InvalidExtension
	}
	m.Shf("goimports -w %s", a.Path)
	return nil
}

var InvalidExtension = fmt.Errorf("invalid extension")

// EmacsOpen parses v for file/path:LINENO and sets cli to open
func EmacsOpen(cli *string, v string) error {
	v = strings.TrimSpace(v)
	first := strings.Split(v, " ")[0]
	parts := strings.Split(first, ":")
	if len(parts) > 1 {
		lineno := parts[1]
		path := parts[0]
		*cli = fmt.Sprintf("emacsclient -n +%s %s", lineno, path)
		return nil
	}
	return NotFound
}

var NotFound = fmt.Errorf("not found")
