// Command order sorts lines on stdin according to patterns in the
// order file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	c := &cli{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	flag.StringVar(&c.filename, "f", "", "order file")
	flag.Parse()
	c.run()
}

type cli struct {
	io.Writer
	io.Reader
	filename string
}

func (c *cli) run() {
	order, err := ioutil.ReadFile(c.filename)
	if err != nil {
		// no order file
		io.Copy(c.Writer, c.Reader)
	}
	// each line in the order file is a pattern
	patterns := strings.Split(string(order), "\n")

	// read stdin as lines
	var content bytes.Buffer
	io.Copy(&content, c.Reader)
	body := bytes.TrimSpace(content.Bytes())
	lines := strings.Split(string(body), "\n")

	sort.Sort(ByPattern{
		lines:    lines,
		patterns: patterns,
	})
	for _, line := range lines {
		fmt.Fprintln(c.Writer, line)
	}
}

type ByPattern struct {
	lines    []string
	patterns []string
}

func (b ByPattern) Less(i, j int) bool {
	return b.patternIndex(i) < b.patternIndex(j)
}

func (b ByPattern) patternIndex(lineIndex int) int {
	for i, pattern := range b.patterns {
		if strings.Index(b.lines[lineIndex], pattern) > -1 {
			return i
		}
	}
	return len(b.patterns)
}

func (b ByPattern) Len() int { return len(b.lines) }

func (b ByPattern) Swap(i, j int) {
	b.lines[i], b.lines[j] = b.lines[j], b.lines[i]
}
