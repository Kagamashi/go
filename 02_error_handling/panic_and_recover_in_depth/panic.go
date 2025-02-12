package main
import "os"

/*
Panic is used to immediately stop the execution of the current function, unwind the stack, terminate the program - unless recovered.
used for unrecoverable errors or critical failures.
- panic should not be used for regular error handling, but reserved for fatal errors
*/

func panic_er() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	panic("a problem")
}
