// Programmer name: Matthew Davenport
// Date completed:  4/1/20
// Description: This program will ask the user for a grade percentage then convert that to a letter grade. Next, it will ask for the users name and year they were born to display back their age.

package main

import "fmt"

// Function to find age from year
func ageFinder(name string, year int){
  age := 2020 - year
  fmt.Println(name, "you are", age, "years old!")
}

// Function to multiply dozens by 12
func eggCounter(dozens int){
  totalEggs := dozens * 12
  fmt.Println("There are a total of", totalEggs, "eggs!")
}

// Function to swap the inputed numbers
func swapper(num1 int, num2 int){
  fmt.Println("Swapped, the numbers are", num2, "and", num1)
}

// Function to decide letter grade
func gradeReport(grade int){
  var letterGrade string
  if (grade >= 90){
    letterGrade = "A"
  } else if (grade >= 80 && grade < 90) {
    letterGrade = "B"
  } else if (grade >= 70 && grade < 80) {
    letterGrade = "C"
  } else if (grade >= 60 && grade < 70) {
    letterGrade = "D"
  } else{
    letterGrade = "F"
  }
  fmt.Println("Your grade is", letterGrade)
}

func main() {
  // Grade Report
  var input int
  fmt.Println("Enter a grade percentage")
  fmt.Scanln(&input)
  gradeReport(input)

  // Age
  var inputName string
  var inputYear int
  fmt.Println("Please enter your name")
  fmt.Scanln(&inputName)
  fmt.Println("Please enter the year you were born")
  fmt.Scanln(&inputYear)
  ageFinder(inputName, inputYear)

  // How many eggs?
  var inputDozens int
  fmt.Println("Enter how many dozens of eggs")
  fmt.Scanln(&inputDozens)
  eggCounter(inputDozens)

  // Swap
  var input1, input2 int
  fmt.Println("Enter 2 numbers")
  fmt.Scanln(&input1, &input2)
  fmt.Println("you entered", input1, "and", input2)
  swapper(input1, input2)
}