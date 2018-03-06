package main

import (
	"log"



	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"time"
)


func main() {

	client, err := gosocketio.Dial(
		gosocketio.GetUrl("hack.multy.io/socket.io/", 80, false),
		transport.GetDefaultWebsocketTransport(),
	)


	client.On("hi",func(c *gosocketio.Channel, args interface{}) {
		//client id is unique
		log.Println("Connected to the server with ID:", args)

		//log.Println(amount, "clients in room")
	})

	log.Printf("Connection done:",  err)

	sendWithAck(client,  "event:receiver:on", "Check my awesome data!")

	time.Sleep(10* time.Second)
	client.Close()
	log.Printf("DONE")
}

func sendWithAck(c *gosocketio.Client, event string, value string){
	log.Printf("Send with ack called")
	result, err := c.Ack(event, value, time.Second * 5)
	if err != nil {
		log.Printf("Error:", err)
	} else {
		log.Printf("Got Ack from Server:", result)
	}
}
