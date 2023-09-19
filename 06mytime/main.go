package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println("The Current time is", presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // its a format must use to get current date in format
	fmt.Println(presentTime.Format("01/02/2006")) // 09/19/2023

	createDate := time.Date(2023, time.September, 20, 23, 23, 0, 0, time.UTC)

	fmt.Println("", createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday")) //09-20-2023 Wednesday
}
