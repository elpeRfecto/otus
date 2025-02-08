package main

import "fmt"

func main() {
	var N int
	println("Enter N: ")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := 0; j < N * 2; j++ {
			if (i+j)%2 == 0 {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
}
