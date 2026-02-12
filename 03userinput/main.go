package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter rating : \n")

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us : ", input)
	fmt.Printf("Variable is of type : %T \n", input)
}