package fo

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

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
		// don't use Sh, recursive errors are bad
		newCmd := exec.Command(c[0], c[1:]...)
		*cmd = *newCmd
	}
	return nil
}
