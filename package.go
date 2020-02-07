package f

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

func OpenError(cmd *exec.Cmd, line, wd string) error {
	var cli string
	err := Emacsclient(&cli, line)
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
		newCmd := exec.Command(c[0], c[1:]...)
		*cmd = *newCmd
	}
	return nil
}

func RunCmd(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Path == "" {
		return MissingCommand
	}
	return cmd.Run()
}

// Emacsclient parses v for file/path:LINENO and sets cli to open
func Emacsclient(cli *string, v string) error {
	v = strings.TrimSpace(v)
	first := strings.Split(v, " ")[0]
	parts := strings.Split(first, ":")
	if len(parts) > 1 {
		lineno := parts[1]
		path := parts[0]
		*cli = fmt.Sprintf("emacsclient -n +%s %s", lineno, path)
		return nil
	}
	return NotFound
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

func TidyImports(a *Args) error {
	if a.Ext != ".go" {
		return InvalidExtension
	}
	return Shf("goimports -w %s", a.Path)
}

/*
func FilterDurations(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Bytes()
		w.Write(word)
	}
	return nil
}
*/

var DefaultTerm = NewTerm()

func Sh(cli string) error {
	return DefaultTerm.Sh(cli)
}

func Shf(format string, args ...interface{}) error {
	return DefaultTerm.Shf(format, args...)
}

func NoExit() {
	DefaultTerm.SetExit(func(int) {})
}

func SetOutput(w io.Writer) {
	DefaultTerm.SetOutput(w)
}

func Verbose() {
	DefaultTerm.Verbose = true
}
