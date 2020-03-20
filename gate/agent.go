package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteByte(data []byte)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
