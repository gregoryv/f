package f

import (
	"testing"
)

func TestNewArgs(t *testing.T) {
	var f string
	ok, _k := assert(t)
	ok(NewArgs(".").Format(&f))
	_k(NewArgs(".", "no such command").Format(&f))
}
