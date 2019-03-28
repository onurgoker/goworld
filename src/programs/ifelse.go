/**
If-else Statement Program with GoLang
Onur GOKER
info@onurgoker.com
*/

package programs

import (
	"fmt"
	"regexp"
)

func Ifelse() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	re := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	isNum := re.Match([]byte(input))
	if isNum {
		fmt.Println("This is an integer")
	} else {
		fmt.Println("This is not an integer!")
	}
}
