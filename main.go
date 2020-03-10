// **********************************************************
// Programmer name: Matthew Davenport
// Date completed:  3/10/20
// Description: This program will pit 2 players against each other in a game of Nim
// **********************************************************
package main

import "fmt"

func main() {
// set a total number of stones
  var pileOne int = 10
  var pileTwo int = 6
  var total = pileOne + pileTwo
// Set variables for player choices
  var playerOnePile int
  var playerTwoPile int
  var playerOnePull int = 0
  var playerTwoPull int = 0
// Start loop to take turns and set win condition
  for total != 0 {
    if (total != 0) {
      fmt.Println("Pile One:", pileOne)
      fmt.Println("Pile Two:", pileTwo)
      fmt.Println("Total:", total)
      fmt.Println("Player 1, choose a pile to pull from")
      fmt.Scanln(&playerOnePile)
      fmt.Println("How many do you want to pull?")
      fmt.Scanln(&playerOnePull)
      if (playerOnePile == 1) {
        pileOne = pileOne - playerOnePull
      } else if (playerOnePile == 2){
        pileTwo = pileTwo - playerOnePull
      }
      total = pileOne + pileTwo
    } else if (total <= 0){
      fmt.Println("Player 2 wins!")
    }
    if (total != 0){
      fmt.Println("Pile One:", pileOne)
      fmt.Println("Pile Two:", pileTwo)
      fmt.Println("Total:", total)
      fmt.Println("Player 2, choose a pile to pull from")
      fmt.Scanln(&playerTwoPile)
      fmt.Println("How many do you want to pull?")
      fmt.Scanln(&playerTwoPull)
      if (playerTwoPile == 1) {
        pileOne = pileOne - playerTwoPull
      } else if (playerTwoPile == 2){
        pileTwo = pileTwo - playerTwoPull
      }
      total = pileOne + pileTwo 
    } else if (total <= 0) {
      fmt.Println("Player 1 wins!")
    }
  }
}