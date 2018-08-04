package main

import(
	"fmt"
	"strings"
	"io"
	"os"
)

const (
	STRING_INPUT =  "1"
	FILE_INPUT	 = "2"
	COMMAND_INPUT = "3"
	EXIT = "q"
)

func main(){
	fmt.Println("hello Reader demo")
	for{
		printMenu();
		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err error
		)
		switch strings.ToLower(ch) {
		case STRING_INPUT:
			helloStr := "hello World!"
			data,err = ReadFrom(strings.NewReader(helloStr),len(helloStr))
		case FILE_INPUT:
			file,_ := os.Open("file.txt")
			data,err = ReadFrom(file,20)
			file.Close()
		case COMMAND_INPUT:
			fmt.Println("请输入数据，不超过9个字符:")
			data, err = ReadFrom(os.Stdin, 11)
		case EXIT:
			fmt.Println("程序退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入！")
			continue
		}

		if err != nil {
			fmt.Println("error ",err)
		}else {
			fmt.Printf("data %s\n",data)
		}
	}
}
func ReadFrom(reader io.Reader,len int)([]byte,error){
	p := make([]byte,len)
	n,err := reader.Read(p)
	if n >0 {
		return p[:n],err
	}
	return p,err
}

func printMenu(){
	fmt.Println("--------------------------------")
	fmt.Println("请选择Reader读取方式：")
	fmt.Println("输入1：读取字符串")
	fmt.Println("输入2：读取文件")
	fmt.Println("输入3：标准输入")
	fmt.Println("输入q：退出")
}
