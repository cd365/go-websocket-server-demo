package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var Ws1 *Hub

func init() {
	Ws1 = NewHub()
	go Ws1.Run()
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(time.Millisecond * 200)
		for range ticker.C {
			Ws1.broadcast <- []byte(fmt.Sprintf("%s %d", time.Now().Format(time.ANSIC), rand.Int63()))
		}
	}()
}

func UpgradeToWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.register <- client
	go client.WritePump()
	go client.ReadPump()
}
