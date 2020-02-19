// Programmer name: __________
// Date completed:  __________
// Description: ___________________________

package main

import (
    "fmt"
    "math/rand"
    //"time"
) //adding the ability to do random numbers

func main() {
    //create a variable for count\
    var count int

    //ask the user to enter a max range for the guessing game and store that value in variable max.
    var max int
    fmt.Println("Please enter a a max range")
    fmt.Scanln(&max)
    
    //this next line creates a random number from 1 to that guess for the computer to know.  You can test this by printing out the variable computerGuess
    var computerGuess = rand.Intn(max)

    //ask the user to enter a guess for the computer number
    var userGuess int
    fmt.Println("Please enter a numerical guess")
    fmt.Scanln(&userGuess)

    //create a loop that compares the computerGuess to the userGuess while they are NOT equal go into the loop
    for userGuess != computerGuess{
      //increase the count by 1
      count ++
      //tell the user that the guessed incorrect
      fmt.Println("Wrong!")
      //ask the user to enter a new guess for the computer number
      fmt.Println("Enter a new guess")
      fmt.Scan(&userGuess)
    }   
    
    //print out that the user got the answer correctly and how many guesses it took (the count)
    fmt.Println("Correct! That took you", count, "tries!")

}