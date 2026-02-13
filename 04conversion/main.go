package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Please enter a rating ")

	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating us : ", val)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(val), 64)

	if err!=nil{
		fmt.Println("Error converting string to number : ", err)
		//panic(err)
	}else{
		fmt.Printf("New Rating %f \n", numRating+1)
	}


}
