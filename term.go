package f

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gregoryv/fox"
)

func NewTerm() *Term {
	m := Term{
		Logger: fox.NewSyncLog(os.Stdout).FilterEmpty(),
		exit:   os.Exit,
	}
	m.errFuncs = []errLiner{
		func(s *string) {
			var cli string
			if EmacsOpen(&cli, *s) == nil {
				// Only open files within the working directory
				if strings.Index(cli, m.wd) > -1 {
					// don't use m.Sh as recursive errors are bad
					c := strings.Split(cli, " ")
					exec.Command(c[0], c[1:]...).Start()
				}
			}
		},
		func(s *string) { Color(s, m.wd) },
		func(s *string) { Strip(s, m.wd) },
		func(s *string) { Color(s, "_test.go") },
	}

	dir, _ := os.Getwd()
	m.wd = dir + "/"
	return &m
}

type Term struct {
	fox.Logger
	Verbose  bool
	exit     func(int)
	wd       string
	errFuncs []errLiner
}

type errLiner func(*string)

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
		m.handleErrLines(out)
		return err
	}
	nice := bytes.TrimSpace(out)
	if len(nice) > 0 {
		fmt.Println(string(nice))
	}
	m.Log("# ", cli, " ", time.Since(start))
	return nil
}

func (m *Term) handleErrLines(out []byte) {
	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		s := string(line)
		for _, fn := range m.errFuncs {
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

var (
	red         = "\033[31m"
	reset       = "\033[0m"
	notColored  = fmt.Errorf("not colored")
	notStripped = fmt.Errorf("not stripped")
)
