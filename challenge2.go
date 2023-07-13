package main

import (
	"strings"
	"fmt"
)

func challenge2() {
	// challengeTwo1()
	// challengeTwo2()
}

func challengeTwo1() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		// A = 0, B = 1, ....Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		//cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

		//increment keyIndex
		keyIndex++

		keyIndex %= len(keyword)
	}

	fmt.Println(message)

	
}

func challengeTwo2() {
	message := "your message goes here"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
	c := message[i]

	if c >=  'A' && c <= 'Z' {
		//A= 0, B=1, ...Z=25

		c -= 'A'
		k := keyword[keyIndex] - 'A'

		//cipher letter + key letter
		c = (c+k)%26 + 'A'

		//increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}

	cipherText += string(c)
 }
 
 fmt.Printf(cipherText)

}