package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Printf("Value of pointer is : %v \n", ptr)

	mynum:=23
	var ptr = &mynum

	fmt.Printf("Value of pointer is : %v \n", ptr)
	fmt.Printf("Value of pointer's reference is : %v \n", *ptr)

	*ptr = *ptr + 1
	fmt.Printf("Value of pointer's reference after increment is : %v \n", mynum)
}
