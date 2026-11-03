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

func connectionHandler(client net.Conn, server net.Conn, connDirection ConnectionType) <-chan byte {
	dataStream := make(chan byte)
	go func() {
		if connDirection == ClientConnection {
			io.Copy(server, client)
		} else {
			io.Copy(client, server)
		}
	}()

	return dataStream
}
