package main

import (
	"bufio"
	"fmt"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // 입력하는만큼 반복
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	// var games map[int]string
	// games = make(map[int]string)
	games := map[int]string{
		456: "성기훈",
		218: "박해수",
		067: "강새벽",
		001: "오일남",
		199: "알리",
		101: "아이오아이",
	}

	// fmt.Println(games[067])
	for _, v := range games {
		fmt.Println(v)
	}

	// update
	games[101] = "장덕수"

	// delete
	delete(games, 199)

	for k, v := range games {
		fmt.Println(k, v)
	}
}
