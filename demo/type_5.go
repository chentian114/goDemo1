package main

import "fmt"

func main() {

	fmt.Println("hello type")
    // bool 类型 1个字节 零值 false
	var a bool
	a = true
	fmt.Println("a= ",a)

	// %v 万能替换符
	var b = false
	fmt.Printf("b = %t %v \n",b,b)

	c := true
	fmt.Println("c",c)

	var d bool
	fmt.Println("d",d)

	// 浮点型

	var f1 float32
	f1 = 3.14
	fmt.Println("f1",f1)

	f2 := 3.14
	fmt.Println("f2",f2)
	fmt.Printf("f2 %f type %T\n",f2,f2)


	var ch byte = 97
	fmt.Printf("ch %d %c %T\n", ch,ch,ch)

	// 字符串类型
	var str  string
	str = "abc"
	fmt.Println("str",str,"len",len(str))

	var ca = 'a'
	fmt.Printf("byte %d %T", ca , ca)
	fmt.Println("str[0]",str[0],"str[1]",str[1])

	// 复合类型
	var comp complex64
	comp = 12 + 13i
	fmt.Println("comp",comp)

	fmt.Println("real(comp)",real(comp),"imag(comp)",imag(comp))

	// 输入 阻塞等待输入
	var scn int
	fmt.Printf("请输入：")
	fmt.Scanf("%d",&scn)
	fmt.Println("scn=",scn)
	fmt.Scan(&scn)
	fmt.Println("scn=",scn)

	//类型
	// bool 不能相互转换整型 ， 不兼容类型
	var flag = true
	fmt.Printf(" flag : %t \n", flag)
	//fmt.Printf("flag %d" , int(flag))  //error

	// 兼容类型 ,可以相互转换
	var chb byte = 'a'
	var cn int = 12

	cn = int(chb)
	fmt.Printf("ch %c %d %d\n" , chb,chb,cn)

	//类型别名 type
	type mybigint int64	//为int64取一个别名bigint
	var bcnt mybigint = 20
	fmt.Printf("bcnt %d %T\n",bcnt,bcnt)

	//多个类型别名
	type (
		mystr string
		myint int
	)
	var mstr mystr = "abc"
	var mint myint = 12
	fmt.Println(" mstr=",mstr,"mint=",mint)


	}