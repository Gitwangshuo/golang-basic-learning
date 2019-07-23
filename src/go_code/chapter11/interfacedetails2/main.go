package main

import "fmt"

type BInterface interface {
	test01()
	test02()
}

type CInterface interface {
	test01()
	test03()
}


type Stu struct {


}


func (stu Stu)test01(){

}

func (stu Stu)test02(){

}

func (stu Stu)test03(){

}

func main() {


	stu := Stu{}
	var b BInterface =stu
	var c CInterface =stu

	//
	fmt.Print(b,c)

}
