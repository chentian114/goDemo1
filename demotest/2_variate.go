package main

import "fmt"

func main(){
	fmt.Println("variate test ")

	//定义一个类型为int的变量
	var num int
	fmt.Println("num",num)
	//定义多个类型为int的变量
	var num1 , num2 int
	fmt.Println("num1",num1,"num2",num2)
	//定义并初始化
	var str string = "hello world"
	fmt.Println("str",str)
	//定义多个变量并初始化
	var str1, str2,str3 string = "abc","def","hlg"
	fmt.Println("str1",str1,"str2",str2,"str3",str3)
	//上面简化写法，根据值的类型初始化
	var n1 , n2 = 1,2
	fmt.Println("n1",n1,"n2",n2)
	//更简洁的写法
	//但这种方式只能在函数内部使用
	//:=这种操作是先声明再初始化两步操作，你对一个变量两次:=操作会报错
	st := "hello world"
	st1,st2,st3 := "abc","def","hlg"
	fmt.Println("st",st,"st1",st1,"st2",st2,"st3",st3)
	//特殊的变量名_，任何赋予它的值都会被丢弃
	a,_ := 1,2
	fmt.Println("a",a)


	//常量可定义为数值、布尔值或字符串
	const NUM = 1.34
	const FLAG = true
	const HELLO = "hello world"
	fmt.Println("NUM",NUM,"FLAG",FLAG,"HELLO",HELLO)
}
