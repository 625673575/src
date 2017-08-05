package command

import (
	"github.com/go-vgo/robotgo"
	"fmt"
)

func getScreenSize(args []string) string {
	x, y := robotgo.GetScreenSize()
	return fmt.Sprintf("Resolution=x:%d  y:%d", x, y)
}