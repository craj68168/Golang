package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["GL"] = "Golang"
	fmt.Println("Languages", languages) //map[GL:Golang RB:Ruby js:JavaScript] no comma seperated
	fmt.Println("JS", languages["js"])  //JavaScript
	delete(languages, "RB")
	fmt.Println("Deleted", languages) //map[GL:Golang js:JavaScript]

	//loops in golang

	for key, value := range languages {
		fmt.Printf("For keys %v, value is %v\n", key, value) //For keys js, value is JavaScript // For keys GL, value is Golang
	}
}
