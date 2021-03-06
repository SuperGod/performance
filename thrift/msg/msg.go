// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package msg

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type MsgService interface {
  // Parameters:
  //  - Msg
  Send(ctx context.Context, msg []byte) (err error)
}

type MsgServiceClient struct {
  c thrift.TClient
}

func NewMsgServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MsgServiceClient {
  return &MsgServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewMsgServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MsgServiceClient {
  return &MsgServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewMsgServiceClient(c thrift.TClient) *MsgServiceClient {
  return &MsgServiceClient{
    c: c,
  }
}

func (p *MsgServiceClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Msg
func (p *MsgServiceClient) Send(ctx context.Context, msg []byte) (err error) {
  var _args0 MsgServiceSendArgs
  _args0.Msg = msg
  var _result1 MsgServiceSendResult
  if err = p.Client_().Call(ctx, "Send", &_args0, &_result1); err != nil {
    return
  }
  return nil
}

type MsgServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler MsgService
}

func (p *MsgServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *MsgServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *MsgServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewMsgServiceProcessor(handler MsgService) *MsgServiceProcessor {

  self2 := &MsgServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["Send"] = &msgServiceProcessorSend{handler:handler}
return self2
}

func (p *MsgServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type msgServiceProcessorSend struct {
  handler MsgService
}

func (p *msgServiceProcessorSend) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := MsgServiceSendArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Send", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := MsgServiceSendResult{}
  var err2 error
  if err2 = p.handler.Send(ctx, args.Msg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Send: " + err2.Error())
    oprot.WriteMessageBegin("Send", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("Send", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Msg
type MsgServiceSendArgs struct {
  Msg []byte `thrift:"msg,1" db:"msg" json:"msg"`
}

func NewMsgServiceSendArgs() *MsgServiceSendArgs {
  return &MsgServiceSendArgs{}
}


func (p *MsgServiceSendArgs) GetMsg() []byte {
  return p.Msg
}
func (p *MsgServiceSendArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MsgServiceSendArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *MsgServiceSendArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Send_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MsgServiceSendArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err) }
  if err := oprot.WriteBinary(p.Msg); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err) }
  return err
}

func (p *MsgServiceSendArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MsgServiceSendArgs(%+v)", *p)
}

type MsgServiceSendResult struct {
}

func NewMsgServiceSendResult() *MsgServiceSendResult {
  return &MsgServiceSendResult{}
}

func (p *MsgServiceSendResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MsgServiceSendResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Send_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MsgServiceSendResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MsgServiceSendResult(%+v)", *p)
}


