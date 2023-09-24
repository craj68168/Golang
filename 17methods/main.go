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
	personalDetail.GetStatus()
	personalDetail.GetEmail() // razz@gmail.com
	fmt.Printf("Personal details: %+v\n ", personalDetail)
}

func (u User) GetStatus() {
	fmt.Println("The status is", u.status)
}

func (u User) GetEmail() {
	u.email = "razz@gmail.com" // i wont change the origin value for that use pointer
	fmt.Println("The email is", u.email)
}
