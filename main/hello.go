package main

import (
	"fmt"
	"chentian114/goDemo1/phone"
)
func main() {
	fmt.Println("Hello Wrold!")
	cnt := 10
	fmt.Println(cnt+10)

	arr := [3] int {1,2,3}
	fmt.Println(arr)

	phone.QueryPhone()
}