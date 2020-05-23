package main

import (
	"os"

	fo "github.com/gregoryv/f"
)

func main() {
	args := fo.NewArgs(os.Args[1:]...)
	var act fo.Action
	if args.UseAction(&act) != nil {
		return
	}
	act(fo.NewTerm())
}
