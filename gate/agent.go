package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteByte(data []byte,id uint16)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
