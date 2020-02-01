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

// EmacsOpen parses v for file/path:LINENO and returns cli to open
func EmacsOpen(cli *string, v string) error {
	v = strings.TrimSpace(v)
	parts := strings.Split(v, ":")
	if len(parts) > 1 {
		p := parts[0]
		if p[0] != '/' {
			return NotFound
		}
		lineno := strings.Split(parts[1], " ")[0]
		*cli = fmt.Sprintf("emacsclient -n +%s %s", lineno, parts[0])
		return nil
	}
	return NotFound
}

var NotFound = fmt.Errorf("not found")
