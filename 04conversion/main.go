package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rating := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your rating:")

	input, err := rating.ReadString('\n') // \n means hitting enter
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your rating is", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // parse string to flat
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Num Rating", numRating+1)
}
