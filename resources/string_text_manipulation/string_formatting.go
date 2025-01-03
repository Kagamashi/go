package main

import (
	"fmt"
	"os"
)

/*
%v print the values (default format)
	%s string
	%f float
	%d integer
%c print character represented by rune
%U print Unicode code point
%+v include struct field names
%T print type of the value
%x provides HEX encoding
%q to double quote strings
%p print representation of a pointer
*/

type point struct {
	x, y int
}

func string_formatting() { // main

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	// struct1: {1 2}

	fmt.Printf("struct2: %+v\n", p) // %+v include struct field name
	// struct2: {x:1, y:2}

	fmt.Printf("struct3: %#v\n", p) // prints Go syntax representation of the value i.e. source code snipper that would prouduce that value
	// struct3: main.point{x:1, y:2}

	fmt.Printf("type: %T\n", p) // print type of the value
	// type: main.point

	fmt.Printf("bool: %t\n", true) // formatting booleans is straightforward
	// bool: true

	fmt.Printf("int: %d\n", 123)
	// int: 132

	fmt.Printf("bin: %b\n", 14) // prints binary representation
	// bin: 1110

	fmt.Printf("char: %c\n", 33) // prints character corresponding to the given integer
	// char: !

	fmt.Printf("hex: %x\n", 456)
	// hex: 1c8

	fmt.Printf("float1: %f\n", 78.9)
	// float1: 78.900000

	fmt.Printf("float2: %e\n", 123400000.0)
	// float2: 1.234000e+08

	fmt.Printf("float3: %E\n", 123400000.0)
	// float3: 1.234000E+08

	fmt.Printf("str1: %s\n", "\"string\"") // basic print string
	// str1: "string"

	fmt.Printf("str2: %q\n", "\"string\"") // double-quote strings
	// str2: "\"string\""

	fmt.Printf("str3: %x\n", "hex this")
	// str3: 6865782074686973

	fmt.Printf("pointer: %p\n", &p)
	// pointer: 0xc0000ba000

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	// width1: |    12|   345|

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// width2: |  1.20|  3.45|

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// width3: |1.20  |3.45  |

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// width4: |   foo|     b|

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	// width5: |foo   |b     |

	s := fmt.Sprintf("sprintf: a %s", "string") // Printf prints formatted string to os.Stdout
	fmt.Println(s)                              //fmt.Sprintf format and returns a string without printing it anywhere
	// sprintf: a string

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // format+print to io.Writes other than os.Stdout using Fprintf
	// io: an error

}
