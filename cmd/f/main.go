package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := f.NewArgs(os.Args[1:]...)
	var act f.Action
	if args.UseAction(&act) != nil {
		return
	}
	act(f.NewTerm())
}
