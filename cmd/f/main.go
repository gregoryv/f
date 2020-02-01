package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := f.NewArgs(os.Args[1:]...)
	f := f.NewTerm()
	var format string
	err := args.Format(&format)
	if err != nil {
		return
	}
	f.Shf(format, args.Path)
}
