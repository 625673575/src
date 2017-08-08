package main

import (
	"path/filepath"
	"github.com/mozillazg/go-pinyin"
	"os"
	"path"
	"unicode"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}
func IsChineseChar(r rune) bool {
	if unicode.Is(unicode.Scripts["Han"], r) {
		return true
	}
	return false
}
func FirstToUpper(str string) string {
	return strings.ToUpper(str[0:1])+str[1:]
}
func main() {
	lenArg := len(os.Args)
	if lenArg < 2 {
		panic("Please specify a valid folder")
		os.Exit(1)
	}
	fold := os.Args[1]
	if IsExist(fold) {
		a := pinyin.NewArgs()

		filepath.Walk(fold, func(folds string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			ext := path.Ext(f.Name())

			if ext == ".mp3" || ext == ".wma" {
				fn := strings.TrimSuffix(f.Name(), ext)
				s := ""
				for _, v := range fn {
					if IsChineseChar(v) {
						t := string(v)
						s += FirstToUpper(pinyin.LazyPinyin(t, a)[0])
					} else {
						s += string(v)
					}
				}
				newName := s + ext
				if newName != f.Name() {
					newName = filepath.Dir(folds) + "/" + newName
					os.Rename(folds, newName)
				}
			}
			return nil
		})
	} else {
		panic("Not a valid folder")
	}
}
