package main

import (
	"fmt"
	"os"
	"path/filepath"
)



func temporary_files_and_directories() {

	f, err := os.CreateTemp("", "sample")
	check_tfd(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check_tfd(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check_tfd(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check_tfd(err)

}

func check_tfd(e error) {
	if e != nil {
		panic(e)
	}
}
