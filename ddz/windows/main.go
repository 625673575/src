package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

var (
	ddzdll         = syscall.NewLazyDLL("ddz.dll")
	TrustShip      = ddzdll.NewProc("TrustShip")
	DealSmoothCard = ddzdll.NewProc("DealSmoothCard")
)

const Identity_Lord = byte(0)
const Identity_Farmer1 = byte(1)
const Identity_Farmer2 = byte(2)

func FormatBytes(input []byte) string {
	x := ""
	for _, v := range input {
		x += fmt.Sprintf("%x ", uint(v))
	}
	return x
}

func TrustShipPlay(playCard []byte, lastPlayCards []byte, lastPlayIdentity byte, curPlayIdentity byte) []byte {
	ret := make([]byte, 20)
	la := len(lastPlayCards)
	l0 := len(playCard)
	lastCardPtr := uintptr(0)
	if lastPlayIdentity != curPlayIdentity {
		lastCardPtr = uintptr(unsafe.Pointer(&lastPlayCards[0]))
	}
	TrustShip.Call(uintptr(unsafe.Pointer(&ret[0])), uintptr(l0),
		uintptr(unsafe.Pointer(&playCard[0])),
		uintptr(la), lastCardPtr,
		uintptr(lastPlayIdentity), uintptr(curPlayIdentity))
	trimIndex := 20
	for i, v := range ret {
		if v == 0 {
			trimIndex = i
			break
		}
	}
	ret = ret[0:trimIndex]
	log.Println(FormatBytes(ret))
	return ret
}

func DealCard(firstStep byte, secondStep byte) ([]byte, []byte, []byte, []byte) {
	desk := make([]byte, 3)
	p0 := make([]byte, 17)
	p1 := make([]byte, 17)
	p2 := make([]byte, 17)
	DealSmoothCard.Call(uintptr(firstStep), uintptr(secondStep), uintptr(unsafe.Pointer(&p0[0])), uintptr(unsafe.Pointer(&p1[0])), uintptr(unsafe.Pointer(&p2[0])), uintptr(unsafe.Pointer(&desk[0])))
	log.Println(FormatBytes(p0))
	log.Println(FormatBytes(p1))
	log.Println(FormatBytes(p2))
	log.Println(FormatBytes(desk))
	return p0, p1, p2, desk
}
func ConvertBytes(input []byte) {
	for i, v := range input {
		input[i] = Convert(v)
	}
}
func Convert(input byte) byte {
	if input == 1 {
		return 0x4f
	}
	if input == 2 {
		return 0x5f
	}
	input -= 0x10
	if (input & 0xf) == 0xf {
		input -= 13
	}
	return input
}

func main() {
	i := []byte{0x14, 0x3f}
	ConvertBytes(i)
	log.Println(FormatBytes(i))
	p0, _, _, d := DealCard(0, 0)
	TrustShipPlay(append(p0, d...), []byte{0x19}, Identity_Farmer1, Identity_Lord)
}
