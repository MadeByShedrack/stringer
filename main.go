package main

import (
	"fmt"
)

func main() {

	// for range time.Tick(time.Second * 2) {
	// 	fmt.Printf("The time is -> %v\n", time.Now())
	// }

	userName := "john"

	for _, u := range userName {
		fmt.Printf("Name letters -> %v\n", u)
	}
}
