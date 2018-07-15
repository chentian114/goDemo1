package main

import "fmt"

func main() {
	fmt.Println("hello 流程控制")

	//if,选择
	str := "abc"
	if str == "abc"{
		fmt.Println("str == abc true")
	}else{
		fmt.Println("error")
	}

	// if 支持一个初始化语句
	if a:=20; a == 10{
		fmt.Println("a== 10 yes")
	}else if a == 20{
		fmt.Println("a==20 yes")
	}else {
		fmt.Println("error")
	}

	//switch 支持一个初始化语句
	switch num:=30; num {
	case 10:
		fmt.Println("num = 10")
		fallthrough //不跳出switch，执行后面的语句
	case 20,30,40:
		fmt.Println("num = 20 || 30 || 40") //默认有break  ,支持多个条件
	default:
		fmt.Println("other")
	}

	// 可以没有条件
	source := 80
	switch  {
	case source>= 90:
		fmt.Println(">=90")
	case source>=80:
		fmt.Println(">=80 && <90")
	default:
		fmt.Println("other")
	}

	//循环
	// for 初始化 ; 判断条件 ; 条件变化
	var sum int = 0
	for cnt:=10 ; cnt<20 ; cnt++ {
		sum += cnt
	}
	fmt.Println(" sum:",sum)

	// range 迭代
	// range 返回两个元素，第一个返回索引，第二个返回值
	strg := "abfg"
	for index,value := range strg {
		fmt.Printf("index %d value %c\n",index,value)
	}
	//第二个默认丢弃
	for index := range strg {
		fmt.Printf("index %d value：%c\n" , index,strg[index])
	}

	//第一个丢弃
	for _,value := range strg{
		fmt.Printf("value %c\n",value)
	}


	//跳转 break ,continue , goto
	var fnum int = 20;
	for{       // for永远为真
		if(fnum > 100 ) {
			break
		}
		fnum+=10
		if(fnum == 40){
			continue
		}
		fmt.Printf(" fnum : %d ",fnum)
	}
	fmt.Println("")
	fmt.Println("1111111111111111")
	goto MyTag
	fmt.Println("2222222222222222")
MyTag:
	fmt.Println("3333333333333333")

}

