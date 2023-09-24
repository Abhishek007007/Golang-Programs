package main

import "fmt"

func main() {

	odd_or_even := []int{}

	for i := 0; i <= 10; i++ {
		odd_or_even = append(odd_or_even, i)
	}

	for _, num := range odd_or_even {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
