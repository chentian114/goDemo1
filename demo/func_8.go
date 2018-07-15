package main

import "fmt"

func main(){
	fmt.Println("hello 函数")

	MyFunc1()

	//调用函数传递的参数叫实参
	MyFunc2(12,"hello")

	MyFunc3(1,2)

	MyFunc4(1,2,3,4,5)

	fmt.Println(MyFunc5())
	fmt.Println(MyFunc6())
	num1,str1 := MyFunc7()
	fmt.Println("num1",num1,"str1",str1)
	num2,str2 :=MyFunc8()
	fmt.Println("num2",num2,"str2",str2)

	a1,b2,c3 :=MyFunc9()
	fmt.Println("a1",a1,"b2",b2,"c3",c3)
	a2,_,_ := MyFunc9()  //丢弃第2，第3个返回值
	fmt.Println("a2",a2)

	fmt.Println("max(20,40)",MyFunc10Max(20,40))

	MyFunc11(5)

	fmt.Printf("sum 1+...+%d = %d\n",100,MyFunc12Sum(100))


	var func1 MyFuncType = Add
 	var func2 MyFuncType = Sub
	fmt.Println("add(1,2)=",func1(1,2))

	fmt.Println("calc add(1,2)",Calc(1,2,func1))
	fmt.Println("calc sub(2,1)",Calc(2,1,func2))

	//匿名函数
	 var neFunc = func(a,b int)(result int){
		result = a+b
		return
	}

	fmt.Println(neFunc(1,2))

	//定义并执行
	re := func(a,b int)(result int){
		fmt.Println("a",a," b",b)
		result = a + b
		return
	}(2,3)
	fmt.Println("re",re)


	//函数闭包
	//一个函数“捕获”了和它在同一作用域的其它常量和变量
	//这种意味着当闭包调用的时候，不管在程序什么地方调用，闭包能使用这些常量和变量
	//它不管这些变量和常量是否超出了作用域，所以只要有闭包在使用它，它就还会存在

	var bnum int = 20
	func(){
		fmt.Println("bnum:",bnum)
	}()

	nfu := closeVar(10,20)
	fmt.Println(nfu())
	fmt.Println(nfu())
	fmt.Println(nfu())

	//defer 延迟调用到函数结束的时候
	//多个defer时候 倒序，先进后出
	defer fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println("bbbbbbbbbbbbb")
	defer fmt.Println("cccccccccccccccc")
	defer fmt.Println("ddddddddddddddddddd")


	//defer与匿名函数
	defer func(){
		fmt.Println("abc")
	}()
}

//1. 无参无返回值函数
func MyFunc1(){
	fmt.Println("MyFunc1")
}

//2. 有参无返回值函数
//定义函数()后定义的参数叫形参
//参数传递只能由实参传递给形参
func MyFunc2(num1 int,str string)  {
	fmt.Println("num1:",num1,"str:",str)
}

func MyFunc3(num1,num2 int){
	fmt.Println("num1",num1,"num2",num2)
}

//3. 不定参数列表 不定参数,只能放在形参中的最后一个
func MyFunc4(args ... int){
	for i:=0 ; i<len(args) ; i++{
		fmt.Printf("arg[%d] = %d ",i,args[i])
	}
	fmt.Println("")

	for idx,data := range args{
		fmt.Printf("arg[%d] = %d ",idx,data)
	}
	fmt.Println("")
}

//4. 无参有一个返回值
func MyFunc5() int{
	return 100
}

func MyFunc6() (result int){
	result = 1000
	return;
}

//5. 有多个返回值
func MyFunc7() (int,string){
	return 100,"Yes"
}
func MyFunc8() (num int,desc string){
	num = 100
	desc = "Yes"
	return
}
func MyFunc9()(a , b , c int){
	a = 1
	b = 2
	c = 3
	return
}


//6.有传参有返回值
func MyFunc10Max(num1,num2 int)(maxNum int)  {
	if( num1 > num2){
		maxNum = num1
	}else{
		maxNum = num2
	}
	return
}

//7.函数递归调用
func MyFunc11(data int){
	if(data==0){
		return
	}
	MyFunc11(data-1)
	fmt.Println("data:",data)
}

func MyFunc12Sum(num int)(sum int){
	if(num==0){
		return 0
	}
	sum = MyFunc12Sum(num-1)+num
	return
}

//8.函数类型
func Add(a,b int)int  {
	return a+b
}
func Sub(a,b int)int{
	return a-b
}
type MyFuncType func(int,int)int

//9. 回调函数 函数有一个参数是函数类型，这个函数就是回调函数
//实现多态 MyFuncType
func Calc(num1,num2 int , fnc MyFuncType)(result int){
	fmt.Println("Calc")
	result = fnc(num1,num2)
	return
}


//闭包
func closeVar(num1,num2 int) func() int {
	var result int
	result += num1
	result += num2
	return func() int{
		result++
		return result
	}
}