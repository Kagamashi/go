package main

import "os"

/* PANIC
Panic is used to immediately stop the execution of the current function, unwind the stack, terminate the program - unless recovered.
It's generally for unrecoverable errors or critical failures.
- panic("Something fked")
- panic should not be used for regular error handling, but reserved for fatal errors
*/

func panic_er() {

    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
