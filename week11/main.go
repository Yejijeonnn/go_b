package main

import "fmt"

func main() {

	primes := [3]int{2, 3, 5} // 배열 리터럴로 primes 배열 초기화
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true} // 초기화하지 않은 배열 원소의 제로 값은 해당 원소의 타입의
	fmt.Println(test[3])

	// for prime := range primes {	// 소수 값을 배열에서 꺼내려고 하지만 인덱스가 출력
	for _, prime := range primes {
		fmt.Println(prime)
	}

	// for idx, prime := range primes {
	// 	fmt.Println(idx, prime)
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	// i := 0
	// for i < 3 {
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// i := 0
	// for i < len(primes) {
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// var primes [3]int = [3]int{2, 3, 5}		// 배열 리터럴로 primes 배열 초기화
	// fmt.Println(primes, primes[1])

	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])
}
