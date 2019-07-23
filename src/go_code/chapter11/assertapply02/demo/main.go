package main

import (
	"fmt"
)

type Student struct {

}



func TypManyItems(items ...interface{}){

	for index,x := range items{
		switch x.(type) {
		case bool:
			fmt.Printf("param %d is a bool value is %v \n",index,x)
		case float64:
			fmt.Printf("param %d is a float64 value is %v \n",index,x)
		case float32:
			fmt.Printf("param %d is a float32 value is %v \n",index,x)
		case int:
			fmt.Printf("param %d is a int value is %v \n",index,x)
		case string:
			fmt.Printf("param %d is a string value is %v \n",index,x)
		case *Student:
			fmt.Printf("param %d is a Student value is %v \n",index,x)
		default:
			fmt.Printf("param %d 's type is unkown value is %v \n,",index,x)

		}

	}


}




func main(){



	var a int = 10

	var b string ="shuo"

	var c *Student

	TypManyItems(a,b,c)




	fmt.Print()
}
