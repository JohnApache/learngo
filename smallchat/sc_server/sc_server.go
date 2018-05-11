package main

import (
	"encoding/binary"
	"encoding/json"
	"net"

	"gopkg.in/ffmt.v1"
)

type client chan<- string

var (
	enter = make(client)
	leave = make(client)
	msg   = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			ffmt.Mark(err)
			continue
		}
		go handdleConn(conn)
	}
}

var buf = make([]byte, 1024*1024)

func Send(conn net.Conn, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	ffmt.Mark(len(b))
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(b)))
	copy(buf[4:], b)
	ffmt.Puts(buf[:4+len(b)])
	_, err = conn.Write(buf[:4+len(b)])
	return err
}

func handdleConn(conn net.Conn) {
	//	ffmt.Mark(conn.RemoteAddr())
	//	for i := 0; i < 10; i++ {
	//		Send(conn, "asdasd")
	//		time.Sleep(1 * time.Second)
	//	}

}
