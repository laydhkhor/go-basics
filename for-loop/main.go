/*

**


for is Go’s only looping construct. Here are some basic types of for loops.
The most basic type, with a single condition.
A classic initial/condition/after for loop.
Another way of accomplishing the basic “do this N times” iteration is range over an integer.
for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
You can also continue to the next iteration of the loop.


**



*/

package main

import "fmt"

func main() {
	// Basic for loop with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic for loop with init, condition, and post statements
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Using range to iterate over an array
	for _, value := range []int{10, 11, 12} {
		fmt.Println(value)
	}

	// Infinite loop with break
	for {
		fmt.Println("loop")
		break
	}

	// Using continue to skip an iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(n) // Print only odd numbers
	}

	// Using range to iterate over a slice and skip even numbers

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
