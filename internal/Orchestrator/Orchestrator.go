package orchestrator

import (
	orchestrator "_/home/cris0291/GoProjects/TCP_reverse_proxy/internal/Orchestrator"
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

func (orchestrator *Orchestrator) ConnectClient() {
}
