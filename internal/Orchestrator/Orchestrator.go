package orchestrator

import (
	"fmt"
	"net"
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

func (orchestrator *Orchestrator) Connet() error {
	var client net.Conn
	var server net.Conn

	listener, err := net.Listen("tcp", orchestrator.proxyPort)
	if err != nil {
		return fmt.Errorf("client connection could not be established: %w", err)
	}

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("connection could not be accepted: %w", err)
		}
	}

	net.Dial("tcp", address string)
}
