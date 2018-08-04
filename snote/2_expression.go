package main

import "fmt"

func main()  {
	fmt.Println("hello expression! 表达式")

	//初始化复合对象，必须使⽤类型标签，且左⼤括号必须在类型尾部。
	var st1 = struct{ a int}{1}
	var sl1 = []int{1,2,3,4}
	fmt.Println("st1:",st1,"sl1:",sl1)

	//初始化值以 "," 分隔。可以分多⾏，但最后⼀⾏必须以 "," 或 "}" 结尾
	var sl2 = []int{
		1,
		2,
	}
	fmt.Println("sl2:",sl2)

	var sl3 = []int{
		1,
		2}
	fmt.Println("sl3:",sl3)

	//可省略条件表达式括号。
	//⽀持初始化语句，可定义代码块局部变量。
	//代码块左⼤括号必须在条件表达式尾部。
	i:=0
	if n:="abc" ; i>0 {
		fmt.Println(n[2])
	}else if(i<0){
		fmt.Println(n[0])
	}else{
		fmt.Println(n[i])
	}

	str1 := "abc"
	for i,n := 0,len(str1) ; i< n; i++{
		fmt.Print(string(str1[i])+" ")
	}
	fmt.Println("")

	n1 := len(str1)
	for n1 > 0 {
		fmt.Println("str[",n1,"]:",string(str1[n1-1]))
		n1--
	}

	for i:= range str1{
		fmt.Println("i:",str1[i])
	}

	for i,v := range str1{
		fmt.Println("i:",i,"v:",v)
	}

	m := map[string]int {"a":1,"b":2}
	fmt.Println("m:",m,"m[a]",m["a"])

	for k,v := range m {
		fmt.Println("k",k,"v",v)
	}

	//分⽀表达式可以是任意类型，不限于常量。
	str3 := "abcd"
	for _,v := range str3{
		switch string(v) {
		case "a":
			fmt.Println("is a")
		case "b":
			fmt.Println("is b")
		case "c":
			fmt.Println("is c")
		default:
			fmt.Println("other")
		}
	}


}

