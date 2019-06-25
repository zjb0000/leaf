package network

import (
	"net"
	"time"
)

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
