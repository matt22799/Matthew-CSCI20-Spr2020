//*********************************************************
// Programmer name: Matthew Davenport
// Date completed:  2/25/20
// Description: 
//*********************************************************

package main

import (
    "fmt"
    "math/rand"
    "time"
) //adding the ability to do random numbers

func main() {
// Rock, Paper, Scissors

    //create two variables - one for the computer and one for the user
    var userGuess int
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 1=rock, 2=scissors, 3=paper
    rand.Seed(time.Now().UnixNano())
    var computerGuess = (rand.Intn(3)+1)

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
  // Compare who won
	if (userGuess == 1 && computerGuess == 3){
		fmt.Println("You Lose!")
		} else if (userGuess == 3 && computerGuess == 2){
		fmt.Println("You Lose!")
		} else if (userGuess == 2 && computerGuess == 1){
		fmt.Println("You Lose!")
		} else if (userGuess == 1 && computerGuess == 2){
		fmt.Println("You Win!")
		} else if (userGuess == 3 && computerGuess == 1){
		fmt.Println("You Win!")
		} else if (userGuess == 2 && computerGuess == 3){
		fmt.Println("You Win!")
    } else {
      fmt.Println("It's a Tie!")
    }

// Pay Stub
  // Collect info
  var name string
  var wage float64
  var hours float64
  var tax float64 = 0.12

  fmt.Println("Please enter your name")
  fmt.Scanln(&name)
  fmt.Println("Please enter your hourly wage")
  fmt.Scanln(&wage)
  fmt.Println("Please enter your hours")
  fmt.Scanln(&hours)
  
  // If no overtime, solve for base wage, otherwise add overtime
  if (hours < 80){
    var taxedWage float64 = wage * tax
    var total float64 = taxedWage * hours
    fmt.Println(name, ", your pay will be approx", total)
  } else {
    var overtimeHours float64 = hours - 80
    var overtimePay float64 = wage * 1.5
    var overtimeTotal float64 = overtimeHours * overtimePay
    var taxedWage float64 = (wage + overtimePay) * tax
    var total = (taxedWage * hours) + overtimeTotal
    fmt.Println(name, "your pay will be approx", total)
  }

// Guessing game
  // Ask for number between 1 and 10
  var player int
  fmt.Println("Enter a number between 1 and 10")
  fmt.Scanln(&player)

  // Generate random number for computer
  rand.Seed(time.Now().UnixNano())
  var computer = rand.Intn(10)

  // Compare player and computer answers
  if (player != computer){
    if (player > computer){
      fmt.Println("Wrong! You were too high!")
    } else {
      fmt.Println("Wrong! You were too low!")
    }
  }else {
    fmt.Println("Correct!")
  }
}