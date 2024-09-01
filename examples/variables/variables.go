package main

import "fmt"

func main() {
	//VALUES
	fmt.Println("go" + "lang")        //golang
	fmt.Println("1+1 =", 1+1)         //1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) // 7.0/3.09 = 2.3333333335
	fmt.Println(true && false)        //false
	fmt.Println(true || false)        //true
	fmt.Println(!true)                //false

	//VARIABLES
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}

