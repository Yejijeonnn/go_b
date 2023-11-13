package main

import "fmt"

func main() {
	var s []int = make([]int, 5)

	for _, value := range s {
		fmt.Println(value)
	}
}
