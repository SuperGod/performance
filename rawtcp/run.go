package rawtcp

import "time"

type RawTCPTest struct {
	addr  string
	total int64
	cost  time.Duration
}

func NewRawTCPTest(addr string) (t *RawTCPTest, err error) {
	t = new(RawTCPTest)
	t.addr = addr
	return
}

func (t *RawTCPTest) Run(isServer bool, msg string, count int) (err error) {
	if isServer {
		err = t.RunServer(msg, count)
	} else {
		err = t.RunClient(msg, count)
	}
	return
}

func (t *RawTCPTest) RunServer(msg string, count int) (err error) {
	server := NewMsgServer(t.addr)
	time.AfterFunc(time.Minute, func() {
		server.Close()
		t.total = server.Total
	})
	err = server.Server()
	return
}

func (t *RawTCPTest) RunClient(msg string, count int) (err error) {
	t.total, t.cost, err = runClient(t.addr, msg, count)
	return
}

func (t *RawTCPTest) Cost() (cost time.Duration) {
	cost = t.cost
	return
}

func (t *RawTCPTest) Total() (total int64) {
	total = t.total
	return
}
