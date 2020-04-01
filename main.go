// Programmer name: Matthew Davenport
// Date completed:  4/1/20
// Description: This program will ask the user for a name then display back a welcome message

package main

import "fmt"

//create a function that accepts a name and prints a greeting message using the name
func welcome(name string){
  fmt.Println("Welcome", name, "!")
}

func main() {
//call your function
var input string
fmt.Println("Please enter your name")
fmt.Scanln(&input)
welcome(input)
}