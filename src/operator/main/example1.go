package main

import (
	"fmt"
)

func main() {
	// test := swap

	// fmt.Printf("%T\n%T", test, swap) //func(int, int) (int, int)
	x, _ := Cal(2, 3)
	fmt.Println(x) //5

}

func swap(a, b int) (int, int) {
	return b, a
}

//自定义数据类型
type myfunc func(int, int) (int, int)

func test(myswap myfunc, a int, b int) (int, int) { //函数数据类型的参数
	return myswap(a, b)
}

//函数返回值命名
func Cal(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
