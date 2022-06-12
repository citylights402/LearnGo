package main

import "fmt"

func main() {
	var myMap map[string]string
	myMap = make(map[string]string, 10) //分配10个空间
	myMap["name"] = "zhou"
	myMap["age"] = "18"
	myMap["age"] = "22"
	fmt.Println(myMap) //map[age:22 name:zhou] map中内容是无序的，key 不能重复

	citys := make(map[string]string)
	citys["No1"] = "北京"
	citys["No2"] = "上海"

	students := map[string]string{
		"1": "John",
		"2": "bob",
		"3": "hery", //注意最后没有元素也要加逗号
	}
	fmt.Println(students) //map[1:John 2:bob 3:hery]
}
