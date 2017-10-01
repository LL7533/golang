package main

import (
	"fmt"
	//	"strconv"
)

func main() {
	//	var number_test int
	//	number_test = 87
	//	//	b := string(number_test)
	//	//	b := strconv.Itoa(number_test)
	//	//	var i int
	//	for i := 0; i < 10000000000; i++ {
	//		number_test += 1
	//	}
	//	//	fmt.Printf("%s \n", number_test)
	//	fmt.Println(number_test)
	//	var nums = [4]int{2, 7, 11, 15}
	//	var left, right int
	a := 10
	a++
	var p *int = &a
	*p++
	fmt.Println(*p)

}
