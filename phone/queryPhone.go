package phone

import (
	"fmt"
	"github.com/xluohome/phonedata"
)
func QueryPhone(){
	phone:="18142816065"
	fmt.Println(phone)

	pr, err :=phonedata.Find(phone)
	if err != nil {
		panic(err)
	}
	fmt.Println(pr)
}


