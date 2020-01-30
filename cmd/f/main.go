package main

import (
	"fmt"
	"os"
)

func main() {
	args := NewArgs(os.Args[1:])
	fmt.Println(args.action, args.path)
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
