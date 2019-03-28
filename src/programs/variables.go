/**
Variables Program with GoLang
Onur GOKER
info@onurgoker.com
*/

package programs

import (
	"fmt"
	"strconv"
)

func Degiskenler() {

	//define primitive types
	var newStr string = "Welcome to Variables!"
	var count int = 1

	fmt.Println(newStr + "This is your " + strconv.Itoa(count) + "st code")

	//define variable without primitive types
	var newStr2 = "For your "
	var count2 = 2

	fmt.Println(newStr2 + strconv.Itoa(count2) + "nd code, variables are defined without primitive types!")

	//define variables with short hand notation
	newStr3 := "It is also possible to define variables with short hand notation."
	count3 := 3

	fmt.Println(newStr3 + "Here you typed your " + strconv.Itoa(count3) + "rd code block!")
}
