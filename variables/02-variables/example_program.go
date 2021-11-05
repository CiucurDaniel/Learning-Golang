package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", input)

	output := input * 2.0

	fmt.Println(fmt.Sprintf("I multiplied your number by 2, here is the result: %f", output))

}
