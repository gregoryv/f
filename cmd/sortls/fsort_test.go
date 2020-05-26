package main

import (
	"os"
)

func ExampleFsortDir() {
	FsortDir(os.Stdout, "./testdata")
	// output:
	// internal
	// README
	// changelog.md
	// file.txt
}

func ExampleFsort_simple_sort() {
	Fsort(os.Stdout, "intro", "file.txt\nintro\nb_intro")
	// output:
	// intro
	// b_intro
	// file.txt
}

func ExampleFsort_no_match() {
	Fsort(os.Stdout, "not", "file.txt\nintro\n")
	// output:
	// file.txt
	// intro
}

func ExampleFsort_colored() {
	Fsort(os.Stdout, "internal\nREADME", `README
[0m[01;34minternal[0m`)
	// output:
	// [0m[01;34minternal[0m
	// README
}

func ExampleFsort_mixed() {
	Fsort(os.Stdout, "internal\nREADME", `drwxrwxr-x 3 gregory gregory 4.0K May 22 08:42 .
drwxrwxr-x 4 gregory gregory 4.0K May 22 08:55 ..
drwxrwxr-x 2 gregory gregory 4.0K May 22 08:38 internal
-rw-rw-r-- 1 gregory gregory    0 May 22 08:27 changelog.md
-rw-rw-r-- 1 gregory gregory   37 May 22 08:42 .lsorder
-rw-rw-r-- 1 gregory gregory    0 May 22 08:27 main.go
-rw-rw-r-- 1 gregory gregory    0 May 22 08:27 README`)
	// output:
	// drwxrwxr-x 3 gregory gregory 4.0K May 22 08:42 .
	// drwxrwxr-x 4 gregory gregory 4.0K May 22 08:55 ..
	// drwxrwxr-x 2 gregory gregory 4.0K May 22 08:38 internal
	// -rw-rw-r-- 1 gregory gregory    0 May 22 08:27 README
	// -rw-rw-r-- 1 gregory gregory    0 May 22 08:27 changelog.md
	// -rw-rw-r-- 1 gregory gregory   37 May 22 08:42 .lsorder
	// -rw-rw-r-- 1 gregory gregory    0 May 22 08:27 main.go
}
