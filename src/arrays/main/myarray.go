package main

import "fmt"

func test(arr *[3]int) {
	(*arr)[0] = 111
}
func main() {
	var arr1 [3]int = [3]int{1, 2, 3}
	// var arr2 = [3]int{1, 2, 3}
	// var arr3 = [...]int{1, 2, 3}                          //用...方式让系统自己判断
	// var arr4 = [3]string{1: "tom", 0: "jack", 2: "marry"} //指定下标赋值

	// fmt.Println(arr1) //[1 2 3]
	// fmt.Println(arr2) //[1 2 3]
	// fmt.Println(arr3) //[1 2 3]
	// fmt.Println(arr4) //[jack tom marry]

	// for index, value := range arr1 {
	// 	fmt.Printf("第%d个是%v\n", index, value)
	// }
	test(&arr1)
	fmt.Println(arr1[0])
}
