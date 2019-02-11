package ws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func runClient(addr string, msg string, count int) (total int64, cost time.Duration, err error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	t1 := time.Now()
	for i := 0; i != count; i++ {
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			err = fmt.Errorf("can't send message on ws socket: %s", err.Error())
			break
		}
		total++
	}
	t2 := time.Now()
	cost = t2.Sub(t1)
	return
}
