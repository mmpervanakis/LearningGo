package main

import "fmt"

func main() {

	/*
		var x = 1
		var boolean = true
		var integer = 1023493
		var str = "my string"
		var float = 12.4
		var character = 'A'

		fmt.Print(x, boolean, integer, str, float, character)

	*/
	var input string
	fmt.Print("Your input here man:")
	n, err := fmt.Scan(&input)
	fmt.Println(input, "Nom of inputs", n, err)
}
