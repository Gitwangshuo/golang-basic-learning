package main

import "fmt"

// declare a usb interface
type Usb interface {

	Start()
	Stop()
	
}


type Phone struct {
	name string
}


func (this Phone)Start(){
	fmt.Println(this.name,"this phone start to use")
}

func (this Phone)Stop(){
	fmt.Println(this.name,"this phone stop to use")
}
func (this Phone)Call(){
	fmt.Println(this.name,"this phone start to call")
}



type Cameral struct {
	name string
}


func (this Cameral)Start(){
	fmt.Println(this.name,"this Cameral start to use")
}

func (this Cameral)Stop(){
	fmt.Println(this.name,"this Cameral stop to use")
}
func (this Cameral)Call(){
	fmt.Println(this.name,"this Cameral start to call")
}




type Computer struct {
	
}

func (c Computer)Working(usb Usb){
	
	if t,ok := usb.(Phone);ok{
		// phone is start to calling now 
		t.Call()
	}
	usb.Stop()


	// var,ok := interface.(Struct)
}



func main(){
	
	
	var usbArr [3]Usb
	
	usbArr[0] = Phone{"xiaomi"}
	usbArr[1] = Phone{"huawei"}
	usbArr[2] = Cameral{"huawei"}

	var computer Computer
	for _,value :=range usbArr{
		computer.Working(value)
		fmt.Println()
	}
	
}
