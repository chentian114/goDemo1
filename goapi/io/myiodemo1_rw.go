package main

import (
	"fmt"
	"io"
)

func main(){
	fmt.Println("hello my io reader,writer")

	//Read 将 len(p) 个字节读取到 p 中。
}

type MyReader struct{
	data string
	len int
}

func (r *MyReader) Read(p []byte)(n int,err error){
	if r.len > len(p) {
	}
	return 0,io.EOF
}