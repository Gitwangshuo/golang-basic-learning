package main
import "fmt"
type Usb interface {
	Say()
}
type Stu struct {

}


func (this *Stu) Say() {
	fmt.Println("Say()")
}




func main() {

	var usb interface{}

	var stu Stu


	// usb interface 接受一个 stu 类型的变量
	usb = stu


	//   声明 stu2 变量企图接受 usb 接口
	var stu2 Stu


	// 类型断言  判断 usb 是不是Stu 类型的变量如果是的话
	stu2 = usb.(Stu)


	fmt.Print(stu2)


	var x float32
	var y interface{}

	y=x


	z,ok := y.(float32)
	if ok {
		fmt.Print("success",z)
	}else{
		fmt.Print("failer ")
	}


}

	