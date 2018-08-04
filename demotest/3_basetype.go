package main

import (
	"fmt"
	"errors"
)

func main()  {
	fmt.Println("hello go base type")
	//bool类型
	var flag bool = true
	fmt.Println("flag",flag)
	//数值类型
	var num int = 20
	fmt.Println("num",num)

	//有符号整型
	var numi8 int8 = 2
	var numi16 int16 = -300
	var numi32 int32 = 600
	fmt.Println("numi8",numi8,"numi16",numi16,"numi32",numi32)
	//无符号整型
	var numu8 uint8 = 2
	var numu16 uint16 = 43
	var numu32 uint32 = 54
	fmt.Println("numu8",numu8,"numu16",numu16,"numu32",numu32)

	//rune是int32的别称，byte是uint8的别称
	var numr rune = 32
	var numb byte = 80
	fmt.Println("numr",numr,"numb",numb)
	fmt.Printf("numb %c\n",numb)

	//不同类型的变量之间不允许互相赋值或操作
	//下面是错误的
	//tmp := a + b;

	//浮点数
	var numf float32 = 3.12
	var numd float64 = 5.6
	fmt.Println("numf",numf,"numd",numd)
	//复数
	var numc complex64 = 32 + 54i
	fmt.Println("numc",numc)
	//字符串
	var str string = "hello"
	fmt.Println("str",str)
	//go中字符串是否可改变的，下面是错误的
	//str1[0] = 'w';

	//如果要改变字符串，需要先转成[]byte类型，修改后再转string
	bys := []byte(str)
	bys[0] = 'I'
	str = string(bys)
	fmt.Println("bys.len",len(bys),"str",str)
	//通过+进行字符串连接
	hello := "hello" + " world"
	fmt.Println("hello",hello)

	//错误类型
	err := errors.New("EOF")
	fmt.Println("err",err)



	//分组定义类型
	type(
		nt int
		er string
	)
	//分组定义常量
	const(
		LENGTH = 20
		WIDTH = 40
	)

	//分组声明变量
	var(
		nut int = 4
		rut float64 = 3.14
	)
	fmt.Println("nut",nut,"rut",rut)
	//iota枚举
	const(
		a1 = iota
		a2 = iota
		a3
		a4
	)
	const a5 = iota
	fmt.Println("a1",a1,"a2",a2,"a3",a3,"a4",a4,"a5",a5)

	//每遇到一个const关键字，iota就会重置

}

