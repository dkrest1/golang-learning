// weight loss program
package main

import "fmt"

func main () {
	const hoursPerday = 24
	var speed, distance = 100800, 9630000 
	
	fmt.Println(distance/speed/hoursPerday, "days")
	
}