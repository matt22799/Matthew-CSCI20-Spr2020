// Programmer name: Matthew Davenport
// Date completed:  4/9/20
// Description: This program will take a total amount of seconds then convert them to hours, minutes and seconds

package main

import "fmt"

// HoursMinutesSeconds computes the number of hours, minutes,
// and seconds given a number of seconds.
// EXAMPLE: 3661 seconds is 1 hours, 1 minutes, 1 seconds
func HoursMinutesSeconds(inSeconds int) (int, int, int){
  hours := inSeconds%60
  minutes := hours%60
  seconds := minutes%60
  
	return hours, minutes, seconds
}

func main() {
//call your function
hours, minutes, seconds := HoursMinutesSeconds(3661)
fmt.Println(hours, minutes, seconds)
}