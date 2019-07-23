package main

import "fmt"

// way 1 declare a struct
type Persion struct {
	name string
	age int
}




type persion struct {
	string
	int
}


// 结构体嵌套

type Student struct {
	  sex int
	  persion
}

func ModidyName(p *Persion)  {
	(*p).name="changed"
	fmt.Println(p.name)
}



type Point struct {
	x ,y int
}


type Rect struct {
	leftUp,rightDown Point
}


type Rect2 struct {
	leftUp,rightDown *Point
}


func main(){

	    per := struct {
			name string
			age int
		}{
			"eagle",24,
		}
	    fmt.Println(per.name)

	    per2 := persion{"eagle",24}

	    fmt.Println(per2.string)


       stu := Student{10,per2}

       fmt.Println(stu.persion)

	    per3 := Persion{"no",18}


	    per4 := &per3 // 进行地址拷贝

	    fmt.Printf("per4 %p, per3 %p \n",&per4,&per3)

	    fmt.Println(per3.name,per4.name)
	    ModidyName(&per3)

	    fmt.Println(per3.name,per4.name)


	   //2： 结构体 中所有字段在内存中是连续分布的

	   r1 := Rect{Point{1,2},Point{3,4}}

	   // 打印地址
	   fmt.Printf("r1.leftUp,x addr=%p \n r1 leftUp.y addr is %p \n",&r1.leftUp.x,&r1.leftUp.y)
	   fmt.Printf("r1.rightDown,x addr=%p \n r1 rightDown.y addr is %p",&r1.rightDown.x,&r1.rightDown.y)
	   /*
	    r1.leftUp,x addr=0xc0000520e0
	    r1 leftUp.y addr is 0xc0000520e8
	    r1.rightDown.x addr=0xc0000520f0
	    r1 rightDown.y addr is 0xc0000520f8
	   */

	    r2 := Rect2{&Point{1,2},&Point{3,4}}

	    fmt.Println("\n")


	    //  他们的地址不是连续的但是 指向的地址是连续的
	    fmt.Printf("r2  leftup is %p r2 rightdown is %p \n",r2.leftUp,r2.rightDown)
	    fmt.Printf("r2  leftup point to  %p r2 rightdown point to  %p \n",&r2.leftUp,&r2.rightDown)
	    /*
	    r2  leftup is 0xc0000540b0 r2 rightdown is 0xc0000540c0
		r2  leftup point to  0xc000040230 r2 rightdown point to  0xc000040238
		*/

		fmt.Printf("r2.leftUp,x addr=%p \n r2 leftUp.y addr is %p \n",&r2.leftUp.x,&r2.leftUp.y)
		fmt.Printf("r2.rightDown,x addr=%p \n r2 rightDown.y addr is %p",&r2.rightDown.x,&r2.rightDown.y)
}
