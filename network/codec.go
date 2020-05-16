package network

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/duration"
	"go-ratel/common"
	"net"
)

const (
	MaxContextLen = 4092
)

type Codec struct {
	Conn   *net.TCPConn
	Buffer Buffer
}

func NewCodec(conn *net.TCPConn) *Codec {
	return &Codec{
		Conn:   conn,
		Buffer: newBuffer(conn, MaxContextLen),
	}
}

func (c *Codec) Read() error {
	return c.Buffer.readFromReader()
}

func (c *Codec) Encode(transferData *common.ServerTransferDataProtoc, duration duration.Duration) error {
	encodeData, e := proto.Marshal(transferData)
	if e != nil {
		return e
	}
	bodyLen := len(encodeData)
	if bodyLen > MaxContextLen {
		return errors.New("not enough")
	}
	header := proto.EncodeVarint(uint64(bodyLen))

	buffer := make([]byte, len(header)+bodyLen)
	copy(buffer, header)
	copy(buffer[len(header):], encodeData)

	_, e = c.Conn.Write(buffer)
	return e
}

func (c *Codec) Decode() (*common.ServerTransferDataProtoc, bool, error) {
	bodyLen, size := proto.DecodeVarint(c.Buffer.buf[c.Buffer.start:])
	if bodyLen > MaxContextLen {
		return nil, false, errors.New("not enough")
	}
	if bodyLen == 0 {
		return nil, false, nil
	}
	body, e := c.Buffer.read(size, int(bodyLen))
	if e != nil {
		return nil, false, nil
	}

	transferData := common.ServerTransferDataProtoc{}
	e = proto.Unmarshal(body, &transferData)
	if e != nil {
		return nil, false, e
	}
	return &transferData, true, nil
}
