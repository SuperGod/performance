package rawtcp

import (
	"net"
)

type MsgServer struct {
	addr     string
	Total    int64
	listener *net.TCPListener
}

func NewMsgServer(addr string) (s *MsgServer) {
	s = new(MsgServer)
	s.addr = addr
	return
}

func (s *MsgServer) Close() (err error) {
	err = s.listener.Close()
	return
}

func (s *MsgServer) Server() (err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.addr)
	if err != nil {
		return
	}
	s.listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}
	for {
		var tcpConn net.Conn
		tcpConn, err = s.listener.AcceptTCP()
		if err != nil {
			return
		}
		go s.handleConn(tcpConn)
	}
	return
}

func (s *MsgServer) handleConn(conn net.Conn) (err error) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	var n int
	for {
		n, err = conn.Read(buf)
		if err != nil {
			break
		}
		s.Total += int64(n)
	}
	return
}
