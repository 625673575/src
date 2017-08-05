package main
import ("fmt"
	"syscall"
	"unsafe"
)
var (
	testdll               = syscall.NewLazyDLL("TestDLL.dll")
	procSetArray            = testdll.NewProc("SetArray")
	procAddMoney              = testdll.NewProc("AddMoney")
	procGetValueRef             = testdll.NewProc("GetValueRef")
	procGetArray             = testdll.NewProc("GetArray")
	procGetValue             = testdll.NewProc("GetValue")
)

type ByteRet *byte
func main(){
	var i int64 =12
	//ip:=&i
	fmt.Println(procAddMoney.Call())
retInputPointer,_,_:=	procGetValueRef.Call(uintptr(unsafe.Pointer(&i)))
	fmt.Println(retInputPointer)
	ar:=[]byte{1,2,4,3}
	procSetArray.Call(uintptr(unsafe.Pointer(&ar)),uintptr(4))
retGetArray,_,_:=	procGetArray.Call()
ps:=	unsafe.Pointer(retGetArray+unsafe.Sizeof(ar[0]))
bret:=	ByteRet(ps)
	fmt.Println(*bret)
vret,_,_:=	procGetValue.Call(uintptr(1))
	fmt.Println(vret)
}