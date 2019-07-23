package work

import "fmt"

type Rect struct {
	length,width int
}


type Math struct {
	x,y float64
}


func (m Math)Caculator(oper byte)float64{

	res := 0.0

	switch oper {

	case '+':
		res =  m.x + m.y
	case '-':
		res =  m.x - m.y
	case '*':
		res =  m.x * m.y
	case '/':
		res =  m.x / m.y
	default:
		fmt.Printf("oper %v was not allowed",oper)

	}


	return res
}

func (r *Rect)MethodUtils(){

	for i :=0 ;i< r.width ; i++ {

		for j:=0 ; j< r.length; j++  {
			fmt.Print("*")

		}
		fmt.Println()
	}

}

func (r *Rect)GetPage () int{



	return r.length*r.width
}


type TwoArr struct {
	change [3][3]int
}

func (t *TwoArr)ChangeArr(){



	for i := 0;i < len(t.change) ; i++ {

		for j :=0 ; j < len(t.change[i]); j++ {

			if i < j {
				tmp := t.change [i][j]
				t.change[i][j] = t.change [j][i]
				t.change[j][i] = tmp
			}

		}


	}
	
   fmt.Print(t.change)

}

type Vistor struct {
	Name string
	Age   int
}

func Factor(name string,age int) *Vistor{


	return &Vistor{
		Name:name,
		Age:age,
	}
}

func (v *Vistor)String() string{


	if v.Age > 18 {
		return fmt.Sprintf("%v 的年龄是 %v,门票的价格是 %v ",v.Name,v.Age,20)
	}else{
		return fmt.Sprintf("%v 的年龄是 %v,门票的价格是 %v ",v.Name,v.Age,20)
	}

}



func main(){

	//
	//r := Rect{10,8}
	//(&r).MethodUtils()
	//
	//page := r.GetPage();
	//
	//fmt.Print(page)
	//
	//m := Math{10,9}
	//
	//res := m.Caculator('+')
	//
	//fmt.Println(res)
	//
	//
	//arr := TwoArr{([3][3]int{{1,2,3},{4,5,6},{7,8,9}})}
	//
	//(&arr).ChangeArr()
	//
	//fmt.Println(arr.change)
	//
	//arr.ChangeArr()
	for {
		var v Vistor
		fmt.Println("please type your name :")
		fmt.Scan(&v.Name)
		fmt.Println("please type your age :")
		fmt.Scan(&v.Age)

		fmt.Println(&v)
	}

}