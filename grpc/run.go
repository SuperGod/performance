package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/SuperGod/performance/grpc/msg"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(isServer bool) {

}

type GRPCTest struct {
	addr  string
	total int64
	cost  time.Duration
}

func NewGRPCTest(addr string) (t *GRPCTest, err error) {
	t = new(GRPCTest)
	t.addr = addr
	return
}

func (t *GRPCTest) Run(isServer bool, msg string, count int) (err error) {
	if isServer {
		err = t.RunServer(msg, count)
	} else {
		err = t.RunClient(msg, count)
	}
	return
}

func (t *GRPCTest) RunServer(msgData string, count int) (err error) {
	lis, err := net.Listen("tcp", t.addr)
	if err != nil {
		return
	}
	s := grpc.NewServer()
	msgServer := NewMsgServer()
	msg.RegisterMsgServiceServer(s, msgServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	time.AfterFunc(time.Minute, func() {
		s.Stop()
		t.total = msgServer.Total
	})
	err1 := s.Serve(lis)
	if err1 != nil {
		return
	}
	return
}

func (t *GRPCTest) RunClient(msgData string, count int) (err error) {
	conn, err := grpc.Dial(t.addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := msg.NewMsgServiceClient(conn)

	var total int64
	ctx := context.Background()
	t1 := time.Now()
	for i := 0; i != count; i++ {
		_, err = c.Send(ctx, &msg.MsgReq{Data: []byte(msgData)})
		if err != nil {
			break
		}
		total++
	}
	t2 := time.Now()
	t.cost = t2.Sub(t1)
	t.total = total
	if total != int64(count) {
		log.Println("GRPC client send not match", total)
	}
	return
}

func (t *GRPCTest) Cost() (cost time.Duration) {
	cost = t.cost
	return
}

func (t *GRPCTest) Total() (total int64) {
	total = t.total
	return
}
