package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := NewArgs(os.Args[1:])
	f := &f.F{}
	f.Shf("%s %s", args.action, args.path)
}

func NewArgs(in []string) *args {
	a := args{path: ".", action: "ls"}
	if len(in) > 0 {
		a.path = in[0]
		if len(in) > 1 {
			a.action = in[1]
		}
	}
	return &a
}

type args struct {
	path      string
	action    string
	dir       string
	filename  string
	extension string
	nameonly  string
}
