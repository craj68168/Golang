package main

import "fmt"

type User struct {
	name   string
	email  string
	phone  int
	status bool
}

func main() {
	personalDetail := User{"Raj", "raj@gmail.com", 9824223239, true}
	fmt.Println("Personal details is", personalDetail)     //{Raj raj@gmail.com 9824223239 true}
	fmt.Printf("Personal details: %+v\n ", personalDetail) //{name:Raj email:raj@gmail.com phone:9824223239 status:true}

}
