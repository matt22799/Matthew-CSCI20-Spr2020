// Programmer name: Matthew Davenport
// Date completed:  5/11/20
// Description: This program will start by asking the user what kind of coffee they want, followed by asking if they want addons and what they want. Then it will run the corresponding function and display back the instructions, addons, and total cost

package main

import "fmt"

// create struct for coffee using cost and extra
type Coffee struct {
  cost_ float64
  extra_ float64
}

// Create functions for each type of coffee using parameters for extras and addons
func frappe(extras_ float64)(float64){
  frappe_ := Coffee{
    cost_: 4.25,
    extra_: extras_ * 1,
  }
  total_ := (frappe_.cost_ + frappe_.extra_)
  return total_
}

func latte(extras_ float64)(float64){
  latte_ := Coffee{
    cost_: 4.25,
    extra_: extras_ * 1,
  }
  total_ := (latte_.cost_ + latte_.extra_)
  return total_
}

func flatwhite(extras_ float64)(float64){
  flatwhite_ := Coffee{
    cost_: 4.25,
    extra_: extras_ * 1,
  }
  total_ := (flatwhite_.cost_ + flatwhite_.extra_)
  return total_
}

func cappuccino(extras_ float64)(float64){
  cappuccino_ := Coffee{
    cost_: 4.25,
    extra_: extras_ * 1,
  }
  total_ := (cappuccino_.cost_ + cappuccino_.extra_)
  return total_
}

func macchiato(extras_ float64)(float64){
  macchiato_ := Coffee{
    cost_: 4.25,
    extra_: extras_ * 1,
  }
  total_ := (macchiato_.cost_ + macchiato_.extra_)
  return total_
}

func main() {
// Collect user input
  var input string
  var extra float64
  var addons string

  // ask for user order
  fmt.Println("What would you like to order?(Frappe, Latte, FlatWhite, Cappuccino, Macchiato)")
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
    fmt.Println("Prepare espresso then chill, combine espresso with milk, sugar and ice in blender, blend until smooth, then pour and serve")
    fmt.Println("Add extra", addons) 
    order := frappe(extra)
    fmt.Println("Total Price:", order)  

  } else if (input == "Latte"){
    fmt.Println("Brew espresso, heat milk, combine milk and espresso, top with foam and serve")
    fmt.Println("Add extra", addons) 
    order := latte(extra)
    fmt.Println("Total Price:", order)   

  } else if (input == "FlatWhite"){
    fmt.Println("Brew espresso, heat milk, combine milk and espresso, top with foam and serve")
    fmt.Println("Add extra", addons) 
    order := flatwhite(extra)
    fmt.Println("Total Price:", order)  

  } else if (input == "Cappuccino"){
    fmt.Println("Brew espresso, heat milk, combine milk and espresso, top with foam and serve")
    fmt.Println("Add extra", addons)  
    order := cappuccino(extra)
     fmt.Println("Total Price:", order)  

  } else {
    fmt.Println("Brew espresso, heat milk, combine milk and espresso, top with foam and serve")
    fmt.Println("Add extra", addons) 
    order := macchiato(extra)
    fmt.Println("Total Price:", order)  
  }
}