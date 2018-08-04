package main

import(
	"fmt"
	"io"
	"errors"
	"io/ioutil"
)

func main(){
	fmt.Println("hello io api")
	var EOF = errors.New("EOF")
	fmt.Println(EOF,io.EOF,(EOF == io.EOF))

	var mybytes = [] byte {1,2,3}

	myread := MyReader{"chen"}
	fmt.Println(myread.Read(mybytes))

	v := interface{}(&myread)
	h,ok := v.(io.Reader)
	fmt.Println("h", h,"ok",ok)

	data := "test data say hello"
	ioutil.WriteFile("file.txt",[]byte(data),0644)
}

type MyReader struct{
	name string
}

func (*MyReader) Read(p []byte) (n int, err error){
	n = len(p)
	if(n  == 0 ){
		err = io.EOF
	}else{
		err = nil
	}
	return
}