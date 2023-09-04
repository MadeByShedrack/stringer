package main

import (
	"fmt"
	"time"
)

func main() {

	for range time.Tick(time.Second * 2) {
		fmt.Printf("The time is -> %v\n", time.Now())
	}
}
