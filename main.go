// Programmer name: Matthew Davenport
// Date completed:  3/3/20
// Description: This program will hold an array for items and another for prices, then  display back the price to each item.

package main

import "fmt"

func main() {
    //Create a string array of at least 5 values.  It should hold 5 products to buy
    var items [5] string
    items [0] = "milk"
    items [1] = "eggs"
    items [2] = "butter"
    items [3] = "chicken"
    items [4] = "salt"
    //Create a float array of at least 5 values.  It should hold a price for each of the products
    var prices [5] float64
    prices [0] = 3.25
    prices [1] = 2.39
    prices [2] = .75
    prices [3] = 5.99
    prices [4] = 1.0
    //Iterate through the array and output the product and it's price (similar to a menu)
    fmt.Println(items[0], ": $", prices[0])
    fmt.Println(items[1], ": $", prices[1])
    fmt.Println(items[2], ": $", prices[2])
    fmt.Println(items[3], ": $", prices[3])
    fmt.Println(items[4], ": $", prices[4])
}