package ws

import (
	"time"
)

type WSTest struct {
	addr  string
	total int64
	cost  time.Duration
}

func NewWSTest(addr string) (t *WSTest, err error) {
	t = new(WSTest)
	t.addr = addr
	return
}

func (t *WSTest) Run(isServer bool, msg string, count int) (err error) {
	if isServer {
		err = t.RunServer(msg, count)
	} else {
		err = t.RunClient(msg, count)
	}
	return
}

func (t *WSTest) RunServer(msg string, count int) (err error) {
	s := NewMsgServer(t.addr)
	err = s.Server()
	t.total = s.Total
	t.cost = s.Cost
	return
}

func (t *WSTest) RunClient(msg string, count int) (err error) {
	t.total, t.cost, err = runClient(t.addr, msg, count)
	return
}

func (t *WSTest) Cost() (cost time.Duration) {
	cost = t.cost
	return
}

func (t *WSTest) Total() (total int64) {
	total = t.total
	return
}
