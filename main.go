// Programmer name: Matthew Davenport
// Date completed:  2/25/20
// Description: 

package main

import (
    "fmt"
    "math/rand"
) //adding the ability to do random numbers

func main() {
    //create two variables - one for the computer and one for the user
    var userGuess int
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 1=rock, 2=scissors, 3=paper
    var computerGuess = rand.Intn(3)

    //prompt the user for an integer value representing the player's choice
    fmt.Println("Enter a number 1-3 (1=rock, 2=scissors, 3=paper)")
    fmt.Scanln(&userGuess)

    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    //You will need to use decisions for this
// computer display back
    if (computerGuess == 1){
      fmt.Println("Computer chose rock")
    } else if (computerGuess == 2){
      fmt.Println("Computer chose scissors")
    } else {
      fmt.Println("Computer chose paper")
    }
// player display back
    if (userGuess == 1){
      fmt.Println("Player chose rock")
    } else if (userGuess == 2){
      fmt.Println("Player chose scissors")
    } else {
      fmt.Println("Player chose paper")
    }
}