package main

import (
	"errors"
	"fmt"
)

func test(name string) (err error) {
	if name == "config" {
		return nil
	} else {
		return errors.New("读取文件错误。。")
	}
}

func main() {
	err := test("conf")
	if err != nil { //error默认值是nil
		panic(err) //程序不会往下执行
	}
	fmt.Println("继续执行。。")
}
