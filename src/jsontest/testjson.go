package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time string
}

func main() {
	var m Message
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":"129"}`) //字符串
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("json Unmarshal error")
	}
	fmt.Println(m.Name)
	// fmt.Printf("%T\n", b)
	// fmt.Println(b)
}
