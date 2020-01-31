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
	dir, _ := os.Getwd()
	m.wd = dir + "/"
	return &m

}

type Term struct {
	fox.Logger
	Verbose bool
	exit    func(int)
	wd      string
}

func (t *Term) Log(p ...interface{}) {
	if t.Verbose {
		t.Logger.Log(p...)
	}
}

func (t *Term) Shf(format string, args ...interface{}) {
	t.Sh(fmt.Sprintf(format, args...))
}

func (t *Term) Sh(cli string) {
	start := time.Now()
	t.Log("# ", cli)

	p := strings.Split(cli, " ")
	out, err := exec.Command(p[0], p[1:]...).CombinedOutput()
	if err != nil {
		lines := bytes.Split(out, []byte("\n"))
		for _, line := range lines {
			s := string(line)
			Color(&s, t.wd)
			Strip(&s, t.wd)
			Color(&s, "_test.go")
			fmt.Println(s)
		}
		fmt.Println(err)
		t.exit(1)
	}
	nice := bytes.TrimSpace(out)
	if len(nice) > 0 {
		fmt.Println(string(nice))
	}
	t.Log("# ", cli, " ", time.Since(start))
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
