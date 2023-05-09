package loops

import (
	"fmt"
)

func CountToTen() {
	for i := 0; i <= 10; i++ {
		// fmt.Println(i)
	}

	// For range loop
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}
