package main

import "fmt"

func main(){

	fmt.Println("hello 常量练习:程序运行期间，不可改变")

	const a int = 10
	// a = 20 //error
	fmt.Println("a",a)

	//自动推导
	const b = 10
	fmt.Println("b",b)

	fmt.Printf("b type is %T\n",b)
	//error
	//const c :=10


	//多个不同类型变量定义
	var (
		cn int
		bf float32
	)
	cn , bf = 1,12.3
	fmt.Println("cn = ",cn)
	fmt.Println( "bf = ",bf)

	//多个不同类型常量定义
	const(
		cp1 = 12
		cf = 13.4
	)
	fmt.Println("cp1 =",cp1,"cf ",cf)
}
