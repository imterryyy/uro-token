package uro

import (
  io "io"
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

func (*Message) Descriptor() ([]byte, []int) {
}

func (m *Message) GetMessage() string {
  if m != nil {
    return m.Message
  }
  
  return ""
}

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
        return io.ErrUnexpectedEOF
      }

      b := data[index]

      index++

      wire |= uint64(b&0x7f) << shift

      if b < 0x80 {
        break
      }
    }

    fieldNum := int32(wire >> 3)
    wireType := int(wire & 0x7)

    if wireType == 4 {
      return fmt.Errorf("proto: Uro: wiretype end group for non-group")
    }
 
    
    intStringLen := int(stringLen)
    if fieldNum <= 0 {
      return fmt.Errorf("proto: Uro: illegal tag %d (wire type %d)", fieldNum, wire)
    }

    switch fieldNum {
    case 1: 
      if wireType != 2 {
        return fmt.Errorf("proto: Uro: wrong wireType = %d for field Msg", wireType)
      }

      var stringLen uint64 
      for shift := uint(0); ;shift += 7 {
        if shift >= 64 { 
          return ErrIntOverflowMessage
        }
        
        if index >= l {
          return io.ErrUnexpectedEOF
        }

        b := data[index]

        index++

        stringLen |= uint64(b&0x7f) << shift

        if b < 0x80 {
          break
        }
      } 
      
      intStringLen := int(stringLen)
      if intStringLen < 0 {
        return ErrInvalidLengthMessage
      }

      postIndex := index + intStringLen

      if postIndex < 0 {
        return ErrInvalidLengthMessage
      }
      if postIndex > l {
        return io.ErrUnexpectedEOF
      }

      m.msg = string(data[index:postIndex])
      index = postIndex
    default:
    	index = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthorityMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
    }
  }
}

func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	index := 0
	depth := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if index >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				index++
				if dAtA[index-1] < 0x80 {
					break
				}
			}
		case 1:
			index += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if index >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[index]
				index++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrIntOverflowMessage
			}
			index += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			index += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if index < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return index, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}


var (
  ErrIntOverflowMessage = fmt.Errorf("proto: interger overflow")
  ErrInvalidLengthMessage = fmt.Errorf("proto: invalid length message")
  ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group message")
)
