package grpc

import (
	"context"
	"sync/atomic"

	"github.com/SuperGod/performance/grpc/msg"
)

func NewMsgServer() *MsgServer {
	ms := new(MsgServer)
	return ms
}

type MsgServer struct {
	Total int64
}

func (s *MsgServer) Send(ctx context.Context, in *msg.MsgReq) (*msg.MsgReply, error) {
	atomic.AddInt64(&s.Total, 1)
	return &msg.MsgReply{}, nil
}
