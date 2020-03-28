package dog

import (
	"fmt"
)

func ExampleYears() {
	fmt.Println(Years(1))
	fmt.Println(Years(7))
	fmt.Println(Years(10))
	// Output:
	// 7
	// 49
	// 70
}