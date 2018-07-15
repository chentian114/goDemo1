package main

import "fmt"

func main()  {

	fmt.Println("iota")
	//iota 常量自动生成器，每个一行，自动累加1
	//iota 给常量赋值使用

	const(
		a = iota
		b = iota
		c = iota
	)
	fmt.Println("a",a,"b",b,"c",c)

	//iota遇到const重置0
	const d = iota
	fmt.Println("d",d)

	//iota可以只赋一个
	const(
		a1 = iota
		a2
		a3
	)
	fmt.Println("a1",a1,"a2",a2,"a3",a3)

	//如果是一行值都一样
	const(
		b1 = iota
		b2 , b3 , b4  = iota , iota , iota
		b5 = iota
	)
	fmt.Println("b1",b1,"b2",b2,"b3",b3,"b4",b4,"b5",b5)
}
