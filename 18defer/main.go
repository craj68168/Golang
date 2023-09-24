package main

import "fmt"

func main() {
	// defer fmt.Println("World") // hello then world
	// fmt.Println("Hello")

	defer fmt.Println("One") // world three two one due to first in last out for defer
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("World")
	myDefer()
}

// World 4 3 2 1 0 Three Two One
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Index", i)
	}
}
