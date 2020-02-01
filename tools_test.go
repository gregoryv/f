package f

import "testing"

func Test_tools(t *testing.T) {
	m := NewTerm()
	ok(t, TidyImports(m, &Args{Ext: ".go"}))
	bad(t, TidyImports(m, &Args{Ext: ".txt"}))
}
