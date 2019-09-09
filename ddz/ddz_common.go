package ddz

import "fmt"

const Identity_Lord = byte(0)
const Identity_Farmer1 = byte(1)
const Identity_Farmer2 = byte(2)

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

func FormatBytes(input []byte) string {
	x := ""
	for _, v := range input {
		x += fmt.Sprintf("%x ", uint(v))
	}
	return x
}

