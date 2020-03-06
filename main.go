// **********************************************************
// Programmer name: Matthew Davenport
// Date completed:  3/4/20
// Description: This program will start by showing a broken
//              array with the users input, then it will
//              display an array and calculate the biggest
//              and smallest number, then finally will Ask
//              for baseball players and stats, then compare
//              the stats.
// **********************************************************
package main

import "fmt"

func main() {
// Broken Array
fmt.Println("Broken Array")
  // Set an array
    var array [10] int
  // Ask for number
    var num int
    for i := 1; i <= 9; i++ {
      fmt.Println("Enter a number")
      fmt.Scanln(&num)
      array[i] = num 
    }
  // Display back
    fmt.Println(array)

// Max and Min
fmt.Println("Max and Min")
	// Set an array of numbers
	a := [10]int {22,68,45,79,52,58,21,1,5,55}
	// Calculate min and max
	var min int = a[0]
	var max int = a[0]

	for _, v := range a {
		if v < min {
		min = v
		}
		if v > max {
		max = v
		}
	}

	// Display back max and min
  fmt.Println(a)
	fmt.Println("Biggest number is", max)
	fmt.Println("Smallest numbers is", min)

// Baseball Stats
fmt.Println("Baseball Stats")
	// Define arrays
	var names [5] string
	var average [5] float64
	var slugging [5] float64

	// Store valus into arrays
	for i := 0; i<= 4; i++ {
		fmt.Println("Enter a players name")
		fmt.Scanln(&names[i])
		fmt.Println("Enter that players batting average")
		fmt.Scanln(&average[i])
		fmt.Println("Enter that players slugging percentage")
		fmt.Scanln(&slugging[i])
	}

	// Display back stats
	for n := 0; n <= 4; n++ {
		fmt.Println("Name:",names[n],"Batting Average:", average[n],"Slugging Percentage:", slugging[n])
	}

	// Compare batting averages 
	var small float64 = average[0] 
	var big float64= average[0]

	for _, value := range average {
		if value < small {
		small = value
		}
		if value > big {
		big = value
		}
	}
	fmt.Println("Best batting average is", big)
	fmt.Println("Worst batting average is", small)

	// Compare slugging percentages
	var worst float64 = slugging[0] 
	var best float64= slugging[0]

	for _, total := range slugging {
		if total < worst {
		worst = total
		}
		if total > best {
		best = total
		}
	}
	fmt.Println("Best slugging percentage is", best)
	fmt.Println("Worst slugging percentage is", worst)

}