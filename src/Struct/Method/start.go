package main

import (
	"fmt"
	"Struct/Method/work"
)
// way 1 declare a struct
type Person struct {
	name string
	age int
}



func  (p *Person)UpdateName(){

	p.name="update"
}


type Human struct {
	sex  int
	Person
}

func (p *Person)String() string{

	str := fmt.Sprintf("Name =[%v] Age =[%v]",p.name,p.age)
	return str
}

func (p Person)changeName(){

	 p.name="jack"
}

func (p *Person)changeNamePtr(){

	p.name="jack"

}

func main(){

	var per = Person{"tom",18}

	fmt.Println(per)

	(&per).changeName()

    fmt.Println(per.name)  // tom

	(&per).changeNamePtr()

	fmt.Println(per.name) //jack

	 Factor("tome",18)
}