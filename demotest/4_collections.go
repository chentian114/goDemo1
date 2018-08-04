package main

import "fmt"

func main(){
	fmt.Println("hello array slice map")
	//声明数组，5表示数组长度，int表示存储的元素类型
	var arri = [5]int{1,2,3,4,5}
	fmt.Println("arri len=",len(arri),"arri",arri)
	arri[4] = 7
	fmt.Println("arri",arri)
	//下标从0开始
	for i:=0 ; i<len(arri) ; i++{
		fmt.Print(arri[i])
	}
	fmt.Println("")

	for i,d := range arri {
		fmt.Print(i)
		fmt.Print(",")
		fmt.Print(d)
		fmt.Print(" ")
	}
	fmt.Println("")
	//长度也是数组类型的一部分，[4]int与[5]int是不同类型
	//数组不能改变长度，数组之间的赋值是值的赋值，而不是指针

	//声明并初始化
	var arr2 = [5]int{1,2,3,4,5}
	fmt.Println("arr2",arr2)
	//会自动计算长度
	arr3 := [...]int{4,5,6,7}
	fmt.Println("arr3",arr3)
	//多维数组，二行二列
	var arr4 = [2][2]int{{1,1},{1,2}}
	fmt.Println("arr4",arr4,"arr[0][0]",arr4[0][0])
	//slice并不是真正意义上的动态数组，而是一个引用类型
	//slice总是指向一个底层array
	//slice的声明类似array，只是不需要长度
	var slice []int
	slice = make([]int,5)
	slice[0] = 1
	fmt.Println("slice",len(slice),"slice[0]",slice[0])

	//slice可以从一个数组或一个已经存在的slice中再次声明
	sarr1 := [5]int{1,2,3,4,5}
	fmt.Println("sarr1",sarr1)
	var slice1 =  sarr1[2:4]
	fmt.Println("slice1",slice1,"len(slice1)",len(slice1),"cap(slice1)",cap(slice1))
	slice1[0] = 67
	slice1 = append(slice1, 15)
	fmt.Println("slice",slice1)
	fmt.Println("sarr1",sarr1)
	//数组和slice声明的不同
	//声明数组时，方括号内写明了数组的长度或使用...自动计算长度
	//声明slice时，方括号内没有任何字符
	var slice2 = []int{1,2,3,4,5}
	fmt.Println("slice2",slice2)

	//声明初始化两个slice
	sarr2 := [5]int{1,2,3,4,5}
	slice3 :=sarr2[2:len(sarr2)]
	fmt.Println("slice3",slice3)
	slice3 = append(slice3,20)
	slice3[0] = 20
	fmt.Println("sarr2",sarr2)
	fmt.Println("slice3",slice3)
	//slice是引用类型，改变j中的内容，i和k的内容也会变

	//map类型的声明,key是字符串，值是int
	//这种方式会创建一个nil map,所以在使用时必须用make初始化。
	var mp map[string]int = make(map[string]int)
	mp["hello"] = 10
	fmt.Println("mp",mp,"mp[hello]",mp["hello"])
	//另一种声明方式
	mp2 := make(map[string]int)
	mp2["world"] = 20
	fmt.Println("mp2[world]",mp2["world"])
	//声明并初始化
	mp3 := map[string]int{"hello":10,"world":20}
	fmt.Println("mp3",mp3)
	//map是无序的，长度不固定，引用类型
	//判断key是否存在
	//map有两个返回值，第二值表示key是否存在
	val,err := mp3["he"]
	val2,err2 := mp3["hello"]
	fmt.Println("val",val,"err",err)
	fmt.Println("val2",val2,"err2",err2)
}
