package main

import "fmt"

func main() {
	var buf []int
	buf = append(buf, 9, 10)
	buf = append(buf, 12)

	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)

	fmt.Println(buf, otherBuf)

}
