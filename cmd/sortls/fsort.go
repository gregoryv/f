package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
)

func FsortDir(w io.Writer, args ...string) {
	cmd := exec.Command("ls", args...)
	order, err := ioutil.ReadFile(path.Join(args[len(args)-1], ".lsorder"))
	if err != nil {
		cmd.Stdout = w
		cmd.Run()
		return
	}
	content, _ := cmd.Output()
	content = bytes.TrimSpace(content)
	Fsort(w, string(order), string(content))
}

func Fsort(w io.Writer, patterns, lines string) {
	c := &cli{
		patterns: strings.Split(patterns, "\n"),
		lines:    strings.Split(lines, "\n"),
	}
	c.run(w)
}

type cli struct {
	patterns []string
	lines    []string
}

func (c *cli) run(w io.Writer) {
	for _, p := range c.patterns {
		c.lines = printMatching(p, c.lines)
		if len(c.lines) == 0 {
			break
		}
	}
	for _, l := range c.lines {
		fmt.Println(l)
	}
}

func printMatching(pattern string, lines []string) []string {
	rest := make([]string, 0)
	for _, line := range lines {
		if strings.Index(line, pattern) == -1 {
			// skip . and ..
			if len(line) > 2 {
				end := line[len(line)-2:]
				if end == " ." || end == ".." {
					fmt.Println(line)
					continue
				}
			}
			rest = append(rest, line)
			continue
		}
		fmt.Println(line)
	}
	return rest
}
