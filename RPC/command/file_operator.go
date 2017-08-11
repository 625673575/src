package command

import (
	"fmt"
	"RPC/file"
)

func renameFile(args []string) string {
	if len(args) < 2 {
		return "参数不正确"
	}
	if file.Rename(args[0], args[1]) {
		return fmt.Sprintf("修改%s为%s", args[0], args[1])
	}else {
		return "修改文件名失败"
	}
}
