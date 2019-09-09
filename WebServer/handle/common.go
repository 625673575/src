package handle

import (
	"fmt"
	"strconv"
	"strings"
)

func convertCard(text string, jinzhi int) []int {
	var tx []int
	if text == "" {
		return tx
	}
	hasDouHao := strings.ContainsRune(text, ',')
	var sss []string
	if hasDouHao {
		sss = strings.Split(text, ",")
	} else {
		sss = strings.Split(text, " ")
	}
	for _, v := range sss {
		if v[0:2] == "0x" {
			v = v[2:]
		}
		hex, err := strconv.ParseInt(v, jinzhi, 64)
		if err == nil {
			tx = append(tx, int(hex))
		}
	}
	return tx
}

func stringArray(t []int) string {
	var r string
	l := len(t)
	for i, v := range t {
		r += fmt.Sprintf("%x", v)
		if i != l-1 {
			r += ","
		}
	}
	return r
}
func ToInts(x []byte) []int {
	r := make([]int, len(x))
	for i, v := range x {
		r[i] = int(v)
	}
	return r
}
func ToBytes(x []int) []byte {
	r := make([]byte, len(x))
	for i, v := range x {
		r[i] = byte(v)
	}
	return r
}
func PlaceAsHtmlInput(x string)string{
	return fmt.Sprintf("<input type=\"text\" value=\"%s\" size=50>",x)
}