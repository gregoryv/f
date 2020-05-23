package fo

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

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
	wd, _ := os.Getwd()
	ok(OpenError(&cmd, "package_test.go:37: message...", wd))
	// to verify
	// cmd.Run()
	_k(OpenError(&cmd, "", ""))
}

func TestTidyImports(t *testing.T) {
	ok, _k := assert(t)
	ok(TidyImports("package_test.go"))
	_k(TidyImports())
	_k(TidyImports("file.txt"))
}

func TestDefaultTerm(t *testing.T) {
	SetOutput(ioutil.Discard)
	Verbose()
	NoExit()
	DefaultTerm.Verbose = false //so below output doesn't show
	Sh("whohooo ")
	Sh("touch package.go")
	Shf("%s %s", "touch", "package_test.go")
}
