package main

import "fmt"

func main() {

	num := myFunc(4, 5)

	fmt.Println("Number is", num)

	returnVal, message := proFunc(4, 2, 10, 14)
	fmt.Println("Total return", returnVal) //30
	fmt.Println("My Message ", message)
}

func myFunc(vala int, valb int) int { //9
	return vala + valb
}

func proFunc(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "Hello"
}
