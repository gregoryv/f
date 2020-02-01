package f

import "path"

func NewArgs(in values) *Args {
	p := in.get(0, ".")
	a := Args{
		Path:     p,
		action:   in.get(1, "ls"),
		dir:      path.Dir(p),
		Ext:      path.Ext(p),
		nameonly: path.Base(p),
	}
	return &a
}

type values []string

func (v values) get(i int, def string) string {
	if i >= len(v) {
		return def
	}
	return v[i]
}

type Args struct {
	Path     string
	action   string
	dir      string
	filename string
	Ext      string
	nameonly string
}

func (a *Args) Format(format *string) error {
	f, found := shellFormats[a.action]
	if !found {
		return NotFound
	}
	*format = f
	return nil
}

var shellFormats = map[string]string{
	"ls": "ls %s",
	"f":  "ls -al %s",
}
