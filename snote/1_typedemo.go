package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("hello go study note type! 类型")

	//使⽤关键字 var 定义变量，⾃动初始化为零值。如果提供初始化值，可省略变量类型，由编译器⾃动推断
	var x int = 32
	var f float32 = 3.12
	var s = "abc"
	fmt.Println("x:",x,"f:",f,"s:",s)

	//在函数内部，可⽤更简略的 ":=" ⽅式定义变量
	x2 := "123"
	fmt.Println("x2:",x2)

	//可⼀次定义多个变量。
	var num1,num2,num3 int = 1,2,3
	var n1,str = 4,"abc"
	fmt.Println("num1:",num1,"num2:",num2,"num3:",num3)
	fmt.Println("n1:",n1,"str:",str)

	n ,s := 0x12345 , "abc"
	fmt.Println("n:",n,"s:",s)

	//多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data , i := []int{1,2,3,4} , 0
	fmt.Println("data:",data,"i:",i,"data[i]",data[i])
	i , data[i] = 1,100
	fmt.Println("data:",data,"i:",i,"data[i]:",data[i])

	//特殊只写变量 "_"，⽤于忽略值占位
	_,str2 := test1()
	fmt.Println("str2",str2)

	//编译器会将未使⽤的局部变量当做错误
	// st := 20
	// 全局变量没问题。

	// 注意重新赋值与定义新同名变量的区别。
	str3 := "abc"
	num4 := 2
	fmt.Println("str3",str3,"&str3:",&str3,"num4:",num4,"&num4:",&num4)

	str3 = "def"
	fmt.Println("str3:",str3,"&str3:",&str3)
	{
		num4 := 5
		fmt.Println("num4:",num4,"&num4:",&num4)
	}

	//常量值必须是编译期可确定的数字、字符串、布尔值。
	// 未使⽤局部常量不会引发编译错误
	const x1,y1 int = 1 ,2
	const cstr = "hello world"
	const(
		 a, b = "a","b"
		 c bool = true
	)
	fmt.Println("x1:",x1,"y1:",y1,"cstr:",cstr)
	fmt.Println("a:",a,"b:",b,"c:",c)

	//在常量组中，如不提供类型和初始化值，那么视作与上⼀常量相同
	const(
		cstr2 = "abc"
		cstr3
	)
	fmt.Println("cstr2:",cstr2,"cstr3:",cstr3)

	//常量值还可以是 len、 cap、 unsafe.Sizeof 等编译期可确定结果的函数返回值。
	const(
		len1 = len("abc")
		cap1 = cap([5]int{1,2,3,4,5})
		size1 = unsafe.Sizeof("abc")
	)
	fmt.Println("len1:",len1,"cap1:",cap1,"size1:",size1)

	// 提供初始化表达式。
	aslice  :=[]int{0,0,0}
	aslice[0] = 10
	fmt.Println("aslice:",aslice)

	// makeslice
	bslice := make([]int,3)
	bslice[0] = 10
	fmt.Println("bslice:",bslice)

	// 类型转换
	//不⽀持隐式类型转换，即便是从窄向宽转换也不⾏
	var by byte = 10
	// var n int = by  //error
	var bn int = int(by)
	fmt.Println("by:",by,"bn:",bn)

	//字符串是不可变值类型，内部⽤指针指向 UTF-8 字节数组
	sb := "abc"
	fmt.Println("sb:",sb,"sb[0]:",sb[0],"sb[1]:",sb[1],"sb[2]:",sb[2])

	//使⽤ "`" 定义不做转义处理的原始字符串，⽀持跨⾏
	sc :=`a
	b\n\r
	c',""'`
	fmt.Println("sc:",sc)

	//连接跨⾏字符串时， "+" 必须在上⼀⾏末尾，否则导致编译错误
	sh := "hello" +
		"world"
	fmt.Println("sh:",sh,"&sh:",&sh)

	//⽀持⽤两个索引号返回⼦串。⼦串依然指向原字节数组，仅修改了指针和长度属性
	sh1 := sh[:5]
	sh2 := sh[1:]
	sh3 := sh[1:5]
	fmt.Println("sh1:",sh1,"sh2:",sh2,"sh3:",sh3)

	//修改字符串
	shb := []byte(sh1)
	fmt.Println("shb:",shb)
	shb[1]='E'
	sh = string(shb)
	fmt.Println("sh:",sh,"&sh:",&sh)

	shc := "中华人民共和国"
	fmt.Println("shc:",shc)
	shr := []rune(shc)
	fmt.Println("shr:",shr)
	shr[2] = '民'
	shr[3] = '族'
	shc = string(shr)
	fmt.Println("shc:",shc)

	shab := "abc汉字"
	for i=0 ; i<len(shab); i++{
		fmt.Print("   i:",i,"shab[i]:",shab[i])
	}
	fmt.Println(" ")
	for i,data := range shab {
		fmt.Printf("   i:%d data:%c",i,data)
	}
	fmt.Println(" ")

	type fdata struct{ p int}
	var pd = fdata{2}
	fmt.Println("pd:",pd,"&pd:",&pd,"pd.p:",pd.p)

	var pp *fdata
	pp = &pd
	fmt.Printf("pp %p, pp.p :%v",pp,pp.p)



}

func test1()(int,string)  {
	return 1,"abc"
}