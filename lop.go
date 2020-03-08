package main

import (
	"fmt"
	"strconv"
)

func testFor() int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

func covertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		testFor(),
		covertToBin(10),
	)
}
