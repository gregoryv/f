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

	m.Verbose = true
	bad(t, silentLog(m))

	m.Shf("%s", "touch term_test.go")
	bad(t, unknownCommand(m))
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

func unknownCommand(m *Term) error {
	var failed bool
	m.exit = func(x int) { failed = true }
	m.Shf("%s", "hubladuble")
	if failed {
		return fmt.Errorf("did not fail when executing unknown command")
	}
	return nil
}
