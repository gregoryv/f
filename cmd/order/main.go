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
	var filename string
	flag.StringVar(&filename, "f", "", "order file")
	flag.Parse()

	order, err := ioutil.ReadFile(filename)
	if err != nil {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	patterns := strings.Split(string(order), "\n")
	var content bytes.Buffer
	io.Copy(&content, os.Stdin)
	body := bytes.TrimSpace(content.Bytes())
	lines := strings.Split(string(body), "\n")

	sort.Sort(ByPattern{
		lines:    lines,
		patterns: patterns,
	})
	for _, line := range lines {
		fmt.Println(line)
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
