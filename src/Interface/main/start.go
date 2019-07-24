package main

import "fmt"

type Usb interface {


	Start()

}

type Phone struct {

}

func(p Phone)Start(){
	fmt.Println("手机开始工作")
}

type Camera struct{

}

func (c Camera)Start(){

	fmt.Println("cameral start working")
}

type Computer struct{



}

func (c Computer) Woring(usb Usb){

	 usb.Start()
}


func main()  {


	computer := Computer{}
    phone    := Phone{}
	camera   := Camera{}

	computer.Woring(phone)
	computer.Woring(camera)

}