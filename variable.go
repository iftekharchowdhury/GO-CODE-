//import fmt and main package

package main

import "fmt"

// always declare main func

func main() {

	var x = 1
	var y = 2
	var z = 3
	// use x again
	x = 3

	//string
	var firstname = "xxx"
	var lastname = " islam"

	var fullname = firstname + lastname

	var name = "Harry"
	var user = "Kodekloud"

	fmt.Print("Welcome to the ", user, ", ", name, "\n")

	// float
	var a = 193.2939
	var n = 0.99
	/*
	   %v = default format
	   %d = int
	   %f = float
	   %.2f = floating numbers upto 2 decimal places
	   %t = true or false
	   %s = plain string
	   %q = quoted characters


	*/
	var grades = 42
	var grades2 = 82.3

	// inner & outer block
	// local variable and global variable
	var city = "London"
	{
		var country = "UK"
		fmt.Println(country)
		fmt.Println(city)

	}

	var s int
	var z float64

	fmt.Printf("Marks: %d", grades)
	fmt.Printf("Marks: %.2f", grades2)

	// print all the message
	fmt.Println(fullname)
	fmt.Println(x + y - z)
	fmt.Println(a + n - 19)
	fmt.Println(s)
	fmt.Printf(".2f", z)

}
