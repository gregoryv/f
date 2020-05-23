package fo

import (
	"io/ioutil"
	"testing"
)

func TestNewArgs(t *testing.T) {
	ok, _k := assert(t)

	var fn Action
	ok(NewArgs(".").UseAction(&fn))
	m := NewTerm()
	m.SetOutput(ioutil.Discard)
	fn(m)

	_k(NewArgs(".", "oups").UseAction(&fn))
}
