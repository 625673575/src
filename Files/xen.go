package  main

import (
	"fmt"
"reflect"
)
const a="fs"
func main() {
	for i:=0;i<5;i++ {
		v:=reflect.ValueOf(a)

		v.SetCap(100)
		fmt.Println(v)
	}
}



