package f

import (
	"testing"
)

func Test_tools(t *testing.T) {
	m := NewTerm()
	ok(t, TidyImports(m, &Args{Ext: ".go"}))
	bad(t, TidyImports(m, &Args{Ext: ".txt"}))
}

func Test_EmacsOpen(t *testing.T) {
	var cli string
	ok(t, EmacsOpen(&cli, "/file:10"))
	if cli == "" {
		t.Fail()
	}
	bad(t, EmacsOpen(&cli, "/path/file 10"))
}
