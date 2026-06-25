/*

**
In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

var declares 1 or more variables.

You can declare multiple variables at once.

Go will infer the type of initialized variables.

Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0

The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
**

*/

package main

import "fmt"

func main() {
	// Declare a variable of type int
	var x int
	fmt.Println("The value of x is:", x) // Output: The value of x is: 0

	// Declare and initialize a variable of type string
	var name string = "Alice"
	fmt.Println("The name is:", name) // Output: The name is: Alice

	// Declare multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println("Values are:", a, b, c) // Output: Values are: 1 2 3

	// Type inference with initialization
	var d = 4.5                          // d is inferred to be of type float64
	fmt.Println("The value of d is:", d) // Output: The value of d is: 4.5

	// Short variable declaration (only inside functions)
	e := "Hello, Messi Won the last world cup!" // e is inferred to be of type string
	fmt.Println("The value of e is:", e)        // Output: The value of e is: Hello, Messi Won the last world cup!
}
