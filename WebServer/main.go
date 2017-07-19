package main

import (
	"net/http"
	"io"
	"fmt"
	"strings"
	"html/template"
	"log"
	"time"
	"crypto/md5"
	"strconv"
	"os"
	"path/filepath"
	"os/exec"
	"bufio"
)

type Hello struct {
	text string
}
type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

var count int = 1

func (Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "liu!\n")
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)

		fmt.Println("val:", strings.Join(v, ""))
	}
	wr, _ := template.ParseFiles("WebServer/login.html")
	wr.Execute(w, nil)

}
func writeFile(fileName string, content string) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	dstFile.WriteString(content + "\n")
}
func execCmdGoRun(fileName string, w http.ResponseWriter) {
	cmd := exec.Command("go", "run", fileName)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("error occur",err)
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(line)
			io.WriteString(w, line)
		}
	cmd.Wait()
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("WebServer/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		msg := r.FormValue("msg")
		io.WriteString(w, username+"  "+password)
		writeFile("./Files/xen.go", msg)
		path, _ := filepath.Abs("./Files/xen.go")
		path = strings.Replace(path, "\\", "/", -1)
		defer execCmdGoRun(path, w)
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("WebServer/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./Files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func init() {
	println("__int")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)

	http.ListenAndServe("localhost:4000", mux)
	// http.ListenAndServe("localhost:4000", http.FileServer(http.Dir(".")))
}
