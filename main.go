// Programmer name: Matthew Davenport
// Date completed:  5/1/20
// Description: This program will create a struct for an animal and then create 2 animal objects using the struct

package main

import "fmt"

//Create a struct that keeps track of animal and stores its name, age, breakfast hour, and dinner hour.
type Animal struct {
  name string
  age int
  breakfastHour int 
  dinnerHour int
}

func main() {
   //create 2 animal objects that store the appropriate data and then print out the data stored
   animalA := Animal {}
   animalA.name = "Princess"
   animalA.age = 3
   animalA.breakfastHour = 8
   animalA.dinnerHour = 5

   animalB := Animal {}
   animalB.name = "Tony"
   animalB.age = 6
   animalB.breakfastHour = 6
   animalB.dinnerHour = 4

   fmt.Println(animalA)
   fmt.Println(animalB)
}