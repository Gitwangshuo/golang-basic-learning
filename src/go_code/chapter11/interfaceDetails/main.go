package main
import (
	"fmt"
)



type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}


type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =" ,i )
}


type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}
type Monster struct {

}
func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

func main() {
	var stu Stu //结构体变量，实现了 Say() 实现了 AInterface
 	var a AInterface = stu
	a.Say()


	var i integer = 10
	var b AInterface = i
	b.Say() // integer Say i = 10


	//Monster实现了AInterface 和 BInterface  自动判断实现的接口
	var monster Monster = Monster{}
	var a2 AInterface = monster
	var b2 BInterface = monster

	if t, ok := b2.(BInterface); ok {
		fmt.Println("s implements BInterface", t)
	}

	 if value,ok := a2.(AInterface);ok{
	  	fmt.Print(value)
	}
	a2.Say()
	b2.Hello()
}