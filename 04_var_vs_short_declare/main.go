package main

import (
	"fmt"
)

func main() {
	var city string
	city = "Colombo"

	var country = "Sri Lanka" // type inference, compiler infers the type of the variable based on the value assigned

	// := is a short declaration operator, it can be used to declare and initialize a variable in one line

	subscribers := 5000

	subscribers = subscribers + 1000

	likes, comments := 100, 30

	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Subscribers:", subscribers)
	fmt.Println("Likes:", likes, "Comments:", comments)

}
