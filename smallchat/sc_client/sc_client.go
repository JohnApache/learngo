package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net"

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
	//	ReadMsg(conn)
	var v string
	for {
		Receive(conn, &v)
		ffmt.Puts(v)
	}
}

func ReadMsg(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		bufs := bytes.NewBuffer(buf)
		for {
			_, err := conn.Read(buf)
			if err != nil {
				ffmt.Mark(err)
			}
			ffmt.Mark(string(buf[:18]))
			bufs.Read(buf)
		}
		ffmt.Mark(bufs.String())
	}
}

func Receive(conn net.Conn, v interface{}) error {
	_, err := conn.Read(Buffer)
	if err != nil {
		return err
	}
	l := binary.LittleEndian.Uint32(Buffer[:4])
	ffmt.Mark(l)
	_, err = conn.Read(BufferBody[:int(l)])
	if err != nil {
		return err
	}
	err = json.Unmarshal(BufferBody[:int(l)], v)
	if err != nil {
		return err
	}
	return nil
}
