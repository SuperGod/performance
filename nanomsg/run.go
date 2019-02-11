package nanomsg

import (
	"strings"
	"time"
)

type NanoTest struct {
	addr  string
	total int64
	cost  time.Duration
}

func NewNanoTest(addr string) (t *NanoTest, err error) {
	t = new(NanoTest)
	t.addr = addr
	if !strings.Contains(addr, "://") {
		t.addr = "tcp://" + addr
	}
	return
}

func (t *NanoTest) Run(isServer bool, msg string, count int) (err error) {
	if isServer {
		err = t.RunServer(msg, count)
	} else {
		err = t.RunClient(msg, count)
	}
	return
}

func (t *NanoTest) RunServer(msg string, count int) (err error) {
	s := NewMsgServer(t.addr)
	time.AfterFunc(time.Minute, func() {
		s.Close()
		t.total = s.Total
	})
	err = s.Server()
	t.cost = s.Cost
	return
}

func (t *NanoTest) RunClient(msg string, count int) (err error) {
	t.total, t.cost, err = runClient(t.addr, msg, count)
	return
}

func (t *NanoTest) Cost() (cost time.Duration) {
	cost = t.cost
	return
}

func (t *NanoTest) Total() (total int64) {
	total = t.total
	return
}
