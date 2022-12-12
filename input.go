/*
This program describes how to take input from answer.
use scanf
*/

package main

import "fmt"

func main() {

	var name, surname string
	var amount int
	var is_muggle bool
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &surname)
	fmt.Print("are you a muggle: ")
	fmt.Scanf("%t", &is_muggle)
	fmt.Println("Hey There, ", name, "is muggle ", is_muggle)
	fmt.Println("Hey There, ", surname)
	fmt.Print("Enter amount: ")
	fmt.Scanf("%s", &amount)
	// fmt.Println("your amount: ", amount)
	fmt.Scanf("%s", &surname)

}
