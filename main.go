package main

import "fmt"

func main() {
// Average Quiz score
  var grade float64
  var arrayPlace int = 0
  var arrayLength float64 = 0
  var grades [10] float64

  // collect info from user
  for grade > -1 {
    fmt.Println("Please enter a grade score")
    fmt.Scanln(&grade)
    grades [arrayPlace] = grade
    arrayPlace++
    arrayLength++
  }
  //take off one place in array
  arrayPlace--
  arrayLength--

  // find average of all grades in array
  var total float64 = (grades[0] + grades[1] + grades[2] + grades [3] + grades [4] + grades [5] + grades [6] + grades [7] + grades [8] + grades [9]) / arrayLength

  // Display total
  fmt.Println(total)

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