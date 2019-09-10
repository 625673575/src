package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	wr, _ := template.ParseFiles("static/index.html")
	wr.Execute(w, nil)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := time.Now().Second()
	f, err := os.Create("message_" + strconv.Itoa(nowStr))
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte("name:" + r.FormValue("name") + "\n" +
			"email:" + r.FormValue("email") + "\n" +
			"message:" + r.FormValue("message")))
	}
}
