// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/stage_mode.proto

package synchronization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// StageMode specifies the mode for file staging.
type StageMode int32

const (
	// StageMode_StageModeDefault represents an unspecified staging mode. It
	// should be converted to one of the following values based on the desired
	// default behavior.
	StageMode_StageModeDefault StageMode = 0
	// StageMode_StageModeMutagen specifies that files should be staged in the
	// Mutagen data directory.
	StageMode_StageModeMutagen StageMode = 1
	// StageMode_StageModeNeighboring specifies that files should be staged in a
	// directory which neighbors the synchronization root.
	StageMode_StageModeNeighboring StageMode = 2
)

var StageMode_name = map[int32]string{
	0: "StageModeDefault",
	1: "StageModeMutagen",
	2: "StageModeNeighboring",
}

var StageMode_value = map[string]int32{
	"StageModeDefault":     0,
	"StageModeMutagen":     1,
	"StageModeNeighboring": 2,
}

func (x StageMode) String() string {
	return proto.EnumName(StageMode_name, int32(x))
}

func (StageMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b961ac0fccf3139, []int{0}
}

func init() {
	proto.RegisterEnum("synchronization.StageMode", StageMode_name, StageMode_value)
}

func init() { proto.RegisterFile("synchronization/stage_mode.proto", fileDescriptor_3b961ac0fccf3139) }

var fileDescriptor_3b961ac0fccf3139 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x2e, 0x49, 0x4c,
	0x4f, 0x8d, 0xcf, 0xcd, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x53,
	0xa1, 0x15, 0xc8, 0xc5, 0x19, 0x0c, 0x52, 0xe4, 0x9b, 0x9f, 0x92, 0x2a, 0x24, 0xc2, 0x25, 0x00,
	0xe7, 0xb8, 0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0x08, 0x30, 0xa0, 0x88, 0xfa, 0x96, 0x82, 0x58,
	0x79, 0x02, 0x8c, 0x42, 0x12, 0x5c, 0x22, 0x70, 0x51, 0xbf, 0xd4, 0xcc, 0xf4, 0x8c, 0xa4, 0xfc,
	0xa2, 0xcc, 0xbc, 0x74, 0x01, 0x26, 0x27, 0xe3, 0x28, 0xc3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24,
	0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x5c, 0x88, 0x0e, 0xdd, 0xcc, 0x7c, 0x18, 0x53, 0xbf, 0x20, 0x3b,
	0x5d, 0x1f, 0xcd, 0x1d, 0x49, 0x6c, 0x60, 0xf7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65,
	0xed, 0x21, 0x4c, 0xc3, 0x00, 0x00, 0x00,
}
