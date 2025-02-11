package main
import "fmt"

/* 
- GO infer type if not explicitly states
- shorthand syntax inside functions ':='
- declaration of multiple variables at once:
		var x, y int = 1, 2
		x, y := 1, 2
- variable names starting with a capital letter are exported and accessible from other packages;
	lowercase ones are package-private
*/

func variables() {

	var a = "initial" // var declares 1 or more variables (here without stating type)
	fmt.Println(a)

	var b, c int = 1, 2 // you can declare multiple variables at once
	fmt.Println(b, c)

	var d = true // go will infer the type of initialized variables
	fmt.Println(d)

	var e int // variables declared without initilization are zero-valued
	fmt.Println(e)

	f := "apple"   // := syntax is shorthand for declaring and initializing a variable
	fmt.Println(f) // this syntax is only available inside functions
}


/* 
ZERO VALUES
int : 0
float : 0.0
complex: 0 + 0i
bool : false
string : ""
*T (any pointer type) : nil
slice : nil
map : nil
channel : nil
interface : nil
struct : each field in a struct gets it's corresponding zero value
*/

func values() {
	fmt.Println("go" + "lang")
	//golang

	fmt.Println("1+1 =", 1+1)
	//1+1 = 2

	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// 7.0/3.09 = 2.3333333335

	fmt.Println(true && false)
	//false

	fmt.Println(true || false)
	//true

	fmt.Println(!true)
	//false
}
