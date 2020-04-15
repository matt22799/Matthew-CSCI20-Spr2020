// Programmer name: Matthew Davenport
// Date completed:  4/9/20
// Description: This program will start by finding the distance between 2 points, then will calculate the miles per gallon of a car, then will find average of points obtained and available points, then finally calculate the tip of a bill

package main

import "fmt"

// create function to find distance between 2 points
func distance(x1 float64,y1 float64,x2 float64,y2 float64)(float64){
  totalDistance := (x2 - x1) + (y2 - y1)
  return totalDistance
}

// create func to divide miles by gallons taking 2 parameters
func milesPerGallon(milesDriven int, gallonsUsed int)(int){
  mpg := milesDriven/gallonsUsed
  return mpg
}

// create func to divide points obtained by available points
func findAverage(totalPoints float64, availablePoints float64)(float64){
  averagePoints := totalPoints/availablePoints
  return averagePoints
}

// create func to multiply total by gratuity
func tipCalculator(totalSpent float64, gratuity float64)(float64){
  tip := gratuity * totalSpent
  return tip
}

func main() {
// Distance Calculator
  totalDistance := distance(20,50,80,87)
  fmt.Println("Total distance is", totalDistance)

// Miles per gallons
  mpg := milesPerGallon(87, 12)
  fmt.Println("Your miles per gallon is", mpg)

// Compute Average 
  average := findAverage(20,80)
  fmt.Println("Your average points obtained was", average)

// Tip Calculator
  tipTotal := tipCalculator(50, .20)
  fmt.Println("Your tip amount is", tipTotal)
}