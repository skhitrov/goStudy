package main

import "fmt"

// C = (F-32) * 5/9
func main() {
	fmt.Print("Enter a number: ")
	var F int64
	fmt.Scan("%f", &F)

	C := (F - 32) * 5 / 9

	fmt.Println(C)

}
