package f

import (
	"fmt"
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
	red   = "\033[31m"
	reset = "\033[0m"
)

var (
	notColored       = fmt.Errorf("not colored")
	notStripped      = fmt.Errorf("not stripped")
	InvalidExtension = fmt.Errorf("invalid extension")
	NotFound         = fmt.Errorf("not found")
	MissingCommand   = fmt.Errorf("missing command")
)

func TidyImports(m *Term, a *Args) error {
	if a.Ext != ".go" {
		return InvalidExtension
	}
	m.Shf("goimports -w %s", a.Path)
	return nil
}
