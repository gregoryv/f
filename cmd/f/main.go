package main

import (
	"os"
)

func main() {
	args := NewArgs(os.Args[1:]...)
	var act Action
	if args.UseAction(&act) != nil {
		return
	}
	act(NewTerm())
}

func TidyImports(args ...string) error {
	a := NewArgs(args...)
	if len(args) == 0 {
		a = NewArgs(os.Args[1:]...)
	}
	if a.Ext != ".go" {
		return InvalidExtension
	}
	return Shf("goimports -w %s", a.Path)
}
