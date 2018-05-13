package main

import (
	"bufio"
	"mypkg/smallchat/sc_msg"
	"net"
	"os"

	_ "github.com/gorilla/websocket"
	"gopkg.in/ffmt.v1"
)

var Buffer = make([]byte, 4)
var BufferBody = make([]byte, 1024*1024)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	go ReceiveMsg(conn)
	SendMsg(conn)
}
func SendMsg(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		err := input.Err()
		if err != nil {
			ffmt.Mark(err)
			break
		}
		err = msg.Send(conn, input.Text())
		if err != nil {
			ffmt.Mark(err)
			break
		}
	}
	err := input.Err()
	if err != nil {
		ffmt.Mark(err)
	}
}

func ReceiveMsg(conn net.Conn) {
	defer conn.Close()
	var v interface{}
	for {
		err := msg.Receive(conn, &v)
		if err != nil {
			ffmt.Mark(err)
			break
		}
		ffmt.Puts(v)
	}
}
