package main

import (
	"mypkg/smallchat/sc_msg"
	"net"

	"gopkg.in/ffmt.v1"
)

type client chan<- interface{}

var (
	enter   = make(chan client)
	leave   = make(chan client)
	message = make(chan interface{})
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	go BroadCaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			ffmt.Mark(err)
			continue
		}
		go handdleConn(conn)
	}
}

func handdleConn(conn net.Conn) {
	defer conn.Close()
	ffmt.Mark(conn.RemoteAddr())
	cli := make(chan interface{})
	go SendMsg(conn, cli)
	message <- conn.RemoteAddr()
	enter <- cli
	ReceiveMsg(conn)
	leave <- cli
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
		message <- v
		ffmt.Puts(conn.RemoteAddr(), v)
	}
}

func SendMsg(conn net.Conn, ch chan interface{}) {
	defer conn.Close()
	for m := range ch {
		err := msg.Send(conn, m)
		if err != nil {
			ffmt.Mark(err)
		}
	}
}

func BroadCaster() {
	clients := make(map[client]bool)
	for {
		select {
		case m := <-message:
			for cli := range clients {
				cli <- m
			}
		case cli := <-enter:
			clients[cli] = true
		case cli := <-leave:
			delete(clients, cli)
			close(cli)
		}
	}
}
