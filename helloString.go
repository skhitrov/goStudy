package main

import "fmt"

func main() {

	helloWorld := "привет мир"

	hello := helloWorld[:12]
	H := helloWorld[0]

	fmt.Println(hello)
	fmt.Println(H)
	
	//конвертация в слайс байт и обратно
	byteString = []byte(helloWorld)
	helloWorld = string(byteString)
}
