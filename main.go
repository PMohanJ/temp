package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C :
			fmt.Println("Hello World!!!")
		}
	}
}