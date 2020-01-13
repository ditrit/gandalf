package message

import(
	"github.com/pebbe/zmq4"
)
//TODO FINIR
type Message interface {
	GetUUID() string
	GetTimeout() string
	SendWith(socket *zmq4.Socket, header string) (isSend bool)
	SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool)
	SendMessageWith(socket *zmq4.Socket) (isSend bool)
}