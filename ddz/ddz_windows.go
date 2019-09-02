package ddz

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	ddzdll         = syscall.NewLazyDLL("ddz.dll")
	TrustShip      = ddzdll.NewProc("TrustShip")
	DealSmoothCard = ddzdll.NewProc("DealSmoothCard")
)


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
	return p0, p1, p2, desk
}
