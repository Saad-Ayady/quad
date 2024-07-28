// coded by : SAAD AYADY
// student in : 01 Talent 
// commit : I Love Coding <3
package main

import "github.com/01-edu/z01"

// func to print error message
func printError(e string) {
	for _, str := range e {
		z01.PrintRune(str)
	}
	z01.PrintRune('\n')
}
// quad A func 
func QuadA(x, y int) {
	// check if x and y are positive
	if x <= 0 || y <= 0 {
		printError("Error : Invalid Input pls enter positive number :[")
		return
	}
	for Y := 0; Y < y; Y++ {
		for X := 0; X < x; X++ {
			if (X == 0 && Y == 0) || (X == x-1 && Y == 0) || (X == 0 && Y == y-1) || (X == x-1 && Y == y-1) {
				z01.PrintRune('o')
			} else if ((X > 0 && X < x-1 ) && Y == 0) || ((X > 0 && X < x-1 ) && Y == y-1) {
				z01.PrintRune('-')
			} else if ((Y > 0 && Y < y-1 ) && X == 0) || ((Y > 0 && Y < y-1 ) && X == x-1) {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
