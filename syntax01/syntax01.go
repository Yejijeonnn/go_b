package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가' // 변수명 c, 타입은 rune (rune은 한글자 -> 따옴표)
	// var a int16 = 7
	// var a = 7
	a := 7
	// var ab int = 7
	// fmt.Print(ab)
	var b float64 = 5.3
	a = int(b) // Type Conversion, Casting : 형변환 -> 소수형에서 int형변화으로 소숫점 잘림
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        // 유니코드값(UTF-8) 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // 룬의 실제 타입을 출력 -> int32 : rune은 integer32타입의 별명이라고 생각하면 됨 (4byte 사용)

	fmt.Println(math.Round(2.71)) // ceil : 올림 , floor : 내림 , round : 반올림
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java")) // .Title() : 단어의 첫글자를 대문자로
}
