package f

import (
	"testing"
)

func Test_tools(t *testing.T) {
	m := NewTerm()
	ok(t, TidyImports(m, &Args{Ext: ".go", Path: "package_test.go"}))
	bad(t, TidyImports(m, &Args{Ext: ".txt"}))
}

func Test_EmacsOpen(t *testing.T) {
	var cli string
	ok(t, EmacsOpen(&cli, "/file:10"))
	if cli == "" {
		t.Fail()
	}
	ok(t, EmacsOpen(&cli, "   file_test.go:10: error"))
	if cli != "emacsclient -n +10 file_test.go" {
		t.Error(cli)
	}

	bad(t, EmacsOpen(&cli, "/path/file 10"))
	bad(t, EmacsOpen(&cli, "--- PASS: TestColor (0.00s)"))
}
