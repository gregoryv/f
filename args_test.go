package f

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewArgs(t *testing.T) {
	ok(t, action(".", "ls"))
	ok(t, action(". f", "f"))
	ok(t, format("."))
	bad(t, format(". ljlj"))
}

func format(in string) error {
	a := NewArgs(strings.Split(in, " "))
	_, found := a.Format()
	if !found {
		return fmt.Errorf("missing format: %s %#v", in, a)
	}
	return nil
}

func action(in string, exp string) error {
	a := NewArgs(strings.Split(in, " "))
	got := a.action
	if got != exp {
		return fmt.Errorf("%q: got %q, exp %q", in, got, exp)
	}
	return nil
}

func ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}

func bad(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("should fail")
	}
}
