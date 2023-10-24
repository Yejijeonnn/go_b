package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	// // 변수명은 영문자로 시작해야된다.
	// var e string
	// var d bool
	// var c rune // 변수명 c, 타입은 rune (rune은 한글자 -> 따옴표)
	// var b float64
	// var a int

	// // naming convention
	// var studentId string
	// var i int32

	// fmt.Println(studentId)
	// fmt.Println(i)

	// // zero value : 값을 할당하지 않았을 때 자동으로 생성되는 값
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(a))
}
