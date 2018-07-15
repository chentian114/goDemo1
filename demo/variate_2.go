package main

import "fmt"

func main(){
	fmt.Println("hello 变量：程序运行期间，可改变的量")
	var cnt int
	fmt.Println(cnt)

	var str1,str2 string
	fmt.Println(str1+" "+str2)

	var count int = 2
	fmt.Println(count)

	//自动推导
	flag := true
	fmt.Println(flag)

	fmt.Println("flag",flag)

	str1 = "chen"
	fmt.Println("count:",count,"flag:",flag)
	fmt.Printf("count %d,str1 %s\n",count,str1)

	//多重变量赋值
	a , b := 10 , 20
	fmt.Printf("a %d b %d\n",a,b)

	//变量值互换
	a , b = b , a
	fmt.Printf("a %d b %d\n",a,b)

	// _匿名变量，丢弃，不处理,配合函数返回值使用
	tmp,_ := 10,20
	fmt.Println(tmp)

	cnt1,cnt2,_ := test()
	fmt.Println("cnt1",cnt1,"cnt2",cnt2)


	fmt.Println("hello")
}

func test()(a int,b int,c int){
	return 1,2,3
}