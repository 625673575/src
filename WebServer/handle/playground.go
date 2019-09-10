package handle

import (
	"ddz"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"html/template"
	"io"
	"log"
	"net/http"
)

var IDENTITY=map[string]int32{"Lord":0,"Farmer1":1,"Farmer2":2}
func PlayGroundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/play.html")
		log.Println(t.Execute(w, nil))
	} else {
		ip := r.FormValue("ip")
		input_type := r.FormValue("input_type")
		jinzhi := 16
		if input_type == "10进制" {
			jinzhi = 10
		}
		l := convertCard(r.FormValue("l"), jinzhi)
		f1 := convertCard(r.FormValue("f1"), jinzhi)
		f2 := convertCard(r.FormValue("f2"), jinzhi)

		last := convertCard(r.FormValue("last"), jinzhi)
		player_identity := r.FormValue("player_identity")
		last_identity := r.FormValue("last_identity")

		io.WriteString(w,"<html><body>")
		io.WriteString(w,"<font size=\"5\">")
		io.WriteString(w, "ip:"+ip+
			"<br>Lord:"+stringArray(l)+
			"<br>Farmer1:"+stringArray(f1)+
			"<br>Farmer2:"+stringArray(f2)+
			"<br>Last:"+stringArray(last)+
			"<br>Player Identity:"+player_identity+
			"<br>Last Identity:"+last_identity+
			"<br>")

		conn, err := grpc.Dial("47.75.241.11:50001", grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect :", err.Error())
		}
		defer conn.Close()
		c := ddz.NewRobotServiceClient(conn)

		lordCard := ToBytes(l)
		farmer1Card := ToBytes(f1)
		farmer2Card := ToBytes(f2)

		lastCards := []byte{0x1e, 0x1e, 0x13, 0x23, 0x33}
		ret, err := c.Play(context.Background(), &ddz.RobotRequest{
			Playeridentity:  IDENTITY[player_identity],
			LordHandcard:    lordCard,
			Farmer1Handcard: farmer1Card,
			Farmer2Handcard: farmer2Card,
			LastIdentity:    IDENTITY[last_identity],
			LastPlaycard:    lastCards,
		})
		if err != nil {
			io.WriteString(w, "Error happened:<br>")
			io.WriteString(w, err.Error())
		} else {
			io.WriteString(w, "The result is :<br>")
			io.WriteString(w, fmt.Sprintf("HandCard : %x<br>", ToInts(ret.Handcard)))
			io.WriteString(w, fmt.Sprintf("CardStyle: %s<br>", ret.Style.String()))
		}
		io.WriteString(w,"</font>")
		io.WriteString(w,"</body></html>")
	}
}

