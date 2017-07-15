package  main

import (
	"fmt"
	"time"
)
func main() {
	for i:=0;i<5;i++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("from hello")
	}
}



