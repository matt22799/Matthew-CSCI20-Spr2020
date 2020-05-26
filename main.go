//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Name: Matthew Davenport
// Date: 5/1/20
// Description: This program will play a simple game of blackjack, starting by dealing 2 cards to the player and house, then by asking the player if they want more cards until they say no, then finally adding house cards until it is 17 or higher, ending with a comparison of hand values
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

package main

import ( 
      "fmt"
      "math/rand"
      "time"
)

// function to create a random card value
func randCard()(int){
  rand.Seed(time.Now().UnixNano())
  cardValue_ := (rand.Intn(10) + 1)
  return cardValue_
}

// function to set starting value of hands
func startValue()(int){
  startValue_ := randCard() + randCard()
  return startValue_
}

// function to draw a card
func drawCard(currentValue_ int)(int){
  newValue_ := currentValue_ + randCard()
  return newValue_
}

func main(){
  // give 2 cards to player to start
  playerValue := startValue()
  fmt.Println("Your starting value:", playerValue)
  // give 2 cards to house to start
  houseValue := startValue()
  fmt.Println("the house's starting value:", houseValue)

  // loop for player's turn; if hit, run drawCard; if stay, display final value
  var choice string
  for playerValue <= 21 {
    fmt.Println("Would you like to hit or stay?")
    fmt.Scanln(&choice)
    if (choice == "hit"){
      playerValue = drawCard(playerValue)
      fmt.Println("Your current value is", playerValue)
    } else{
        break
    }
  }
    fmt.Println("Your final value is", playerValue)  

  // loop for house's turn; ends when value is >= 17
  for houseValue < 17 {
    houseValue = drawCard(houseValue)
    fmt.Println("The house's hand value stands at", houseValue)
  }
  fmt.Println("The house's final value stands at", houseValue)

  // compare player's hand to house's hand
  if (playerValue > 21 || playerValue < houseValue && houseValue <= 21){
    fmt.Println("You Lose!")
  } else if (houseValue > 21 || houseValue < playerValue && playerValue <= 21){
    fmt.Println("You Win!")
  } else {
    fmt.Print("Draw!")
  }
}