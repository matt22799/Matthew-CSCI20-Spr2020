package main

import "fmt"

func main() {
// Average Quiz Score
  total := 0
  numGrades := 0
  var grade int

for grade > -1 {
    fmt.Println("Enter a grade from 0-100")
    fmt.Scanln(&grade)

    if (grade > -1) {
      total = total + grade
      numGrades++
    } else if (grade == -1){
      average := total/numGrades
      fmt.Println("Total Average Grade:", average)

    } else {
      fmt.Println("Invalid Response")
    }
  }

// Obnoxious Child
  // collect value
  var age int
  fmt.Println("How old is the child?")
  fmt.Scanln(&age)

  // start loop based on value
  for i := 0; i < age; i++{
  fmt.Println("Are we there yet?")
  }  

// Song Lyrics 
  // Ask for number of repeats
  var repeat int
  fmt.Println("How many times do you want to hear the song?")
  fmt.Scanln(&repeat)

  // start loop based on value
  for i := 0; i < repeat; i++{
  fmt.Println("We all live in a yellow submarine, a yellow submarine, a yellow submarine")
  }  
}