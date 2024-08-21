package main

import "fmt"

func main() {

	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	newBuf := buf[:]
	newBuf[0] = 9

	fmt.Println(buf, newBuf)

	newBuf = append(newBuf, 6)
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)
}
