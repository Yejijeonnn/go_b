package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2, 10))
	fmt.Println(mymath.MyAbs(-99))
	greeting.Hi()
	greeting.Hello()
}
