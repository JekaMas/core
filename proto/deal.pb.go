// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deal.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DealStatus int32

const (
	DealStatus_ANY_STATUS DealStatus = 0
	DealStatus_PENDING    DealStatus = 1
	DealStatus_ACCEPTED   DealStatus = 2
	DealStatus_CLOSED     DealStatus = 3
)

var DealStatus_name = map[int32]string{
	0: "ANY_STATUS",
	1: "PENDING",
	2: "ACCEPTED",
	3: "CLOSED",
}
var DealStatus_value = map[string]int32{
	"ANY_STATUS": 0,
	"PENDING":    1,
	"ACCEPTED":   2,
	"CLOSED":     3,
}

func (x DealStatus) String() string {
	return proto.EnumName(DealStatus_name, int32(x))
}
func (DealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type Deal struct {
	BuyerID    string     `protobuf:"bytes,1,opt,name=BuyerID" json:"BuyerID,omitempty"`
	SupplierID string     `protobuf:"bytes,2,opt,name=SupplierID" json:"SupplierID,omitempty"`
	Status     DealStatus `protobuf:"varint,3,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
	// todo(sshaman1101): need to change the type from a string to a bigInt
	// todo(sshaman1101):  but for now it will be out of scope.
	Price             string     `protobuf:"bytes,4,opt,name=price" json:"price,omitempty"`
	StartTime         *Timestamp `protobuf:"bytes,5,opt,name=startTime" json:"startTime,omitempty"`
	EndTime           *Timestamp `protobuf:"bytes,6,opt,name=endTime" json:"endTime,omitempty"`
	SpecificationHash string     `protobuf:"bytes,7,opt,name=SpecificationHash" json:"SpecificationHash,omitempty"`
	WorkTime          uint64     `protobuf:"varint,8,opt,name=workTime" json:"workTime,omitempty"`
	Id                string     `protobuf:"bytes,9,opt,name=id" json:"id,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Deal) GetBuyerID() string {
	if m != nil {
		return m.BuyerID
	}
	return ""
}

func (m *Deal) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Deal) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_ANY_STATUS
}

func (m *Deal) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Deal) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Deal) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Deal) GetSpecificationHash() string {
	if m != nil {
		return m.SpecificationHash
	}
	return ""
}

func (m *Deal) GetWorkTime() uint64 {
	if m != nil {
		return m.WorkTime
	}
	return 0
}

func (m *Deal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Deal)(nil), "sonm.Deal")
	proto.RegisterEnum("sonm.DealStatus", DealStatus_name, DealStatus_value)
}

func init() { proto.RegisterFile("deal.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x4d, 0xd7, 0xb5, 0xdb, 0x7f, 0xb2, 0xd5, 0x3f, 0x1e, 0xc2, 0x0e, 0x52, 0x3c, 0x55,
	0xd1, 0x1d, 0xe6, 0x27, 0x98, 0x6d, 0xd1, 0x81, 0xcc, 0xd1, 0xd4, 0x83, 0x27, 0x89, 0x6b, 0xc4,
	0xe0, 0xd6, 0x86, 0x26, 0x43, 0xfc, 0x5e, 0x7e, 0x40, 0x69, 0xea, 0x9c, 0xa0, 0xa7, 0xf0, 0xde,
	0xef, 0xe5, 0x3d, 0x42, 0x00, 0x0a, 0xc1, 0xd7, 0x13, 0x55, 0x57, 0xa6, 0x42, 0x57, 0x57, 0xe5,
	0x66, 0x3c, 0x92, 0x65, 0x73, 0x96, 0x92, 0xb7, 0xf6, 0xe9, 0xa7, 0x03, 0x6e, 0x22, 0xf8, 0x1a,
	0x29, 0xf8, 0xd7, 0xdb, 0x0f, 0x51, 0xcf, 0x13, 0x4a, 0x42, 0x12, 0xf5, 0xb3, 0x9d, 0xc4, 0x13,
	0x00, 0xb6, 0x55, 0x6a, 0x2d, 0x2d, 0x74, 0x2c, 0xfc, 0xe5, 0x60, 0x04, 0x9e, 0x36, 0xdc, 0x6c,
	0x35, 0xed, 0x84, 0x24, 0x1a, 0x4e, 0x83, 0x49, 0x33, 0x31, 0x69, 0x5a, 0x99, 0xf5, 0xb3, 0x6f,
	0x8e, 0xc7, 0xd0, 0x55, 0xb5, 0x5c, 0x09, 0xea, 0xda, 0x92, 0x56, 0xe0, 0x25, 0xf4, 0xb5, 0xe1,
	0xb5, 0xc9, 0xe5, 0x46, 0xd0, 0x6e, 0x48, 0xa2, 0xc1, 0x74, 0xd4, 0x56, 0x34, 0x8e, 0x36, 0x7c,
	0xa3, 0xb2, 0x7d, 0x02, 0xcf, 0xc0, 0x17, 0x65, 0x61, 0xc3, 0xde, 0xff, 0xe1, 0x1d, 0xc7, 0x0b,
	0x38, 0x62, 0x4a, 0xac, 0xe4, 0x8b, 0x5c, 0x71, 0x23, 0xab, 0xf2, 0x96, 0xeb, 0x57, 0xea, 0xdb,
	0xed, 0xbf, 0x00, 0xc7, 0xd0, 0x7b, 0xaf, 0xea, 0x37, 0xdb, 0xdc, 0x0b, 0x49, 0xe4, 0x66, 0x3f,
	0x1a, 0x87, 0xe0, 0xc8, 0x82, 0xf6, 0xed, 0x55, 0x47, 0x16, 0xe7, 0x31, 0xc0, 0xfe, 0x7d, 0x38,
	0x04, 0x98, 0x2d, 0x1e, 0x9f, 0x58, 0x3e, 0xcb, 0x1f, 0x58, 0x70, 0x80, 0x03, 0xf0, 0x97, 0xe9,
	0x22, 0x99, 0x2f, 0x6e, 0x02, 0x82, 0x87, 0xd0, 0x9b, 0xc5, 0x71, 0xba, 0xcc, 0xd3, 0x24, 0x70,
	0x10, 0xc0, 0x8b, 0xef, 0xee, 0x59, 0x9a, 0x04, 0x9d, 0x67, 0xcf, 0x7e, 0xc1, 0xd5, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x45, 0xa1, 0x9e, 0x03, 0xa7, 0x01, 0x00, 0x00,
}
