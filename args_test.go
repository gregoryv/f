package f

import (
	"testing"
)

func TestNewArgs(t *testing.T) {
	var f string
	ok, _k := assert(t)
	ok(NewArgs([]string{"."}).Format(&f))
	_k(NewArgs([]string{".", "no such command"}).Format(&f))
}
