package main

import "fmt"

/* ENUMS
Enumerated types (enums) are a special case of sum types (https://en.wikipedia.org/wiki/Algebraic_data_type)
Enum is a type that has a fixed number of possible values, each with a distinct name.
Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
- 


https://pkg.go.dev/golang.org/x/tools/cmd/stringer
https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate
*/

// TYPE refers to classification of values and variables, specifying the kind of data they hold and the operations that can be performed on them.
type ServerState int  // our enum type ServerState has an underying int type


const (  // possible values for ServerState are defined as constants
    StateIdle = iota  // special keyword iota generated successive constant values automatically; in this case 0, 1, 2..
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{  // by implementing fmt.Stringer interface, vales of ServerState can be printed out or converted to strings
    StateIdle:      "idle",  
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
}

func enums() {  // 
    ns := transition(StateIdle)
    fmt.Println(ns)
    //

    ns2 := transition(ns)
    fmt.Println(ns2)
    //
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
