package gate

import (
	"net"
	"time"
)

type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
