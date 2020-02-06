// Programmer name: Matthew Davenport
// Date completed:  2/6/20
// Description: This program will as the user for their name and age
              //then display back a welcome message of their name, age and
              //my favorite food.
package main

import "fmt" 

func main() {
    //declare variable for favorite food and store your favorite food.
    var favFood string = "Macaroni and Cheese"

    //declare variables for name and age (make sure they are appropriate data types)
    var name string
    var age int

    //ask the user to enter their answer for name and age.
    fmt.Println("Please enter your name")
    fmt.Scanln(&name)
    fmt.Println("Please enter your age")
    fmt.Scanln(&age)

    //output a welcome statement using their name
    fmt.Println("Welcome ", name, "!")

    //output a statement that says At their age you enjoyed the favorite food
    fmt.Println("At ", age, "years old, my favorite food was ", favFood)
}