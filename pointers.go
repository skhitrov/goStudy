package main

import "fmt"

func main()

var a := 2
fmt.Println(a)
var b := &a
fmt.Println(b)
*b = 3 // a = 3
fmt.Println(a)
var c := &a //new point to var a
