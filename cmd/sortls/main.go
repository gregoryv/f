package main

import "os"

func main() {
	FsortDir(os.Stdout, os.Args[1:]...)
}
