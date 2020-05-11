//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Name: Matthew Davenport
// Date: 5/11/20
// Description: This program will play a simple game of blackjack, starting by dealing 2 cards to the player and house, then by asking the player if they want more cards until they say no, then finally adding house cards until it is 17 or higher, ending with a comparison of hand values
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

package main

import ( 
      "fmt"
      "math/rand"
      "time"
)

// create variables for player and house hands
var playerHand int = 0
var houseHand int = 0
var cardValue int

// create function to add a card to players hand when “hit”
func playerDraw()(int){
    rand.Seed(time.Now().UnixNano())
    cardValue = (rand.Intn(10) + 1)
    playerHand = playerHand + cardValue
    
    return playerHand
}
// create function to add a card to house's hand until its 17 or more
func houseDraw()(int){
  rand.Seed(time.Now().UnixNano())
  cardValue = (rand.Intn(10) + 1)
  houseHand = houseHand + cardValue

  return houseHand
}

func main() {
  // add 2 cards to player hand and add
  rand.Seed(time.Now().UnixNano())
  cardValue = (rand.Intn(10) + 1)
  playerHand = playerHand + cardValue
  cardValue = (rand.Intn(10) + 1)
  playerHand = playerHand + cardValue
  
  // add 2 cards to house hand and add
  cardValue = (rand.Intn(10) + 1)
  houseHand = houseHand + cardValue
  cardValue = (rand.Intn(10) + 1)
  houseHand = houseHand + cardValue

  // display back current hands
  fmt.Println("Your hand value is",playerHand)
  fmt.Println("The house's hand value is",houseHand) 

  // loop for players turn; if hit, run playerDraw; if stay, display final value
  var choice string
  if (playerHand <= 21){
    for playerHand < 21 {
      fmt.Println("Do you want to hit or stay?")
      fmt.Scanln(&choice)
      if (choice == "hit"){
        playerValue := playerDraw()
        fmt.Println("Your hand value stands at", playerValue)
      } else {
        fmt.Println("Your final hand value is", playerHand)
        break
      }
    }
  }

  // loop for house's turn; ends when value is >= 17
  for houseHand < 17 {
    houseValue := houseDraw()
    fmt.Println("The house's value stands at", houseValue)
  }
  // display house's final value
  fmt.Println("The house's final hand value is", houseHand)

  // compare player hand to house hand
  if (playerHand > 21 || playerHand < houseHand && houseHand <= 21){
    fmt.Println("You Lose!")
  } else if (houseHand > 21 || houseHand < playerHand && playerHand <= 21){
    fmt.Println("You Win!")
  } else {
    fmt.Print("Draw!")
  }
}