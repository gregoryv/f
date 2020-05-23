package main

import "testing"

func TestTidyImports(t *testing.T) {
	ok, _k := assert(t)
	ok(TidyImports("main.go"))
	_k(TidyImports())
	_k(TidyImports("file.txt"))
}
