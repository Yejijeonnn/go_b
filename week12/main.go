package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // type, length, capacity
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "Z"
	c := append(a, "y")
	// c := append(a, "y", "x")

	c[0] = "q"

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0]) // c만 주소가 다름 왜? -> 용량이 더 필요하니까 -> 그렇다면 c의 주소는 a, as와 같을 수 없음

}
