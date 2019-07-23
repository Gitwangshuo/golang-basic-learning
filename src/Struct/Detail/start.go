package main

import "fmt"

// way 1 declare a struct
type Persion struct {
	name string
	age int
}


type interger int


type Human struct {
	name string
	age   int
}

func main(){

    // 转换需要两个 结构体之间完全一致的字段
	var per Persion
	var human Human
	human = Human(per)
    fmt.Print(human)

}