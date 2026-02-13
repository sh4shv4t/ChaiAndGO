package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is :", len(fruitList))
}
