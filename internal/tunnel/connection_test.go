package tunnel_test

import (
	"io"
	"net"
	"testing"

	"github.com/Cris0291/myproxy/internal/tunnel"
)

func TestConnectionHandler(t *testing.T) {
	// Arrange
	test := "Hello World"

	test_client, proxy_server, err := createConnectionPair()
	if err != nil {
		t.Fatalf("client-proxy connection could not be established returned an error: %v", err)
	}

	test_server, proxy_client, err := createConnectionPair()
	if err != nil {
		t.Fatalf("proxy-server connection could not be established returned an error: %v", err)
	}

	// Act
	test_client.Write([]byte(test))

	go func() {
		buf := make([]byte, 15)

		_, err := io.ReadAll(test_server)
		if err != nil {
			t.Errorf("test server could not read message returned an error: %v", err)
			return
		}

		for total := 0; total < len(buf); {
			n, err := test_server.Write(buf)
			total += n

			if err != nil {
				t.Errorf("test server could not read message returned an error: %v", err)
				return
			}
		}
	}()

	done, errCh1, errCh2 := tunnel.ConnectionHandler(proxy_client, proxy_server)

	data, err := io.ReadAll(test_client)
	if err != nil {
		t.Fatalf("client read returned an error: %v", err)
	}
	res := string(data)

	// Assert
	for {
		select {
		case err := <-errCh1:
			t.Fatalf("proxy server-client returned an error: %v", err)
		case err := <-errCh2:
			t.Fatalf("proxy client-server returned an error: %v", err)
		case <-done:
			select {
			case err := <-errCh1:
				t.Fatalf("proxy server-client after done returned an error: %v", err)
			case err := <-errCh2:
				t.Fatalf("proxy cleint-server after done returned an error: %v", err)
			default:
				if res != test {
					t.Fatalf("test failed waht client read is not the same that it sent")
				}
			}
		}
	}
}

func createConnectionPair() (net.Conn, net.Conn, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, nil, err
	}

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		return nil, nil, err
	}

	serverConn, err := listener.Accept()
	if err != nil {
		return nil, nil, err
	}

	return clientConn, serverConn, nil
}
