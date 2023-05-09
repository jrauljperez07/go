package main

import (
	"fmt"
)

func main() {
	maps()
}

func maps() {
	teams := make(map[string]interface{})

	teams["Juan"] = 25
	teams["Pedro"] = "30"

	m := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	v, ok := m["Hello"]
	fmt.Println(v, ok)

	v, ok = m["Goodbye"]
	fmt.Println(v, ok)
}

func slices() {
	x := []int{1, 2, 3}
	fmt.Println(x)
}

func arrays() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	var x = [...]int{1, 2, 3}
	fmt.Println(x)

	var y [2][3]int
	fmt.Println(y)

	for i := 0; i < len(y); i++ {
		for j := 0; j < len(y[i]); j++ {
			y[i][j] = i + j
		}
	}

	fmt.Println(len(y) == len(x))
}
