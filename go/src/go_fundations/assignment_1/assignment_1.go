// ask user for First Name
// ask user for Last Name
// your name is: Juan Rojas
// Hint - Read Printf in fmt package
// use fmt.Scan() or fmt.Scanln() for reading user input

package main

import "fmt"

func main() {
	var firstname string
	var lastname string

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Printf("Your name is: %s %s", firstname, lastname)
}
