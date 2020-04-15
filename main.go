// Programmer name: Matthew Davenport
// Date completed:  4/15/20
// Description: This program will start by asking the user what kind of coffee they want, followed by asking if they want addons and what they want. Then it will run the corresponding function and display back the instructions, addons, and total cost

package main

import "fmt"

// Create functions for each type of coffee using parameters for extras and addons

func frappe(extras_ float64, addons_ string){
  cost_ := 5.25
  extra_ := extras_ * 1
  total_ := (cost_ + extra_)
  instructions := "Prepare espresso then chill, combine espresso with milk, sugar and ice in blender, blend until smooth, then pour and serve"
  addedInstructions := addons_
  
  fmt.Println("Instructions:", instructions)
  fmt.Println("Addons:", addedInstructions)
  fmt.Println("Your total cost is", total_)
}

func latte(extras_ float64, addons_ string){
  cost_ := 3.50
  extra_ := extras_ * 1
  total_ := (cost_ + extra_)
  instructions := "Brew espresso, heat milk, combine milk and espresso, top with foam and serve"
  addedInstructions := addons_
  
  fmt.Println("Instructions:", instructions)
  fmt.Println("Addons:", addedInstructions)
  fmt.Println("Your total cost is", total_)
}

func flatWhite(extras_ float64, addons_ string){
  cost_ := 2.75
  extra_ := extras_ * 1
  total_ := (cost_ + extra_)
  instructions := "Brew espresso, heat milk, combine milk and espresso, top with foam, fold foam and serve"
  addedInstructions := addons_
  
  fmt.Println("Instructions:", instructions)
  fmt.Println("Addons:", addedInstructions)
  fmt.Println("Your total cost is", total_)
}

func cappuccino(extras_ float64, addons_ string){
  cost_ := 2.50
  extra_ := extras_ * 1
  total_ := (cost_ + extra_)
  instructions := "Brew espresso, heat milk, combine milk and espresso, top with foam and serve"
  addedInstructions := addons_
  
  fmt.Println("Instructions:", instructions)
  fmt.Println("Addons:", addedInstructions)
  fmt.Println("Your total cost is", total_)
}

func macchiato(extras_ float64, addons_ string){
  cost_ := 3.99
  extra_ := extras_ * 1
  total_ := (cost_ + extra_)
  instructions := "Brew espresso, heat milk, combine milk and espresso, top with foam and serve"
  addedInstructions := addons_
  
  fmt.Println("Instructions:", instructions)
  fmt.Println("Addons:", addedInstructions)
  fmt.Println("Your total cost is", total_)
}

func main() {
// Collect user input
  var input string
  var extra float64
  var addons string
  // ask for user order
  fmt.Println("What would you like to order?(Frappe, Latte, Flat White, Cappucino, Macchiato)")
  fmt.Scanln(&input)
  // ask for any extras
  fmt.Println("Would you like anything extra? ($1 to add extras; enter 0 if no, 1 if yes)")
  fmt.Scanln(&extra)
  // if yes, ask what kind of extras
  if (extra == 1) {
    fmt.Println("What would you like to add?")
    fmt.Scanln(&addons)
  }

// Display back instructions and cost
  if (input == "Frappe"){
    frappe(extra, addons)
  } else if (input == "Latte"){
    latte(extra, addons)
  } else if (input == "Flat White"){
    flatWhite(extra, addons)
  } else if (input == "Cappuccino"){
    cappuccino(extra, addons)
  } else {
    macchiato(extra, addons)
  }
}