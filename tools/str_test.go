package tools

import (
	"fmt"
)

func ExampleGenRandomName() {
	fmt.Println(len(GenRandomName(10)))
	fmt.Println(len(GenRandomName(20)))
	fmt.Println(len(GenRandomName(0)))
	// Output:
	// 10
	// 20
	// 0
}
