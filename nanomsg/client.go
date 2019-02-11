package nanomsg

import (
	"fmt"
	"time"

	mangos "nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/push"
)

func runClient(addr string, msg string, count int) (total int64, cost time.Duration, err error) {
	var sock mangos.Socket

	if sock, err = push.NewSocket(); err != nil {
		err = fmt.Errorf("can't get new push socket: %s", err.Error())
		return
	}

	if err = sock.DialOptions(addr, map[string]interface{}{mangos.OptionNoDelay: true}); err != nil {
		err = fmt.Errorf("can't dial on push socket: %s", err.Error())
		return
	}
	defer sock.Close()
	fmt.Printf("NODE1: SENDING \"%s\"\n", msg)
	t1 := time.Now()
	for i := 0; i != count; i++ {
		if err = sock.Send([]byte(msg)); err != nil {
			err = fmt.Errorf("can't send message on push socket: %s", err.Error())
			break
		}
		total++
	}
	t2 := time.Now()
	cost = t2.Sub(t1)
	time.Sleep(time.Second)
	return
}
