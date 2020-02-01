package f

import "fmt"

func TidyImports(m *Term, a *Args) error {
	if a.Ext != ".go" {
		return InvalidExtension
	}
	m.Shf("goimports -w %s", a.Path)
	return nil
}

var InvalidExtension = fmt.Errorf("invalid extension")
