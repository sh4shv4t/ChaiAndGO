package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Mango", "Grapes"}

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is :", len(fruitList))

	fruitList = append(fruitList, "Banana", "Papaya")

	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list is :", fruitList)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 999
	highscore[2] = 456
	highscore[3] = 567

	highscore = append(highscore, 678, 789, 890)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

	//removing values from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
