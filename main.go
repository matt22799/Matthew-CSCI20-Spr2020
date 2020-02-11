package main

import "fmt"

func main() {
  var input float64
  fmt.Println("What is the distance in Rods?")
  fmt.Scanln(&input)

  var feet = input * 17.0
  fmt.Println("Distance in feet is", feet)

  var meters = feet / 3.281
  fmt.Println("Distance in meters is", meters)

  var miles = feet / 5280.0
  fmt.Println("Distance in miles is", miles)

  var totalTime = miles/ 3.1
  fmt.Println("Total time to travel is", totalTime, "hours")

}