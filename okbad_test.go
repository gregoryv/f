package f

import (
	"strings"
	"testing"
)

func ok(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err != nil {
		if len(msg) > 0 {
			t.Error(strings.Join(msg, " ")+":", err)
			return
		}
		t.Error(err)
	}
}

func bad(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err == nil {
		if len(msg) > 0 {
			t.Error(strings.Join(msg, " "), "should fail")
			return
		}
		t.Error("should fail")
	}
}
