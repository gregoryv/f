package f

import (
	"strings"
	"testing"
)

func assertOk(t *testing.T) assertFunc {
	return func(err error, msg ...string) {
		t.Helper()
		if err != nil {
			if len(msg) > 0 {
				t.Error(strings.Join(msg, " ")+":", err)
				return
			}
			t.Error(err)
		}
	}
}

func assertBad(t *testing.T) assertFunc {
	return func(err error, msg ...string) {
		t.Helper()
		if err == nil {
			if len(msg) > 0 {
				t.Error(strings.Join(msg, " "), "should fail")
				return
			}
			t.Error("should fail")
		}
	}
}

func assert(t *testing.T) (ok, bad assertFunc) {
	return assertOk(t), assertBad(t)
}

type assertFunc func(error, ...string)
