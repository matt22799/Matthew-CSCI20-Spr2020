// Programmer name: Matthew Davenport
// Date completed:  5/1/20
// Description: This program will start by creating structs for animals, keepers, and zones then initialize each object within a zone object. Then, it will create a struct for inventory and ask the user what they want and then finally update the new inventory

package main

import "fmt"

// Struct for animal
  type Animal struct {
    name string
    species string
    age int
    breakfastHour int 
    dinnerHour int
  }
// struct for keeper
  type Keeper struct {
    name string
    zone string
    animals int 
  }
// struct for zone
  type Zone struct {
    cleaner Keeper
    animal1 Animal
    animal2 Animal
  }

// struct for inventory
  type Inventory struct {
    partNumber int
    unitPrice float64
    quantity int
  }

func main() {
// Zoo Two
  // zone objects that initialize keeper objects and animal objects
    zoneMammal := Zone{
      cleaner: Keeper{"Joe", "Mammal", 6},
      animal1: Animal{"Tony", "Tiger", 4, 8, 6},
      animal2: Animal{"Pepe", "Mule", 6, 7, 5},
    }

    zoneReptile := Zone{
      cleaner: Keeper{"Abby", "Reptile", 8},
      animal1: Animal{"Momo", "Komodo Dragon", 9, 8, 6},
    }

    zoneBird := Zone{
      cleaner: Keeper{"Anthony", "Bird", 7},
      animal1: Animal{"Patchy", "Parrot", 2, 5, 4},
      animal2: Animal{"Snowball", "Penguin", 6, 9, 6},
    }

  // display Zone objects
    fmt.Println("Mammal zone info:", zoneMammal)
    fmt.Println("Reptile zone info:", zoneReptile)
    fmt.Println("Bird zone info:", zoneBird)


// Inventory
  // create initialized objects of Inventory
    snack := Inventory{
      partNumber: 789,
      unitPrice: 3.5,
      quantity: 4,
    }
    meal := Inventory{
      partNumber: 256,
      unitPrice: 8.99,
      quantity: 7,
    }
    drink := Inventory{
      partNumber: 782,
      unitPrice: 1.25,
      quantity: 10,
    }

  // display full inventory
    fmt.Println("Snacks:", snack.quantity)
    fmt.Println("Meals:", meal.quantity)
    fmt.Println("Drinks:", drink.quantity)

  // ask user what they are buying
    var item string
    fmt.Println("What would you like to buy?")
    fmt.Scanln(&item)

  // decide which object needs to be updated
    if (item == "Snack" || item == "snack"){
      snack.quantity--
    } else if (item == "Meal" || item == "meal"){
      meal.quantity--
    } else if (item == "Drink" || item == "drink"){
      drink.quantity--
    } else {
      fmt.Println("Sorry! We are sold out of that!")
    }

  // display updated inventory
    fmt.Println("Snacks:", snack.quantity)
    fmt.Println("Meals:", meal.quantity)
    fmt.Println("Drinks:", drink.quantity)
}