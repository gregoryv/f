package f

import "path"

func NewArgs(in ...string) *Args {
	a := Args{
		Path:   ".",
		action: "ls",
	}
	set(&a.Path, in, 0)
	set(&a.action, in, 1)
	p := a.Path
	a.dir = path.Dir(p)
	a.Ext = path.Ext(p)
	a.nameonly = path.Base(p)
	return &a
}

func set(v *string, in []string, i int) error {
	if i >= len(in) {
		return NotFound
	}
	*v = in[i]
	return nil
}

type Args struct {
	Path     string
	action   string
	dir      string
	filename string
	Ext      string
	nameonly string
}

func (a *Args) UseAction(fn *Action) error {
	var f string
	err := a.Format(&f)
	if err == nil {
		*fn = func(m *Term) {
			m.Shf(f, a.Path)
		}
		return nil
	}

	return NotFound
}

type Action func(*Term)

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
	"e":  "emacsclient -n %s",
}
