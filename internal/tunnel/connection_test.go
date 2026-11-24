package tunnel_test

import (
	"net"
	"testing"

	"github.com/Cris0291/myproxy/internal/tunnel"
)

func TestConnectionHandler(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("connection could not be established returned an error: %v", err)
	}

	serverConn, err := listener.Accept()
	if err != nil {
		t.Fatalf("server could not accept connection returned an error: %v", err)
	}

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("client could not accept connection returned an error: %v", err)
	}

	tunnel.ConnectionHandler(clientConn, serverConn)
}
