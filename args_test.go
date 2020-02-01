package f

import (
	"testing"
)

func TestNewArgs(t *testing.T) {
	ok, _k := assert(t)

	var fn Action
	ok(NewArgs(".").UseAction(&fn))
	fn(NewTerm())

	_k(NewArgs(".", "oups").UseAction(&fn))
}
