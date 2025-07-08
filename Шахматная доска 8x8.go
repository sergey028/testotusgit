package main

import "fmt"

func main() {
	size := 8

	for i := 0; i < size; i += 1 {
		for j := 0; j < size; j += 1 {

			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
git