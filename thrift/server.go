package thrift

import (
	"context"
	"sync/atomic"
)

type MsgServer struct {
	Total int64
}

func NewMsgServer() *MsgServer {
	s := new(MsgServer)
	return s
}

func (s *MsgServer) Send(ctx context.Context, msg []byte) (err error) {
	atomic.AddInt64(&s.Total, 1)
	return
}
