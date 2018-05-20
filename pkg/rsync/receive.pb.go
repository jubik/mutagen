// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/rsync/receive.proto

package rsync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ReceivingStatus encodes that status of an rsync receiver.
type ReceiverStatus struct {
	// Path is the path currently being received.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// Received is the number of paths that have already been received.
	Received uint64 `protobuf:"varint,2,opt,name=received" json:"received,omitempty"`
	// Total is the total number of paths expected.
	Total                uint64   `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiverStatus) Reset()         { *m = ReceiverStatus{} }
func (m *ReceiverStatus) String() string { return proto.CompactTextString(m) }
func (*ReceiverStatus) ProtoMessage()    {}
func (*ReceiverStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_receive_95b50423a8e95979, []int{0}
}
func (m *ReceiverStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverStatus.Unmarshal(m, b)
}
func (m *ReceiverStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverStatus.Marshal(b, m, deterministic)
}
func (dst *ReceiverStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverStatus.Merge(dst, src)
}
func (m *ReceiverStatus) XXX_Size() int {
	return xxx_messageInfo_ReceiverStatus.Size(m)
}
func (m *ReceiverStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverStatus proto.InternalMessageInfo

func (m *ReceiverStatus) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ReceiverStatus) GetReceived() uint64 {
	if m != nil {
		return m.Received
	}
	return 0
}

func (m *ReceiverStatus) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*ReceiverStatus)(nil), "rsync.ReceiverStatus")
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/rsync/receive.proto", fileDescriptor_receive_95b50423a8e95979)
}

var fileDescriptor_receive_95b50423a8e95979 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7,
	0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0x2a, 0xae, 0xcc,
	0x4b, 0xd6, 0x2f, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x0b, 0x2a, 0x85, 0x71, 0xf1, 0x05, 0x41, 0xc4, 0x8b, 0x82, 0x4b, 0x12, 0x4b, 0x4a,
	0x8b, 0x85, 0x84, 0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x29, 0x2e, 0x0e, 0xa8, 0xee, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20,
	0x38, 0x5f, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0xbf, 0x24, 0x31, 0x47, 0x82, 0x19, 0x2c, 0x01, 0xe1,
	0x24, 0xb1, 0x81, 0x6d, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xe1, 0x0d, 0x6d, 0x9c,
	0x00, 0x00, 0x00,
}