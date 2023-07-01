package main

import (
	"fmt"
	"math/rand"
	"time"
)

//countdown using the for loop function
func countDown () {
	var count  = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}

	fmt.Println("Liftoff!")
}

// to infinity and beyound 
func infinity() {
	var degrees = 0
	for {
		fmt.Println(degrees)
		degrees++
		if(degrees >= 360) {
			degrees = 0
			if(rand.Intn(2) == 0) {
				fmt.Println("I am going light off right now!")
				break
			}
		}
	}
}

// exercise three 
func exercise3 () {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if(rand.Intn(100) == 0) {
			break
		}
		count--
	}

	if(count == 0) {
		fmt.Println("Liftoff!")
	}else {
		fmt.Println("Launch failed.")
	}
}

//////////////// variable scope ////////////////
func scope() {
	var count  = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}
}

//////// short declaration ///////
// func shortDeclaration() {
// 	var count = 10
// 	word := "ten"
// }


func loop() {
	var count = 0

	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}

// short declearation in if statement
func shortDec() {
	switch num := rand.Intn(10); num {
	case 0: 
	fmt.Println("Space ")
	case 2: 
	fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num )
	}
	
}

var era  = "AD"

func scopeRule() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day )
}

