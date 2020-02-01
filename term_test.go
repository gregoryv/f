package f

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/fox"
)

func TestTerm(t *testing.T) {
	m := NewTerm()
	ok(t, loggerSet(m))
	ok(t, silentLog(m))
	ok(t, TidyImports(m, NewArgs([]string{"term_test.go"})))

	m.Verbose = true
	bad(t, silentLog(m))
	bad(t, TidyImports(m, NewArgs([]string{"term_test.txt"})))
	m.Shf("%s", "touch term_test.go")
	bad(t, m.Shf("%s", "hubladuble"))
	// output is trimmed
	m.Sh("echo '  hello '")
}

func TestColor(t *testing.T) {
	line := "/home/john"
	ok(t, Color(&line, "/home"))
	bad(t, Color(&line, "/etc"))
}

func TestStrip(t *testing.T) {
	line := "/home/john"
	ok(t, Strip(&line, "/home"))
	line2 := "/home/john"
	bad(t, Strip(&line2, "/etc"))
}

func loggerSet(m *Term) error {
	if m.Logger == nil {
		return fmt.Errorf("Logger is nil")
	}
	return nil
}

func silentLog(m *Term) error {
	var buf bytes.Buffer
	l := fox.NewSyncLog(&buf)
	m.Logger = l
	m.Log("x")
	got := buf.String()
	if got != "" {
		return fmt.Errorf("Default Verbose should be silent")
	}
	return nil
}
