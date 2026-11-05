package orchestrator

import (
	"fmt"
	"net"	
	tunnel "github.com/Cris0291/myproxy/internal/tunnel"
)

type ConnectionType int

const (
	ClientConnection ConnectionType = 0
	ServerConnection ConnectionType = 1
)

type ConnectionDirection int

const (
	ReadDirection  ConnectionDirection = 0
	WriteDirection ConnectionDirection = 1
)

type PairConn struct {
	Client net.Conn
	Server net.Conn
}

type Orchestrator struct {
	connections []PairConn
	proxyIp     string
	proxyPort   string
}

func (orchestrator *Orchestrator) HandleHalfClose(connection net.Conn, connDir ConnectionDirection) {
	// for now i am going to handle the case in which the net.Conn interface
	// is implemented by net.TcpConn
	tcpConn, ok := connection.(*net.TCPConn)
	if !ok {
		return
	}

	if connDir == ReadDirection {
		tcpConn.CloseRead()
	} else {
		tcpConn.CloseWrite()
	}
}

func (orchestrator *Orchestrator) Connet(serverAddress string) error {
	var clientConn net.Conn

	listener, err := net.Listen("tcp", orchestrator.proxyPort)
	if err != nil {
		return fmt.Errorf("client connection could not be established: %w", err)
	}

	clientConn, err = listener.Accept()
	if err != nil {
		return fmt.Errorf("connection could not be accepted: %w", err)
	}

	serverConn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return fmt.Errorf("server connection could not be established: %w", err)
	}

	go tunnel.ConnectionHandler(clientConn, serverConn)

	pair := PairConn{clientConn, serverConn}
	orchestrator.connections = append(orchestrator.connections, pair)

	return nil
}

func (orchestrator *Orchestrator) 
