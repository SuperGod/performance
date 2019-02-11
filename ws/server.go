package ws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type MsgServer struct {
	Total  int64
	Addr   string
	Cost   time.Duration
	server *http.Server
}

func NewMsgServer(addr string) (s *MsgServer) {
	s = new(MsgServer)
	s.Addr = addr
	return
}
func (s *MsgServer) Close() (err error) {
	err = s.server.Close()
	return
}

func (s *MsgServer) handleConn(conn *websocket.Conn) {
	var msgType int
	var buf []byte
	var t1 time.Time
	var t2 time.Time
	var err error
	var lastMsg []byte
	for {
		msgType, buf, err = conn.ReadMessage()
		if err != nil {
			break
		}
		lastMsg = buf
		s.Total++
		if s.Total == 1 {
			t1 = time.Now()
		}
		t2 = time.Now()
	}
	s.Cost = t2.Sub(t1)
	fmt.Println("ws msg:", msgType, string(lastMsg))
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	s.Close()
}
func (s *MsgServer) handleWS(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	go s.handleConn(conn)
	return
}

func (s *MsgServer) Server() (err error) {
	http.HandleFunc("/ws", s.handleWS)
	s.server = &http.Server{Addr: s.Addr, Handler: nil}
	err = s.server.ListenAndServe()
	return
}
