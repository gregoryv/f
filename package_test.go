package f

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

func Test_tools(t *testing.T) {
	m := NewTerm()
	ok(t, TidyImports(m, &Args{Ext: ".go", Path: "package_test.go"}))
	bad(t, TidyImports(m, &Args{Ext: ".txt"}))
}

func Test_EmacsClient(t *testing.T) {
	var cli string
	ok(t, Emacsclient(&cli, "/file:10"))
	if cli == "" {
		t.Fail()
	}
	ok(t, Emacsclient(&cli, "   file_test.go:10: error"))
	if cli != "emacsclient -n +10 file_test.go" {
		t.Error(cli)
	}

	bad(t, Emacsclient(&cli, "/path/file 10"))
	bad(t, Emacsclient(&cli, "--- PASS: TestColor (0.00s)"))
}

func TestRunCmd(t *testing.T) {
	ok(t, RunCmd(exec.Command("echo")))
}

func Test_OpenError(t *testing.T) {
	var cmd exec.Cmd
	wd, _ := os.Getwd()
	line := path.Join(wd, "package_test.go:10")
	ok(t, OpenError(&cmd, line, wd))
	bad(t, OpenError(&cmd, "", wd))
}
