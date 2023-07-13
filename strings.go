package main

import (
	"fmt"
	"strconv"
)

func stringFunc () {
	// rawString()
	// listing()
	// runeByte()
	// pullingString()
	// conVar()
	strCon()
}

func rawString() {
	fmt.Println("peace be upon you \nupon you be peace")

	fmt.Println(`strings can span multiple lines with \n escape sequence`)

	fmt.Println(`
		peace be upon you 
		upon you be peace
	`)
}

func listing() {
	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}

func runeByte() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
}

func pullingString() {
	str := "hello boy"
	fmt.Println(len(str))
}

func conVar() {
	age := 41
	marsAge := float64(age)

	marsDays := 687.0 
	earthDays := 365.3435

	marsAge = marsAge * earthDays / marsDays

	fmt.Println("I am", marsAge, "years old on mars")

	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h)

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Print(string(pi), string(omega), string(bang), string(alpha))
	
}

func strCon() {
	countdown := 10

	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds"

	fmt.Println(str)

	str2 := fmt.Sprintf("Launch in T minus %v seconds", countdown)

	fmt.Println(str2)

	countdown, err := strconv.Atoi("10")

	if err != nil {
		fmt.Println("oh no, something went wrong")
	}

	fmt.Println(countdown)
}
func typeStatic() {
	
}