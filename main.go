package main

import "fmt"

func main() {

  // Population
  var currentPop = 329252070.0
  var secPerYear = 31536000.0
  var birthPerSecond = 1.0 / 9.0
  var deathPerSecond = 1.0 / 11.0
  var immPerSecond = 1.0 / 44.0

  var birthPerYear = birthPerSecond * secPerYear
  var deathPerYear = deathPerSecond * secPerYear
  var immPerYear = immPerSecond * secPerYear

  var nextYear = (currentPop + birthPerYear + immPerYear) - deathPerYear
  var tenYears = nextYear * 10.0

  fmt.Println("The population in the next 10 years will reach", tenYears)


  // Today
  var day = "Tuesday"
  var month = "February"
  var date = 4
  var year = 2020

  fmt.Println("Todays date is",day,",",month,date,",",year)

}