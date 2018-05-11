package msg

import (
	"encoding/binary"
	"net"
)

var (
	BufLen  = make([]byte, 4)
	BufSend = make([]byte, 1024*1024)
)

func Send(conn net.Conn, m string) {
	l := len(m)
	binary.LittleEndian.PutUint16()
	conn.Write(BufSend)
}

func Receive() {

}
