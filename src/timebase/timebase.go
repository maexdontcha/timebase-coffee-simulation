// websockets.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var p = fmt.Println
var myHour, myMinute, mySecond int = 0, 10, 0
var timeOutInSec = 2
var startTime = time.Date(
	2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

var connMap []*websocket.Conn

func main() {

	go emmitToAll()
	// append works on nil slices.
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		connMap = append(connMap, conn)
		// fmt.Println(connMap)
		// msgType, msg, err := conn.ReadMessage()
		// fmt.Printf("connect")
		// for {
		// 	// Read message from browser

		// 	if err != nil {
		// 		return
		// 	}

		// 	// Print the message to the console
		// 	// fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 	// Write message back to browser
		// 	if err = conn.WriteMessage(msgType, msg); err != nil {
		// 		return
		// 	}
		// }
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "websockets.html")
	// })

	http.ListenAndServe(":8080", nil)

	fmt.Println("abhier1")

}

func emmitToAll() {
	Fstart := startTime
	upHour := 0
	upMinute := 0
	upSecond := 0
	// fmt.Println(start)
	for {
		start := Fstart.Local().Add(time.Hour*time.Duration(upHour) +
			time.Minute*time.Duration(upMinute) +
			time.Second*time.Duration(upSecond))
		upHour += myHour
		upMinute += myMinute
		upSecond += mySecond
		time.Sleep(time.Duration(timeOutInSec) * time.Second)
		fmt.Println(start)
		for i := 0; i < len(connMap); i++ {
			// fmt.Println(connMap[i].RemoteAddr())
			connMap[i].WriteMessage(1, []byte(start.String()))
		}
	}
}
