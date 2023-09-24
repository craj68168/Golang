package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // generate random number between 0 - 5 and + 1 means 1 - 6
	fmt.Println("Random number is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Its Dice Number  1")
	case 2:
		fmt.Println("Its Dice Number  2")
	case 3:
		fmt.Println("Its Dice Number  3")
		fallthrough // if 3 than it fall one step down and also print below which is 4
	case 4:
		fmt.Println("Its Dice Number  4")
	case 5:
		fmt.Println("Its Dice Number  5")
	case 6:
		fmt.Println("Its Dice Number  6")
	default:
		fmt.Println("Opps what the number ??")
	}

}
