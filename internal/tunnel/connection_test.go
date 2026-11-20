package tunnel_test

import (
	"net"
	"testing"
)

func TestConnectionHandler(t *testing.T) error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	serverConn, err := listener.Accept()
	if err != nil {
		return err
	}

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		return err
	}
}
