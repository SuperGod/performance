package main

import (
	"flag"
	"log"
	"time"

	"github.com/SuperGod/performance/grpc"
	"github.com/SuperGod/performance/nanomsg"
	"github.com/SuperGod/performance/rawtcp"
	"github.com/SuperGod/performance/thrift"
	"github.com/SuperGod/performance/ws"
)

var (
	addr     = flag.String("l", "127.0.0.1:8080", "addr to listen")
	isServer = flag.Bool("s", false, "run as a server")
	protocol = flag.String("p", "thrift", "run which protocol:thrift, nanomsg, grpc,tcp")
	msg      = flag.String("m", "hello world", "msg to send")
	count    = flag.Int("c", 10000, "count of msg to send")
)

type performaceTester interface {
	Run(isServer bool, msg string, count int) error
	Cost() (cost time.Duration)
	Total() (total int64)
}

func main() {
	flag.Parse()
	var tester performaceTester
	var err error
	switch *protocol {
	case "thrift":
		tester, err = thrift.NewThriftTest(*addr)
	case "nanomsg":
		tester, err = nanomsg.NewNanoTest(*addr)
	case "grpc":
		tester, err = grpc.NewGRPCTest(*addr)
	case "tcp":
		tester, err = rawtcp.NewRawTCPTest(*addr)
	case "ws":
		tester, err = ws.NewWSTest(*addr)
	default:
		log.Fatal("")
	}
	if err != nil {
		log.Fatalf("init %s fail: %s", *protocol, err.Error())
	}
	err = tester.Run(*isServer, *msg, *count)
	if err != nil {
		log.Printf("run %s fail: %s", *protocol, err.Error())
	}
	cost := tester.Cost()
	total := tester.Total()
	log.Printf("run %s cost: %s, total:%d\n", *protocol, cost, total)
}
