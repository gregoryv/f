package fo

import (
	"os"
	"os/exec"
	"testing"
)

func Test_EmacsClient(t *testing.T) {
	var cli string
	ok, bad := assert(t)
	ok(Emacsclient(&cli, "/file:10"))
	ok(Emacsclient(&cli, "   file_test.go:10: error"))

	bad(Emacsclient(&cli, "/path/file 10"))
	bad(Emacsclient(&cli, "--- PASS: TestColor (0.00s)"))
}

func Test_OpenError(t *testing.T) {
	var cmd exec.Cmd
	ok, bad := assert(t)
	ok(OpenError(&cmd, "term.go:10", ""))
	wd, _ := os.Getwd()
	ok(OpenError(&cmd, "term_test.go:37: message...", wd))
	// to verify
	// cmd.Run()
	bad(OpenError(&cmd, "", ""))
}
