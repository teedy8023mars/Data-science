package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Enter a string:\n")
	fmt.Scan(&str)
	if strings.Contains(str, "a") && (str[0] == 'i' || str[0] == 'I') && (str[len(str)-1] == 'n' || str[len(str)-1] == 'N') {
		fmt.Print("Found!\n")
	} else {
		fmt.Print("Not Found!\n")
	}
}
