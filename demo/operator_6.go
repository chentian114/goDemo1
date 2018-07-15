package main

import "fmt"

func main() {
	fmt.Println("hello 运算符")

	// 算术运算符 +,-,*,/,++,--
	//关系运行付 ==,>=,<=,!=
	fmt.Println("1==2",(1==2))
	//逻辑运行符 !,&&,||
	fmt.Println("!(4>3)",!(4>3))
	// fmt.Println("0 < 8 < 10",(0<8<10)) //error
	fmt.Println("0 < 8 && 8<10 ",(0<8 && 8<10))
	//位运算符 &,|,^,<<
	//赋值运算符 =,+=,
	//& 取地址运行符
	//* 取值运算符



}
