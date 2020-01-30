package f

func NewArgs(in []string) *Args {
	a := Args{Path: "."}
	if len(in) > 0 {
		a.Path = in[0]
		if len(in) > 1 {
			a.action = in[1]
		}
	}
	return &a
}

type Args struct {
	Path      string
	action    string
	dir       string
	filename  string
	extension string
	nameonly  string
}

func (a *Args) Format() (string, bool) {
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
