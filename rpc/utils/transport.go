/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/21 下午5:30
***********************************************/

package utils

import (
	"encoding/binary"
	"io"
	"net"
)

// Transport struct
type Transport struct {
	conn net.Conn
}

func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn: conn}
}

// Send data
func (tran *Transport) Send(req Data) (err error) {
	b, err := encode(req)
	if err != nil {
		return
	}
	buf := make([]byte, 4+len(b))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(b)))
	copy(buf[4:], b)
	_, err = tran.conn.Write(buf)
	return
}

// Receive data
func (tran *Transport) Receive() (data Data, err error) {
	header := make([]byte, 4)
	_, err = io.ReadFull(tran.conn, header)
	if err != nil {
		return
	}
	dataLen := binary.BigEndian.Uint32(header)
	dataByte := make([]byte, dataLen)
	_, err = io.ReadFull(tran.conn, dataByte)
	if err != nil {
		return
	}
	data, err = decode(dataByte)
	return
}
