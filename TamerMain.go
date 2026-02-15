package main

import (
	"fmt"
)

func main() {
	blackList := "admin"
	maxAttempts := 3

	var userName string
	var attempts int

	fmt.Println("=== System (Go Edition) ===")

	for attempts < maxAttempts {
		fmt.Printf("\nAttempt %d of %d. Enter your name : ", attempts+1, maxAttempts)

		fmt.Scan(&userName)

		if userName == blackList {
			fmt.Println("You are blocked")
		} else if userName == "tamer" {
			fmt.Println("Welcome <3")
			return
		} else {
			fmt.Println("Who are you, bro?", userName)
		}
		attempts++
	}
	fmt.Println("\nGoodbye!")
}
