package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/* FILE PATHS
path/filepath package provides utilities to work with file and directory paths across different operating systems
- Ensures cross-platform compability when handling file paths

- Common functions:
filepath.Join(elements...) - combines multiple path elements into a single path
filepath.Abs(path) - returns the absolute path of a relative path
filepath.Base(path) - returns the last element of the path (e.g. filename)
*/

func file_paths() {

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	
}
