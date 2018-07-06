package main

import "fmt"

func main() {
	// fmt.Printf("%d \t - %b \t - %x \n", 42, 42, 42)
	// fmt.Printf("%d \t - %b \t - %#x  - %q \n", 42, 42, 42) // Decimal Binary Hexdecimal UTF8
	// fmt.Printf("%d \t - %b \t - %#X \n", 42, 42, 42)

	for i := 2000; i < 10000; i++ {
		fmt.Printf("%d \t - %b \t - %x  - %q \n", i, i, i, i)
	}

}
