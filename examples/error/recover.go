package main

import "fmt"

/* RECOVER
Recover is used to regain control of a panicking goroutine, effectively stopping the panic and preventing the program from crashing.

- Recover only works inside deferred functions - if called during a panic, it catches the panic and allows the program to continue
- If there's no panic recover returns nil
- Used for graceful error recovery, especially in servers or critical systems where abrupt termination must be avoided
*/

func mayPanic() {
    panic("a problem")
}

func recover_er() {

    defer func() {
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
}