package verify

import (
	"math/rand"
	"time"
	"sync"
)
const (
	SOURCE_STRING = "aAcCdDeEfFgGhHiIkKmMnNpPqQxXyYzZ123456789"
 	MAX_LENGTH  = len(SOURCE_STRING)
 	DEFAULT_LENGTH  = 4
)
var (

)
func CreateDefaultRandCode()(result string){
	len := DEFAULT_LENGTH
	for i:=0 ; i<len ; i++ {
		idx := GetRand()
		result += string(SOURCE_STRING[idx])
	}
	return
}
func CreateRandCode(len int)(result string){
	if len>MAX_LENGTH {
		len = MAX_LENGTH
	}
	for i:=0 ; i<len ; i++ {
		idx := GetRand()
		result += string(SOURCE_STRING[idx])
	}
	return
}

func GetRand() int{
	rand.Seed(time.Now().UnixNano())
	var mu  sync.Mutex
	mu.Lock()
	n := rand.Intn(MAX_LENGTH)
	mu.Unlock()
	return n
}
