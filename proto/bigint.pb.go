// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bigint.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// BigInt represents multi-precision value stored as a big integer.
type BigInt struct {
	Neg bool   `protobuf:"varint,1,opt,name=neg" json:"neg,omitempty"`
	Abs []byte `protobuf:"bytes,2,opt,name=abs,proto3" json:"abs,omitempty"`
}

func (m *BigInt) Reset()                    { *m = BigInt{} }
func (m *BigInt) String() string            { return proto.CompactTextString(m) }
func (*BigInt) ProtoMessage()               {}
func (*BigInt) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BigInt) GetNeg() bool {
	if m != nil {
		return m.Neg
	}
	return false
}

func (m *BigInt) GetAbs() []byte {
	if m != nil {
		return m.Abs
	}
	return nil
}

func init() {
	proto.RegisterType((*BigInt)(nil), "sonm.BigInt")
}

func init() { proto.RegisterFile("bigint.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xca, 0x4c, 0xcf,
	0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xce, 0xcf, 0xcb, 0x55, 0xd2,
	0xe1, 0x62, 0x73, 0xca, 0x4c, 0xf7, 0xcc, 0x2b, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4b, 0x4d, 0x97,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x02, 0x31, 0x41, 0x22, 0x89, 0x49, 0xc5, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x20, 0x66, 0x12, 0x1b, 0x58, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x52, 0x50, 0x08, 0x79, 0x4a, 0x00, 0x00, 0x00,
}