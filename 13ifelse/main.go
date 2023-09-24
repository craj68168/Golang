package main

import "fmt"

func main() {

	count := 21
	var age int

	if count < 10 {
		age = 10
		fmt.Println("Age is", age)
	} else if count > 20 {
		age = 25
		fmt.Println("Age is", age)
	} else {
		age = 0
		fmt.Println("Age is", age)
	}

	if num := 5; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater")
	}
}
