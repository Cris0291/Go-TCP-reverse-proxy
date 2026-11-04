package tunnel

import (
	"io"
	"net"
)

type ConnectionType int

const (
	ClientConnection ConnectionType = 0
	ServerConnection ConnectionType = 1
)

func ConnectionHandler(client net.Conn, server net.Conn) {
	go func() {
		io.Copy(server, client)
	}()

	go func() {
		io.Copy(client, server)
	}()
}
