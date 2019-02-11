# performance test
send msg performance test

1. websocket
2. rawtcp
3. thrift
4. nanomsg(mangos)
5. grpc

## Result
send 1000 times

msg content: "hello world"

time cost:(less is better)
rawtcp <= websocket < nanomsg < thrift < grpc
