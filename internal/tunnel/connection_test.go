package tunnel_test

import (
	"net"
	"testing"

	"github.com/Cris0291/myproxy/internal/tunnel"
	"golang.org/x/tools/go/callgraph/cha"
)

func TestConnectionHandler(t *testing.T) {
	test_client, proxy_server, err := createConnectionPair()

  if err != nil{
		t.Fatalf("client-proxy connection could not be established returned an error: %v", err)
	}

  test_server, proxy_client, err := createConnectionPair()

  if err != nil{
		t.Fatalf("proxy-server connection could not be established returned an error: %v", err)
	}


	go func(){

	}

	go func(){
		tunnel.ConnectionHandler(clientConn, serverConn)

	}	
}

func createConnectionPair() (net.Conn, net.Conn, error) {
listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, nil, err
	}

	serverConn, err := listener.Accept()
	if err != nil {
		return nil, nil, err
	}

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		return nil, nil, err
	}

	return clientConn, serverConn, nil
}

func createTunnel() chan bool {
	done := make(chan bool)
	go func(){
		tunnel.ConnectionHandler(clientConn, serverConn)
	}()
}	
