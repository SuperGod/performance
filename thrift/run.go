package thrift

import (
	"context"
	"fmt"
	"time"

	"github.com/SuperGod/performance/thrift/msg"

	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftTest struct {
	addr  string
	total int64
	cost  time.Duration
}

func NewThriftTest(addr string) (t *ThriftTest, err error) {
	t = new(ThriftTest)
	t.addr = addr
	return
}

func (t *ThriftTest) Run(isServer bool, msg string, count int) (err error) {
	if isServer {
		err = t.RunServer(msg, count)
	} else {
		err = t.RunClient(msg, count)
	}
	return
}

func (t *ThriftTest) RunServer(msgData string, count int) (err error) {

	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(t.addr)
	if err != nil {
		return err
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	fmt.Printf("%T\n", transport)
	handler := NewMsgServer()

	processor := msg.NewMsgServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", t.addr)
	time.AfterFunc(time.Minute, func() {
		server.Stop()
		t.total = handler.Total
	})
	err = server.Serve()
	return
}

func (t *ThriftTest) RunClient(msgData string, count int) (err error) {
	var transport thrift.TTransport
	transport, err = thrift.NewTSocket(t.addr)
	if err != nil {
		return err
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	ctx := context.Background()
	clt := msg.NewMsgServiceClient(thrift.NewTStandardClient(iprot, oprot))
	t1 := time.Now()
	for i := 0; i != count; i++ {
		err = clt.Send(ctx, []byte(msgData))
		if err != nil {
			break
		}
		t.total++
	}
	t2 := time.Now()
	t.cost = t2.Sub(t1)
	return
}

func (t *ThriftTest) Cost() (cost time.Duration) {
	cost = t.cost
	return
}

func (t *ThriftTest) Total() (total int64) {
	total = t.total
	return
}
