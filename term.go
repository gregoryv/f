package f

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gregoryv/fox"
)

func NewTerm() *Term {
	m := Term{
		Logger: fox.NewSyncLog(os.Stderr).FilterEmpty(),
		exit:   os.Exit,
	}
	m.SetOutput(os.Stdout)
	m.errFuncs = []liner{
		func(s *string) {
			var cmd exec.Cmd
			OpenError(&cmd, *s, m.wd)
			RunCmd(&cmd)
		},

		func(s *string) { Color(s, m.wd) },
		func(s *string) { Strip(s, m.wd) },
		func(s *string) { Color(s, "_test.go") },
	}
	m.okFuncs = []liner{}
	dir, _ := os.Getwd()
	m.wd = dir + "/"
	return &m
}

type Term struct {
	fox.Logger
	output   io.Writer
	Verbose  bool
	exit     func(int)
	wd       string
	errFuncs []liner
	okFuncs  []liner
}

func (m *Term) SetOutput(w io.Writer) {
	m.output = w
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
		fmt.Fprintln(m.output, s)
	}
}
