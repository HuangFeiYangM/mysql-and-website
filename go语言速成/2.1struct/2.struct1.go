package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	//下面的3种写法都对
	//但是不能同时写,也就是多次在同一层使用":="定义同一个变量

	//1
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	//2
	//goku := Saiyan{}

	// or

	//goku := Saiyan{Name: "Goku"}
	//goku.Power = 9000

	//3

	//goku := Saiyan{"Goku", 9000}
	//print(goku.Name)

	//同时,不同作用域允许同名变量

	fmt.Println(goku.Name) //孩子们Println前面一定要加fmt

	//===========================================================
	Super1(goku)
	fmt.Printf("P1=%d\n",goku.Power)//运行结果是9000,要想改就需要使用指针

	
	Super2(&goku)
	fmt.Printf("P2=%d",goku.Power)
	
	
	fmt.Scanln()
}

func Super1(s Saiyan) {
	s.Power += 10000
}
func Super2(s *Saiyan) {
	s.Power += 10000
}
