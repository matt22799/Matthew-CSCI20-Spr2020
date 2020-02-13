// Programmer Name: Matthew Davenport
// Date Created: 2/13/20
// Program Description: Ask user for temp in Celsius, convert to Fahrenheit and Kelvin
//                      then display back all values.

package main

import "fmt"

func main() {
  // ask for input
    fmt.Println("Please enter your temperature in Celsius")
    var celsius float64
    fmt.Scanln(&celsius)

  // convert input to F
    var fahrenheit float64 = (celsius * (9.0/5.0)) + 32

  // convert input to K
    var kelvin float64 = celsius + 273.15

  // display back user input and conversions
    fmt.Println("Your temperature in Celsius is", celsius)
    fmt.Println("Your temperature in Fahrenheit is", fahrenheit)
    fmt.Println("Your temperature in Kelvin is", kelvin)

}