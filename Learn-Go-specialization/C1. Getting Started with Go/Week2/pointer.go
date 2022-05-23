package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int
	ip = &x // ip = address of x
	y = *ip // y = data at the address:ip

	//new() function creates a variable and returns a pointer to the variable
	prt := new(int)
	*prt = 4

	//https://www.w3schools.com/go/go_formatting_verbs.php output verbs
	fmt.Print("Value of prt=", *prt, "\nValue of y=", y)
}
