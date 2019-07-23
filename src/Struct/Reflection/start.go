package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// way 1 declare a struct
type Persion struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


// struct 的每个字段上可以写一个tag 该tag 可以通过反射机制获取 最常见的场景 就是序列化和反序列化

func main(){


	 var per = Persion{ "wangshuo",20}

	 fmt.Println("type:",reflect.TypeOf(per))

	Json, error := json.Marshal(per)

	if error ==nil {
		fmt.Println(string(Json))
		//{"name":"wangshuo","age":20}
	}else{
		fmt.Println(error)
	}

}