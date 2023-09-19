package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fullName := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your Rating:")

	input, err := fullName.ReadString('\n') // \n means hitting enter

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your rating is", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // parse string to flat

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("num rating", numRating+1)
}
