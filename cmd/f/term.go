package main

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
		/*		func(s *string) {
				var cmd exec.Cmd
				OpenError(&cmd, *s, m.wd)
				RunCmd(&cmd)
			},*/

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

func (m *Term) SetOutput(w io.Writer) { m.output = w }
func (m *Term) SetExit(fn func(int))  { m.exit = fn }

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
	m.Log("# ", cli, " ", time.Since(start).Round(time.Millisecond))
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

// ----------------------------------------

var DefaultTerm = NewTerm()

func NoExit()               { DefaultTerm.SetExit(func(int) {}) }
func SetOutput(w io.Writer) { DefaultTerm.SetOutput(w) }
func Verbose()              { DefaultTerm.Verbose = true }
func Sh(cli string) error   { return DefaultTerm.Sh(cli) }

func Shf(format string, args ...interface{}) error {
	return DefaultTerm.Shf(format, args...)
}

// ----------------------------------------

func RunCmd(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Path == "" {
		return MissingCommand
	}
	return cmd.Run()
}

func Color(line *string, contains string) error {
	found := (strings.Index(*line, contains) > -1)
	if !found {
		return Unchanged
	}
	*line = red + *line + reset
	return nil
}

func Strip(line *string, part string) error {
	stripped := strings.ReplaceAll(*line, part, "")
	if stripped == *line {
		return Unchanged
	}
	*line = stripped
	return nil
}

var (
	red   = "\033[31m"
	reset = "\033[0m"
)

var (
	Unchanged        = fmt.Errorf("unchanged")
	InvalidExtension = fmt.Errorf("invalid extension")
	NotFound         = fmt.Errorf("not found")
	MissingCommand   = fmt.Errorf("missing command")
)
