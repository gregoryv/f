package main

import (
	"os"

	"github.com/gregoryv/f"
)

func main() {
	args := NewArgs(os.Args[1:])
	f := f.NewTerm()
	format, found := args.Format()
	if !found {
		return
	}
	f.Shf(format, args.path)
}

func NewArgs(in []string) *args {
	a := args{path: "."}
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

func (a *args) Format() (string, bool) {
	format, found := shellFormats[a.action]
	if !found {
		return "", false
	}
	return format, true
}

var shellFormats = map[string]string{
	"":  "ls %s",
	"f": "ls -al %s",
}
