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

// set timejumps per step
var myHour, myMinute, mySecond int = 0, 10, 0

// timout per emit
var timeOutInSec = 1
var startTime = time.Date(
	2019, 02, 01, 8, 0, 0, 651387237, time.UTC)

var connMap []*websocket.Conn

func main() {
	// emit to all connected
	go emitToAll()

	http.HandleFunc("/clock", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		connMap = append(connMap, conn)
	})
	http.ListenAndServe(":8080", nil)
}

func emitToAll() {
	Fstart := startTime
	upHour := 0
	upMinute := 0
	upSecond := 0
	for {
		start := Fstart.Local().Add(time.Hour*time.Duration(upHour) +
			time.Minute*time.Duration(upMinute) +
			time.Second*time.Duration(upSecond))
		upHour += myHour
		upMinute += myMinute
		upSecond += mySecond
		time.Sleep(time.Duration(timeOutInSec) * time.Second)
		for i := 0; i < len(connMap); i++ {
			connMap[i].WriteMessage(1, []byte(start.String()))
		}
	}
}
