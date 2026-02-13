package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time studies")
	presentTime := time.Now()
	fmt.Println("Present time is : ", presentTime)
	fmt.Println("Present time in format : ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.February, 16, 12, 0, 0, 0, time.UTC)
	fmt.Println("Created date is : ", createdDate)
}

//created CLI executable using go build command in terminal