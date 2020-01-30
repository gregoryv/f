package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := f.NewArgs(os.Args[1:])
	f := f.NewTerm()
	format, found := args.Format()
	if !found {
		return
	}
	f.Shf(format, args.Path)
}
