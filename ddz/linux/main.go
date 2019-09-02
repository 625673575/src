package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -lstdc++ -L. -lddz
#include <c_api.h>
*/
import "C"
import (
	"fmt"
	"log"
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
func TrustShip(playCard []byte, lastPlayCards []byte, lastPlayIdentity byte, curPlayIdentity byte) []byte {
	ret := make([]byte, 20)
	la := byte(len(lastPlayCards))
	l0 := byte(len(playCard))
	lastCardPtr := (*C.uchar)(nil)
	if lastPlayIdentity != curPlayIdentity {
		lastCardPtr = (*C.uchar)(&lastPlayCards[0])
	}
	C.TrustShip((*C.uchar)(&ret[0]),(C.uchar)(l0),
		(*C.uchar)(&playCard[0]),
		(C.uchar)(la), lastCardPtr,
		(C.uchar)(lastPlayIdentity), (C.uchar)(curPlayIdentity))
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
	C.DealSmoothCard((C.uchar)(firstStep), (C.uchar)(secondStep),(*C.uchar)(&p0[0]), (*C.uchar)(&p1[0]), (*C.uchar)(&p2[0]), (*C.uchar)(&desk[0]))
	log.Println(FormatBytes(p0))
	log.Println(FormatBytes(p1))
	log.Println(FormatBytes(p2))
	log.Println(FormatBytes(desk))
	return p0, p1, p2, desk
}

func main() {
	i := []byte{0x14, 0x3f}
	log.Println(FormatBytes(i))
	p0, _, _, d := DealCard(0, 0)
	TrustShip(append(p0, d...),[]byte{}, Identity_Lord, Identity_Lord)
}
