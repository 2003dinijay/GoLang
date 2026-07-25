package main

import (
	"fmt"
)

func main() {
	views1 := 100
	views2 := 200
	totalViews := views1 + views2

	likes := 50
	dislikes := 10
	totalLikes := likes - dislikes

	fmt.Println("Total Views:", totalViews, "Total Likes:", totalLikes)

}
