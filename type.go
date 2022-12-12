package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("prinkya"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
	var grades int = 42
	var message string = "hello world"
	var grades_C = 89

	fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))

	fmt.Printf("variable grades=%v is of type %v \n", message, reflect.TypeOf(message))
	fmt.Println("grade c type is ", reflect.TypeOf(grades_C))
}
