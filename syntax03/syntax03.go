package main

import (
	"bufio"
	"fmt"
	"log" // Fatal function
	"os"
	"strconv" // TrimSpace
	"strings" //
)

func main() {
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')
	if err != nil { // <nil> : 에러가 없다
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)   // 불필요한 앞뒤 문자 제거
	dan, err := strconv.Atoi(in) // ParseInt -> in(입력받은 숫자), n진수, 비트 사이즈
	if err != nil {
		log.Fatal(err)
	}

	// 다른 언어의 while문 구현
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i)) // -> C언어처럼
		i++
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log" // Fatal function
// 	"os"
// 	"strconv" // TrimSpace
// 	"strings" //
// )

// func main() {
// 	fmt.Print("숫자 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')
// 	if err != nil { // <nil> : 에러가 없다
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in) // 불필요한 앞뒤 문자 제거
// 	// dan, err := strconv.ParseInt(in, 10, 32)  // 강제로 32로 바꾼 것이기 때문에 사실 64로 돌아가는 것이다.
// 	dan, err := strconv.Atoi(in) // ParseInt -> in(입력받은 숫자), n진수, 비트 사이즈
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for i := 1; i < 10; i++ {
// 		// fmt.Println(dan, " * ", i, " = ", (int(dan) * i))  // -> ParseInt 사용 시
// 		fmt.Println("%d * %d = %d\n", dan, i, (dan * i)) // -> C언어처럼
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("숫자 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, _ := rd.ReadString('\n')
// 	fmt.Println(in)
// }
