// Programmer name: Matthew Davenport
// Date completed:  3/30/20
// Description: This program will ask the user a question then tell them if they were correct or not in a randomized fashion

package main

import (
  "fmt"
  "math/rand"
  "time"
)

// Generate response if correct
func correct(){
  // Generate random number
  rand.Seed(time.Now().UnixNano())
  correctNumb := (rand.Intn(4 - 1 + 1) + 1)
  // Compare random number
  if (correctNumb == 1) {
    fmt.Println("Good job!")
  } else if (correctNumb == 2) {
    fmt.Println("Nice work!")
  } else if (correctNumb == 3) {
    fmt.Println("Keep it up!")
  } else {
    fmt.Println("You're doing amazing!")
  }
}
// Generate response if incorrect
func wrong(){
  // Generate random number
  rand.Seed(time.Now().UnixNano())
  wrongNumb := (rand.Intn(4 - 1 + 1) + 1)
  // Compare random number
  if (wrongNumb == 1) {
    fmt.Println("Wrong!")
  } else if (wrongNumb == 2) {
    fmt.Println("Try again!")
  } else if (wrongNumb == 3) {
    fmt.Println("Nope!")
  } else {
    fmt.Println("Try harder next time!")
  }
}

func main() {
  // Ask a quiz question
  var answer float64
  fmt.Println("What is the square root of pi?")
  fmt.Scanln(&answer)

  // Compare if answer is correct or wrong
  if (answer == 1.77){
    correct()
  } else {
    wrong()
  }
}