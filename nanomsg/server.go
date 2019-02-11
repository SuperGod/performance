package nanomsg

import (
	"fmt"
	"sync/atomic"
	"time"

	"nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/pull" // register transports
	_ "nanomsg.org/go/mangos/v2/transport/all"
)

type MsgServer struct {
	Total int64
	Addr  string
	sock  mangos.Socket
	Cost  time.Duration
}

func NewMsgServer(addr string) (s *MsgServer) {
	s = new(MsgServer)
	s.Addr = addr
	return
}
func (s *MsgServer) Close() (err error) {
	err = s.sock.Close()
	return
}

func (s *MsgServer) Server() (err error) {

	var msg []byte
	if s.sock, err = pull.NewSocket(); err != nil {
		err = fmt.Errorf("can't get new pull socket: %s", err)
		return
	}

	if err = s.sock.ListenOptions(s.Addr, map[string]interface{}{mangos.OptionNoDelay: true}); err != nil {
		err = fmt.Errorf("can't listen on pull socket: %s", err.Error())
		return
	}
	var t1 time.Time
	for {
		// Could also use sock.RecvMsg to get header
		msg, err = s.sock.Recv()
		if err != nil {
			break
		}
		atomic.AddInt64(&s.Total, int64(len(msg)))
		if s.Total == 1 {
			t1 = time.Now()
		}
	}
	t2 := time.Now()
	s.Cost = t2.Sub(t1)
	fmt.Println("nanomsg recv:", string(msg))
	return
}
