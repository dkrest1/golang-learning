package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	// countDown()
	// infinity()
	exercise3()
}

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
