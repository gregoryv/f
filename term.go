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
	m.wd = dir
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
			StripAndColor(&s, t.wd)
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

func StripAndColor(line *string, contains string) error {
	mycode := strings.Index(*line, contains) > -1
	stripped := strings.ReplaceAll(*line, contains, "")
	if !mycode {
		return notColored
	}
	colored := red + stripped + reset
	line = &colored
	return nil
}

var (
	red        = "\033[31m"
	reset      = "\033[0m"
	notColored = fmt.Errorf("not colored")
)
