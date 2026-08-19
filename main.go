package main

import (
	"fmt"
	"os"
)

// log_aggregator - Centralized log collection
func log_aggregator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Log-Aggregator")
	fmt.Println("  Centralized log collection")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	log_aggregator(path)
}
