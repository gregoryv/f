package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func Test_cli(t *testing.T) {
	var out bytes.Buffer
	tmp, _ := ioutil.TempFile("", "order")
	exp := `internal
README
changelog.md
file.txt
`
	tmp.WriteString(exp)
	tmp.Close()
	defer os.RemoveAll(tmp.Name())
	c := &cli{
		Writer:   &out,
		Reader:   strings.NewReader("internal\nREADME\nchangelog.md\nfile.txt\n"),
		filename: tmp.Name(),
	}
	c.run()
	if out.String() != exp {
		t.Error(out.String())
	}
}

func TestFilterLines(t *testing.T) {
	lines := []string{"a", "b", "a", "c", "d"}
	sort.Sort(ByPattern{
		lines:    lines,
		patterns: []string{"a", "c"},
	})
	exp := []string{"a", "a", "c", "b", "d"}
	if !reflect.DeepEqual(lines, exp) {
		t.Error(lines)
	}
}
