package mymath

func MyPower(base int, exponent int) int {
	res := 1
	for i := 1; i <= exponent; i++ {
		res = res * base
	}
	return res
}

func MyAbs(number int) int {
	if number < 0 { // 음수면
		return number * -1
	}
	return number // 양수면
}
