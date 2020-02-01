package f

import (
	"os/exec"
	"testing"
)

func Test_tools(t *testing.T) {
	m := NewTerm()
	ok, _k := assert(t)
	ok(TidyImports(m, &Args{Ext: ".go", Path: "package_test.go"}))
	_k(TidyImports(m, &Args{Ext: ".txt"}))
}

func Test_EmacsClient(t *testing.T) {
	var cli string
	ok, _k := assert(t)
	ok(Emacsclient(&cli, "/file:10"))
	ok(Emacsclient(&cli, "   file_test.go:10: error"))
	_k(Emacsclient(&cli, "/path/file 10"))
	_k(Emacsclient(&cli, "--- PASS: TestColor (0.00s)"))
}

func TestRunCmd(t *testing.T) {
	ok, _k := assert(t)
	ok(RunCmd(exec.Command("echo")))
	_k(RunCmd(exec.Command("")))
}

func Test_OpenError(t *testing.T) {
	var cmd exec.Cmd
	ok, _k := assert(t)
	ok(OpenError(&cmd, "package_test.go:10", ""))
	_k(OpenError(&cmd, "", ""))
}
