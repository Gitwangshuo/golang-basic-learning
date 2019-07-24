package main



import (
	"fmt"
	"flag"
)


var name string
var age   int

func init(){

	flag.StringVar(&name,"name","zhangsan","姓名")
    flag.IntVar(&age,"age",0,"年龄（default 0）")
}


func main() {
	flag.Parse()
	fmt.Printf("name is %v \n",name)
	fmt.Printf("age  is %v  \n",age)
}
