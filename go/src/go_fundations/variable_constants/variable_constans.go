package main

import "fmt"

func main() {
	//variable - keyboard
	var username string
	username = "Australia"
	fmt.Println("Username: ", username)
	firstName := "Juan"
	lastName := "Rojas"
	fmt.Println("Fullname:", firstName, lastName)
	var (
		a string = "tic"
		b int    = 10
		c bool   = true
	)

	fmt.Println(a, b, c)
	var d, e, f = "I am a string", 9.0, true
	fmt.Printf("d: %v, e: %.1f, f: %v\n", d, e, f)

	d = "Is now a new string"
	fmt.Println("d:", d)

	const API_KEY = "sadf.sadfsdf.555"
	fmt.Println(API_KEY)

}
