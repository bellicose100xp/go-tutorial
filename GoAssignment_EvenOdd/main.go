package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}

	for _, num := range slice {
		if num%2 == 0 {
			fmt.Printf("The number %v is even\n", num)
		} else {
			fmt.Printf("The number %v is odd\n", num)
		}
	}
}
