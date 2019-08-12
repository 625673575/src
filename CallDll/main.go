package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

var (
	testdll         = syscall.NewLazyDLL("TestDLL.dll")
	procSetArray    = testdll.NewProc("SetArray")
	procAddMoney    = testdll.NewProc("AddMoney")
	procGetValueRef = testdll.NewProc("GetValueRef")
	procGetArray    = testdll.NewProc("GetArray")
	procGetValue    = testdll.NewProc("GetValue")
)

type ByteRet *byte

func Test0() {
	var i int64 = 12
	//ip:=&i
	fmt.Println(procAddMoney.Call())
	retInputPointer, _, _ := procGetValueRef.Call(uintptr(unsafe.Pointer(&i)))
	fmt.Println(retInputPointer)
	ar := []byte{1, 2, 4, 3}
	procSetArray.Call(uintptr(unsafe.Pointer(&ar)), uintptr(4))
	retGetArray, _, _ := procGetArray.Call()
	ps := unsafe.Pointer(retGetArray + unsafe.Sizeof(ar[0]))
	bret := ByteRet(ps)
	fmt.Println(*bret)
	vret, _, _ := procGetValue.Call(uintptr(1))
	fmt.Println(vret)
}

var (
	throwdll        = syscall.NewLazyDLL("ThrowDll.dll")
	LiterallyThrow  = throwdll.NewProc("LiterallyThrow")
	CatchIndexError = throwdll.NewProc("CatchIndexError")
	RunIndexError   = throwdll.NewProc("RunIndexError")
	NoError         = throwdll.NewProc("NoError")
)

func Test1() {
	NoError.Call()
	//LiterallyThrow.Call()
	//CatchIndexError.Call()
	//RunIndexError.Call()
}

var (
	ddzdll         = syscall.NewLazyDLL("ddz.dll")
	RobotPlay      = ddzdll.NewProc("RobotPlay")
	DealSmoothCard = ddzdll.NewProc("DealSmoothCard")
	Add            = ddzdll.NewProc("Add")
	AddBytes 	   = ddzdll.NewProc("AddBytes")
)

func TestRobotPlay() []byte {
	ret := make([]byte, 20)
	p0 := []byte{0x14, 0x24, 0x16, 0x48}
	p1 := []byte{0x1}
	p2 := []byte{0x2}
	last := []byte{0x2f}
	lastPlayIdentity := byte(0)
	curPlayIdentity := byte(0)
	l0 := byte(4)
	l1 := byte(1)
	l2 := byte(1)
	RobotPlay.Call(uintptr(unsafe.Pointer(&ret[0])),uintptr(l0), uintptr(l1), uintptr(l2), uintptr(unsafe.Pointer(&p0[0])),
		uintptr(unsafe.Pointer(&p1[0])),
		uintptr(unsafe.Pointer(&p2[0])),
		uintptr(l2), uintptr(unsafe.Pointer(&last[0])),
		uintptr(lastPlayIdentity), uintptr(curPlayIdentity))
	log.Println(ret)
	return ret
}

func TestDealCard() {
	desk := make([]byte, 3)
	p0 := make([]byte, 17)
	p1 := make([]byte, 17)
	p2 := make([]byte, 17)
	DealSmoothCard.Call(uintptr(unsafe.Pointer(&p0[0])), uintptr(unsafe.Pointer(&p1[0])), uintptr(unsafe.Pointer(&p2[0])), uintptr(unsafe.Pointer(&desk[0])))
	log.Println(p0)
	log.Println(p1)
	log.Println(p2)
	log.Println(desk)
}
func TestAddBytes(){
	x:=[]byte{15,16,18,12}
	r,_,_:=AddBytes.Call(uintptr(unsafe.Pointer(&x[0])),4)
	log.Println(r)
}
func main() {
	a, _, _ := Add.Call(uintptr(5), uintptr(6))
	log.Println(a)
	TestAddBytes()
	TestDealCard()
	TestRobotPlay()
}
