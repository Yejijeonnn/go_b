package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) // Args == Arguments
	fmt.Println(os.Args[2])
}

// go run test.go cs inha university
