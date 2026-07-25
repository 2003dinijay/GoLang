package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Dinithi"
	lastName := "Jayasinghe"
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)
	fmt.Println(strings.ToUpper(fullName))
}
