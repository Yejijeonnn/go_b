package main

import "fmt"

// shadowing
func main() {
	// 자료타입을 변수명으로 사용
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	var univ string = "inha"

	var f1 string = "functions"
	var f2 = append([]string{}, "함수")

	fmt.Println(test1 + test2)
	fmt.Println(univ)
	fmt.Println(f1)
	fmt.Println(f2)

	// 자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// vat test float64 = 7.9
	// fmt.Println(float64)

	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println()

	// 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]sgring{}, "함수")

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// // 입력(0과 1처리)된 수의 소수 판정 프로그램 v0.9
// func main() {
// 	var number int

// 	fmt.Print("정수 입력: ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(1)		// 0을 쓰면 성공적인 종료를 의미함 -> exit status 1 출력 없이 종료됨
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number일 때 loop을 돌지 않음
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수 입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.7
// func main() {
// 	var number int
// 	fmt.Print("정수 입력: ")
// 	_, err := fmt.Scanln(&number)

// 	// fmt.Println(n)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number일 때 loop을 돌지 않음
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수 입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // 입력된 수의 소수 판정 프로그램 v0.7
// func main() {
// 	fmt.Print("정수 입력(2 이상의 수): ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)      // 불필요한 앞뒤 문자 제거
// 	number, err := strconv.Atoi(in) // in(입력받은 숫자) -> 정수형, n진수, 비트 사이즈
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ { // 1과 number일 때 loop을 돌지 않음
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수 입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }
