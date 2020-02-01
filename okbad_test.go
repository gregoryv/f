package f

import (
	"strings"
	"testing"
)

func ok(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err != nil {
		t.Error(strings.Join(msg, " ")+":", err)
	}
}

func bad(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err == nil {
		t.Error(strings.Join(msg, " "), "should fail")
	}
}
