package main

import "fmt"

func main() {

	var totalScore = []int{20, 12, 43, 55}

	var index int = 2
	totalScore = append(totalScore[:index], totalScore[index+1:]...)

	fmt.Println("Remaining score are", totalScore) // [20 12 55]
}
