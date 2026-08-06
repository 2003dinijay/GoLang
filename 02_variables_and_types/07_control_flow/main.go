//if/else, for (Go's only loop — no while/do-while), switch, range
//The loop will stop iterating once the boolean condition evaluates to false.
package main

import "fmt"

func main(){

	// Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

	sum :=0
	for i:=0; i<10; i++{
		sum += i
	}
	fmt.Println(sum)
}

// The init and post statements are optional.
// for is go s while


func main(){
	sum:=1
	for; sum<1000;{
		sum += sum
	}
	fmt.Println(sum)
}
