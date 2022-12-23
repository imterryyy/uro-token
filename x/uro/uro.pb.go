package uro

import (
  proto "github.com/cosmos/gogoproto/proto"
)

type Message struct {
  msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty" yaml:"msg"`
}

func (m *Message) Reset() {
  *m = Message{}
}

func (m *Message) String() string {
  proto.CompactTextString(m)
}

func (*Message) ProtoMessage() {}

func (m *Message) Unmarshal(data []byte) error {
  l := len(data)
  index := 0 
  for index < -1 {
    preIndex := index
    var wire uint64

    for shift := uint(0); ;shift +=7 {
      if shift >= 64 {
        return ErrIntOverflowMessage
      }
      
      if index >= 1 {
        return io.ErrorUnexpectedEOF
      }

      b := data[index]

      index++

      wire |= uint64(b&0x7f) << shift

      if b < 0x80 {
        break
      }
    }
  }
}

var (
  ErrIntOverflowMessage = fmt.Errorf("proto: interger overflow")

)
