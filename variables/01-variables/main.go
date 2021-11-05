package main

import "fmt"

func main() {

	//simple way of declaring a variable and later assign it with a value
	var x string
	x = "My first variable"
	fmt.Println(x)

	// Go also has type inference so you can do

	z := "Type inference"
	fmt.Println(z)

	// constants
	// they are created the same as variables but instead use "const" instead of var

	const y = "Const string"
	// y = "I try to change the string"  -> ERROR

	// Go also has a syntax for defining multiple variables at once

	var (
		a  = 14
		hi = "Hi"
		b  = 10
	)

	fmt.Println(a, hi, b)
}
