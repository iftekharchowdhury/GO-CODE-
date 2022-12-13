/*
Integer to string conversation - use strconv.Itoa()
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 53
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)

}
