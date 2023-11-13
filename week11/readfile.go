// average calculates the average of several numbers.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([4]float64, error) {
	var numbers [4]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

// average calculates the average of several numbers.
func main() {
	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

// func main() {
// 	f, err := os.Open("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		fmt.Println(s.Text()) // .Text = 한줄씩 가져오는
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if s.Err() != nil {
// 		log.Fatal(s.Err())
// 	}
// }
