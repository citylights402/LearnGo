package main

import "fmt"

func main() {
	str := "hello world"
	arr2 := []rune(str)
	arr2[0] = '北' //注意这里是字符
	str = string(arr2)
	fmt.Println("str=", str) //str= 北ello world
}
