package msg

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"

	"gopkg.in/ffmt.v1"
)

var (
	BufTransfer = make([]byte, 1024*1024)
)

func Send(conn net.Conn, m interface{}) error {
	v, err := json.Marshal(m)
	if err != nil {
		ffmt.Mark(err)
		return err
	}
	l := len(v)
	binary.LittleEndian.PutUint32(BufTransfer[:4], uint32(l))
	copy(BufTransfer[4:], v)
	_, err = conn.Write(BufTransfer[:4+l])
	if err != nil {
		return err
	}
	return nil
}

func Receive(conn net.Conn, v interface{}) error {
	_, err := conn.Read(BufTransfer[:4])
	if err != nil {
		return err
	}
	l := binary.LittleEndian.Uint32(BufTransfer[:4])
	if int(l) > 1024*1024 {
		return errors.New("消息体异常")
	}
	_, err = conn.Read(BufTransfer[4 : l+4])
	err = json.Unmarshal(BufTransfer[4:l+4], v)
	if err != nil {
		return err
	}
	return nil
}
