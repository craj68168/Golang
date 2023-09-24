package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://localhost:3000/learn?course=react&paymentid=1234"

func main() {

	result, _ := url.Parse(myUrl)
	// fmt.Println("Result scheme is", result.Scheme)     //https
	// fmt.Println("Result Host is", result.Host)         //localhost:3000
	// fmt.Println("Result Path is", result.Path)         // /learn
	// fmt.Println("Result Port is", result.Port())       //3000
	// fmt.Println("Result RawQuery is", result.RawQuery) //course=react&paymentid=1234

	qparams := result.Query()

	// fmt.Println("The query is", qparams["course"])    //[react]
	// fmt.Println("The query is", qparams["paymentid"]) //[1234]

	for _, v := range qparams {
		fmt.Println("The values are", v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "localhost:3000",
		Path:    "/learn",
		RawPath: "course=react&paymentid=1234",
	}
	myresult := partsOfUrl.String()
	fmt.Println("result", myresult) //result https://localhost:3000/learn
}
