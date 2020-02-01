package f

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/fox"
)

func TestTerm(t *testing.T) {
	m := NewTerm()
	ok := assertOk(t)
	ok(loggerSet(m))
	ok(silentLog(m))

	m.Logger = t
	m.exit = func(int) {}
	m.Verbose = true
	bad := assertBad(t)
	bad(silentLog(m))
	m.Shf("%s", "touch term_test.go")
	bad(m.Shf("%s", "hubladuble"))
	// output is trimmed
	m.Sh("echo '  hello '")
}

func TestColor(t *testing.T) {
	line := "/home/john"
	ok := assertOk(t)
	ok(Color(&line, "/home"))

	bad := assertBad(t)
	bad(Color(&line, "/etc"))
}

func TestStrip(t *testing.T) {
	line := "/home/john"
	ok := assertOk(t)
	ok(Strip(&line, "/home"))

	line2 := "/home/john"
	bad := assertBad(t)
	bad(Strip(&line2, "/etc"))
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
