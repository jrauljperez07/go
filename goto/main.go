package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("Do something when the loop is done")

done:
	fmt.Println("Do complex stuff here")
	fmt.Println(a)
}
