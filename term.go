package f

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/gregoryv/fox"
)

func NewTerm() *Term {
	m := Term{
		Logger: fox.NewSyncLog(os.Stdout).FilterEmpty(),
		exit:   os.Exit,
	}
	m.errFuncs = []liner{
		func(s *string) {
			OpenError(s, m.wd)
		},
		func(s *string) { Color(s, m.wd) },
		//		func(s *string) { Strip(s, m.wd) },
		func(s *string) { Color(s, "_test.go") },
	}
	m.okFuncs = []liner{}
	dir, _ := os.Getwd()
	m.wd = dir + "/"
	return &m
}

type Term struct {
	fox.Logger
	Verbose  bool
	exit     func(int)
	wd       string
	errFuncs []liner
	okFuncs  []liner
}

// liner funcs modify an output line
type liner func(*string)

func (m *Term) Log(p ...interface{}) {
	if m.Verbose {
		m.Logger.Log(p...)
	}
}

func (m *Term) Shf(format string, args ...interface{}) error {
	return m.Sh(fmt.Sprintf(format, args...))
}

func (m *Term) Sh(cli string) error {
	start := time.Now()
	m.Log("# ", cli)

	p := strings.Split(cli, " ")
	out, err := exec.Command(p[0], p[1:]...).CombinedOutput()
	if err != nil {
		m.adaptOutput(out, m.errFuncs)
		m.exit(1)
		return err
	}
	nice := bytes.TrimSpace(out)
	if len(nice) > 0 {
		m.adaptOutput(nice, m.okFuncs)
	}
	m.Log("# ", cli, " ", time.Since(start))
	return nil
}

func (m *Term) adaptOutput(out []byte, liners []liner) {
	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		s := string(line)
		for _, fn := range liners {
			fn(&s)
		}
		fmt.Println(s)
	}
}

func Color(line *string, contains string) error {
	found := (strings.Index(*line, contains) > -1)
	if !found {
		return notColored
	}
	*line = red + *line + reset
	return nil
}

func Strip(line *string, part string) error {
	stripped := strings.ReplaceAll(*line, part, "")
	if stripped == *line {
		return notStripped
	}
	*line = stripped
	return nil
}

func OpenError(s *string, wd string) error {
	var cli string
	err := EmacsOpen(&cli, *s)
	if err != nil {
		return err
	}
	c := strings.Split(cli, " ")
	// emacsclient -n +lineno path
	_, err = os.Stat(path.Join(wd, c[3]))
	isLocal := (err == nil)
	// Only open files within the working directory
	if strings.Index(cli, wd) > -1 || isLocal {
		// don't use m.Sh as recursive errors are bad
		exec.Command(c[0], c[1:]...).Run()
	}
	return nil
}

var (
	red         = "\033[31m"
	reset       = "\033[0m"
	notColored  = fmt.Errorf("not colored")
	notStripped = fmt.Errorf("not stripped")
)
