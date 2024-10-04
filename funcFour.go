package main

import "fmt"

// func singleIn(in int) int {
// return in
// }

// func multIn(a, b int, c int) int {
// return a+b+c
// }

// именованный результат
// func namedReturn() (out int) {
// out = 2
// return
// // корректная работа
// // return 3
// }

// несоклько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happened")
	}
	return in, nil
}

// несколько именованных результатов
func multipleNameReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmtErrorf("some error happend")
		return

	}
	rez = 2
	return
}

// не фиксиированное количество параметров
// ...int дает слайс интов на входе
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	fmt.Println(multipleNamedReturn(true))
	return

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return
}
