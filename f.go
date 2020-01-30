package f

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type F struct {
	Verbose bool
}

func (f *F) Shf(format string, args ...interface{}) {
	f.Sh(fmt.Sprintf(format, args...))
}

func (f *F) Sh(cli string) {
	start := time.Now()
	if f.Verbose {
		fmt.Println("# ", cli)
	}
	p := strings.Split(cli, " ")
	out, err := exec.Command(p[0], p[1:]...).CombinedOutput()
	if err != nil {
		lines := bytes.Split(out, []byte("\n"))
		for _, line := range lines {
			fmt.Println(ColorWorkingDir(line))
		}
		fmt.Println(err)
		os.Exit(1)
	}
	nice := bytes.TrimSpace(out)
	if len(nice) > 0 {
		fmt.Println(string(nice))
	}
	fmt.Println("# ", cli, time.Since(start))
}

func ColorWorkingDir(line []byte) string {
	dir, err := os.Getwd()
	if err != nil {
		return string(line)
	}
	s := string(line)
	mycode := strings.Index(s, dir) > -1
	s = strings.ReplaceAll(s, dir, "")
	if !mycode {
		return s
	}
	return red + s + reset
}

var (
	red   = "\033[31m"
	reset = "\033[0m"
)
