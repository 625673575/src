package file

import (
	"os"
	"io/ioutil"
	"RPC/remote"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
func ListDisk() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return
}
func ListDirAndFile(path string) (remote.FileEntrys,error){
	var dirs []*remote.FileEntry
	files, err := ioutil.ReadDir(path)
	for _, f := range files {
		dirs = append(dirs,&remote.FileEntry{ IsDir: f.IsDir(), FullName: f.Name() ,ModeTime:f.ModTime().Unix()})
	}
	return remote.FileEntrys{Count:(int32)(len(dirs)),Entrys:dirs} ,err
}
func Rename(oldPath string,newPath string)bool{
	if IsExist(oldPath){
	err:=os.Rename(oldPath,newPath)
		if err==nil{
			return true
		}
	}
	return false
}
func ReadBytes(path string)([]byte,error){
return	ioutil.ReadFile(path)
}
func GetFileSize(path string)int64{
	f,err:=os.Stat(path)
	if err==nil{
		return f.Size()
	}
	return -1
}
func ReadString(path string)(string,error){
	b,err:=ioutil.ReadFile(path)
	return string(b),err
}