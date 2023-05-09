package main

import (
	"fmt"

	"github.com/jrauljperez07/structures/loops"
	"github.com/jrauljperez07/structures/structs"
)

func main() {
	loops.CountToTen()

	person := structs.Person{
		Name: "John",
		Age:  30,
		Pet:  "Dog",
	}

	fmt.Println(person)
}
