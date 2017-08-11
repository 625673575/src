package file

import (
	"testing"
)

func Test_ListDisk(t *testing.T){
	t.Log(ListDisk())
}
func Test_ListDir(t *testing.T){
	t.Log(ListDirAndFile("D:\\behaviac"))
}
func Test_Rename(t *testing.T){
	Rename("D:\\备份","D:\\BackUp")
	Rename("D:\\behaviac\\version.txt","D:\\behaviac\\ver.txt")
}
func Test_ReadFile(t *testing.T){
	b,_:=ReadBytes("D:\\behaviac\\ver.txt")
	t.Log(b)
}
