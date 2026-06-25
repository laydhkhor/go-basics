/*

** - Go has various value types including strings, integers, floats, booleans, etc
** - Bellow are some examples of value types in Go:
** - We can use the "+" operator to concatenate strings, the "+" operator to add integers, the "/" operator to divide floats, and logical operators like "&&", "||", and "!" for boolean values.
 */

package main

import "fmt"

func main() {
	fmt.Println("Go Lang" + " " + "With Values")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
