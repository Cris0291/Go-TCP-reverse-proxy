package tunnel

import (
	"io"
	"net"
	"sync"
)

type ConnectionType int

const (
	ClientConnection ConnectionType = 0
	ServerConnection ConnectionType = 1
)

func ConnectionHandler(client net.Conn, server net.Conn) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		io.Copy(server, client)
		tcpConn, ok := client.(*net.TCPConn)

		if ok {
			tcpConn.CloseWrite()
		}
	}()

	go func() {
		defer wg.Done()

		io.Copy(client, server)
		tcpConn, ok := server.(*net.TCPConn)

		if ok {
			tcpConn.CloseWrite()
		}
	}()
}
