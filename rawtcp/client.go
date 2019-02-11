package rawtcp

import (
	"fmt"
	"net"
	"time"
)

func runClient(addr, msg string, count int) (total int64, cost time.Duration, err error) {
	fmt.Println("connect:", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	defer conn.Close()
	var n int
	t1 := time.Now()
	for i := 0; i != count; i++ {
		n, err = conn.Write([]byte(msg))
		if err != nil {
			break
		}
		total += int64(n)
	}
	t2 := time.Now()
	cost = t2.Sub(t1)
	return
}
