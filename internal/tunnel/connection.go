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

func ConnectionHandler(client net.Conn, server net.Conn) (<-chan bool, <-chan error, <-chan error) {
	errCh1 := make(chan error)
	errCh2 := make(chan error)
	done := make(chan bool)

	go func() {
		defer close(done)
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()

			_, err := io.Copy(server, client)
			if err != nil {
				errCh1 <- err
			}

			tcpConn, ok := server.(*net.TCPConn)

			if ok {
				tcpConn.CloseWrite()
			}
		}()

		go func() {
			defer wg.Done()

			_, err := io.Copy(client, server)
			if err != nil {
				errCh2 <- err
			}

			tcpConn, ok := client.(*net.TCPConn)

			if ok {
				tcpConn.CloseWrite()
			}
		}()

		wg.Wait()
	}()

	return done, errCh1, errCh2
}
