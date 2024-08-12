package socket

import "github.com/zishang520/socket.io/v2/socket"

func NewSocketIo() *socket.Server {
	socketIo := socket.NewServer(nil, nil)

	socketIo.On("connection", func(clients ...interface{}) {

		client := clients[0].(*socket.Socket)

		client.On("message", func(args ...interface{}) {
			client.Emit("message-back", args...)
		})

		client.Emit("auth", client.Handshake().Auth)

		client.On("message-with-ack", func(args ...interface{}) {
			ack := args[len(args)-1].(func([]any, error))
			ack(args[:len(args)-1], nil)
		})
	})

	socketIo.Of("/custom", nil).On("connection", func(clients ...interface{}) {
		client := clients[0].(*socket.Socket)
		client.Emit("auth", client.Handshake().Auth)
	})
	return socketIo
}
