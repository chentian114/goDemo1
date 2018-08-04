package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("hello strings api")
	// 判断两个utf-8编码字符串，大小写不敏感
	a, b :="hello world", "Hello World"
	res := strings.EqualFold(a, b)
	fmt.Println("equal:",res)

	// 判断s是否有前缀字符串prefix
	fmt.Println(strings.HasPrefix(a,"hello"))
	fmt.Println(strings.HasPrefix(a,"heello"))

	// 判断s是否有后缀字符串suffix
	fmt.Println(strings.HasSuffix(a,"old"))
	fmt.Println(strings.HasSuffix(a,"orld"))

	// 判断字符串s是否包含子串substr
	fmt.Println(strings.Contains(a,"world"))
	fmt.Println(strings.Contains(a,"word"))


}
