package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := f.NewArgs(os.Args[1:]...)
	var fn f.Action
	err := args.UseAction(&fn)
	if err != nil {
		return
	}
	m := f.NewTerm()
	fn(m)
}
