package main

import (
	"bufio"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
)
type Reply struct {
	Data string
}
func main() {
	client, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("err11!! : ",err)
	}
	log.Println("here : ",os.Stdin)
	c := jsonrpc.NewClient(client)
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal("err22!! : ",err)
		}
		var reply Reply
		err = c.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal("err33!! : ",err)
		}
		log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	}
}

