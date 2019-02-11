export PATH=$PATH:/opt/thrift/bin
thrift -o . --out . --gen go msg.thrift
