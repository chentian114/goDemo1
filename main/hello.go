package main

import (
	"fmt"
	"github.com/chentian114/goDemo1/phone"
	"github.com/chentian114/goDemo1/verify"
)
func main() {
	fmt.Println("Hello Wrold!")
	cnt := 10
	fmt.Println(cnt+10)

	arr := [3] int {1,2,3}
	fmt.Println(arr)

	phone.QueryPhone()
	fmt.Println(verify.CreateDefaultRandCode())
	fmt.Println(verify.CreateRandCode(11))
	for i:=0;i<10 ;i++  {
		fmt.Println(verify.GetRand())
	}
}