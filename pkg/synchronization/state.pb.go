// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/state.proto

package synchronization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/havoc-io/mutagen/pkg/synchronization/core"
	rsync "github.com/havoc-io/mutagen/pkg/synchronization/rsync"
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

type Status int32

const (
	Status_Disconnected           Status = 0
	Status_HaltedOnRootDeletion   Status = 1
	Status_HaltedOnRootTypeChange Status = 2
	Status_ConnectingAlpha        Status = 3
	Status_ConnectingBeta         Status = 4
	Status_Watching               Status = 5
	Status_Scanning               Status = 6
	Status_WaitingForRescan       Status = 7
	Status_Reconciling            Status = 8
	Status_StagingAlpha           Status = 9
	Status_StagingBeta            Status = 10
	Status_Transitioning          Status = 11
	Status_Saving                 Status = 12
)

var Status_name = map[int32]string{
	0:  "Disconnected",
	1:  "HaltedOnRootDeletion",
	2:  "HaltedOnRootTypeChange",
	3:  "ConnectingAlpha",
	4:  "ConnectingBeta",
	5:  "Watching",
	6:  "Scanning",
	7:  "WaitingForRescan",
	8:  "Reconciling",
	9:  "StagingAlpha",
	10: "StagingBeta",
	11: "Transitioning",
	12: "Saving",
}

var Status_value = map[string]int32{
	"Disconnected":           0,
	"HaltedOnRootDeletion":   1,
	"HaltedOnRootTypeChange": 2,
	"ConnectingAlpha":        3,
	"ConnectingBeta":         4,
	"Watching":               5,
	"Scanning":               6,
	"WaitingForRescan":       7,
	"Reconciling":            8,
	"StagingAlpha":           9,
	"StagingBeta":            10,
	"Transitioning":          11,
	"Saving":                 12,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8699c6f4e92f6557, []int{0}
}

type State struct {
	Session                         *Session              `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Status                          Status                `protobuf:"varint,2,opt,name=status,proto3,enum=synchronization.Status" json:"status,omitempty"`
	AlphaConnected                  bool                  `protobuf:"varint,3,opt,name=alphaConnected,proto3" json:"alphaConnected,omitempty"`
	BetaConnected                   bool                  `protobuf:"varint,4,opt,name=betaConnected,proto3" json:"betaConnected,omitempty"`
	LastError                       string                `protobuf:"bytes,5,opt,name=lastError,proto3" json:"lastError,omitempty"`
	SuccessfulSynchronizationCycles uint64                `protobuf:"varint,6,opt,name=successfulSynchronizationCycles,proto3" json:"successfulSynchronizationCycles,omitempty"`
	StagingStatus                   *rsync.ReceiverStatus `protobuf:"bytes,7,opt,name=stagingStatus,proto3" json:"stagingStatus,omitempty"`
	Conflicts                       []*core.Conflict      `protobuf:"bytes,8,rep,name=conflicts,proto3" json:"conflicts,omitempty"`
	AlphaProblems                   []*core.Problem       `protobuf:"bytes,9,rep,name=alphaProblems,proto3" json:"alphaProblems,omitempty"`
	BetaProblems                    []*core.Problem       `protobuf:"bytes,10,rep,name=betaProblems,proto3" json:"betaProblems,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}              `json:"-"`
	XXX_unrecognized                []byte                `json:"-"`
	XXX_sizecache                   int32                 `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_8699c6f4e92f6557, []int{0}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *State) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Disconnected
}

func (m *State) GetAlphaConnected() bool {
	if m != nil {
		return m.AlphaConnected
	}
	return false
}

func (m *State) GetBetaConnected() bool {
	if m != nil {
		return m.BetaConnected
	}
	return false
}

func (m *State) GetLastError() string {
	if m != nil {
		return m.LastError
	}
	return ""
}

func (m *State) GetSuccessfulSynchronizationCycles() uint64 {
	if m != nil {
		return m.SuccessfulSynchronizationCycles
	}
	return 0
}

func (m *State) GetStagingStatus() *rsync.ReceiverStatus {
	if m != nil {
		return m.StagingStatus
	}
	return nil
}

func (m *State) GetConflicts() []*core.Conflict {
	if m != nil {
		return m.Conflicts
	}
	return nil
}

func (m *State) GetAlphaProblems() []*core.Problem {
	if m != nil {
		return m.AlphaProblems
	}
	return nil
}

func (m *State) GetBetaProblems() []*core.Problem {
	if m != nil {
		return m.BetaProblems
	}
	return nil
}

func init() {
	proto.RegisterEnum("synchronization.Status", Status_name, Status_value)
	proto.RegisterType((*State)(nil), "synchronization.State")
}

func init() { proto.RegisterFile("synchronization/state.proto", fileDescriptor_8699c6f4e92f6557) }

var fileDescriptor_8699c6f4e92f6557 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xc9, 0xba, 0xa6, 0xed, 0xed, 0xbf, 0x60, 0x06, 0x44, 0x05, 0x44, 0x34, 0x10, 0x8a,
	0x10, 0x24, 0x5a, 0xf7, 0xc8, 0x13, 0xeb, 0x40, 0x7b, 0x03, 0x39, 0x93, 0x26, 0xf1, 0xe6, 0x7a,
	0x5e, 0x62, 0x91, 0xda, 0x95, 0xed, 0x56, 0x2a, 0xdf, 0x9b, 0x57, 0x84, 0x9c, 0xb8, 0x74, 0x4d,
	0x27, 0xf1, 0x56, 0x9f, 0xfb, 0x3b, 0xf6, 0xbd, 0xf7, 0x34, 0xf0, 0x42, 0x6f, 0x04, 0x2d, 0x94,
	0x14, 0xfc, 0x17, 0x31, 0x5c, 0x8a, 0x54, 0x1b, 0x62, 0x58, 0xb2, 0x54, 0xd2, 0x48, 0x34, 0x6e,
	0x14, 0x27, 0x6f, 0x9a, 0xb4, 0xb2, 0x42, 0xaa, 0x18, 0x65, 0x7c, 0xed, 0x5c, 0x93, 0x57, 0x07,
	0x57, 0x32, 0xad, 0xb9, 0x14, 0xae, 0x7c, 0x70, 0x07, 0x95, 0x8a, 0xa5, 0x54, 0x8a, 0xbb, 0x92,
	0x53, 0xe3, 0xa0, 0xd3, 0x07, 0xa1, 0xa5, 0x92, 0xf3, 0x92, 0x2d, 0x6a, 0xe6, 0xf4, 0x77, 0x0b,
	0xda, 0x99, 0xed, 0x16, 0x4d, 0xa1, 0xe3, 0xde, 0x08, 0xbd, 0xc8, 0x8b, 0xfb, 0xd3, 0x30, 0x69,
	0xf8, 0x93, 0xac, 0xae, 0xe3, 0x2d, 0x88, 0x52, 0xf0, 0xed, 0xa8, 0x2b, 0x1d, 0x1e, 0x45, 0x5e,
	0x3c, 0x9a, 0x3e, 0x3f, 0xb4, 0x54, 0x65, 0xec, 0x30, 0xf4, 0x0e, 0x46, 0xa4, 0x5c, 0x16, 0x64,
	0x26, 0x85, 0x60, 0xd4, 0xb0, 0xdb, 0xb0, 0x15, 0x79, 0x71, 0x17, 0x37, 0x54, 0xf4, 0x16, 0x86,
	0x73, 0x66, 0xee, 0x61, 0xc7, 0x15, 0xb6, 0x2f, 0xa2, 0x97, 0xd0, 0x2b, 0x89, 0x36, 0x5f, 0x94,
	0x92, 0x2a, 0x6c, 0x47, 0x5e, 0xdc, 0xc3, 0x3b, 0x01, 0x5d, 0xc1, 0x6b, 0xbd, 0xa2, 0x94, 0x69,
	0x7d, 0xb7, 0x2a, 0xb3, 0xfd, 0xbe, 0x66, 0x1b, 0x5a, 0x32, 0x1d, 0xfa, 0x91, 0x17, 0x1f, 0xe3,
	0xff, 0x61, 0xe8, 0x13, 0x0c, 0xb5, 0x21, 0x39, 0x17, 0x79, 0x3d, 0x4e, 0xd8, 0xa9, 0x16, 0xf4,
	0x34, 0xa9, 0x92, 0x4b, 0x70, 0x9d, 0x9c, 0x72, 0xb3, 0xee, 0xb3, 0xe8, 0x03, 0xf4, 0xb6, 0xb9,
	0xe8, 0xb0, 0x1b, 0xb5, 0xe2, 0xfe, 0x74, 0x94, 0xd8, 0x24, 0x92, 0x99, 0x93, 0xf1, 0x0e, 0x40,
	0xe7, 0x30, 0xac, 0x56, 0xf1, 0xbd, 0x4e, 0x49, 0x87, 0xbd, 0xca, 0x31, 0xac, 0x1d, 0x4e, 0xc5,
	0xfb, 0x0c, 0x3a, 0x83, 0x81, 0x5d, 0xcc, 0x3f, 0x0f, 0x3c, 0xe4, 0xd9, 0x43, 0xde, 0xff, 0xf1,
	0xc0, 0x77, 0x0d, 0x06, 0x30, 0xb8, 0xe4, 0x9a, 0x6e, 0xb7, 0x1a, 0x3c, 0x42, 0x21, 0x9c, 0x5c,
	0x91, 0xd2, 0xb0, 0xdb, 0x6f, 0x02, 0x4b, 0x69, 0x2e, 0x59, 0xc9, 0xec, 0x36, 0x02, 0x0f, 0x4d,
	0xe0, 0xd9, 0xfd, 0xca, 0xf5, 0x66, 0xc9, 0x66, 0x05, 0x11, 0x39, 0x0b, 0x8e, 0xd0, 0x13, 0x18,
	0xbb, 0x68, 0xb8, 0xc8, 0x3f, 0xdb, 0x06, 0x83, 0x16, 0x42, 0x30, 0xda, 0x89, 0x17, 0xcc, 0x90,
	0xe0, 0x18, 0x0d, 0xa0, 0x7b, 0x43, 0x0c, 0x2d, 0xb8, 0xc8, 0x83, 0xb6, 0x3d, 0x65, 0x94, 0x08,
	0x61, 0x4f, 0x3e, 0x3a, 0x81, 0xe0, 0x86, 0x70, 0x0b, 0x7f, 0x95, 0x0a, 0x33, 0x4d, 0x89, 0x08,
	0x3a, 0x68, 0x0c, 0x7d, 0xcc, 0xa8, 0x14, 0x94, 0x97, 0x16, 0xeb, 0xda, 0x9e, 0xb3, 0x7a, 0xcb,
	0xf5, 0x43, 0x3d, 0x8b, 0x38, 0xa5, 0x7a, 0x05, 0xd0, 0x63, 0x18, 0x5e, 0x2b, 0x22, 0x34, 0xb7,
	0xad, 0x5b, 0x57, 0x1f, 0x01, 0xf8, 0x19, 0x59, 0xdb, 0xdf, 0x83, 0x8b, 0xb3, 0x1f, 0x69, 0xce,
	0x4d, 0xb1, 0x9a, 0x27, 0x54, 0x2e, 0xd2, 0x82, 0xac, 0x25, 0xfd, 0xc8, 0x65, 0xba, 0x58, 0x19,
	0x92, 0x33, 0x91, 0x2e, 0x7f, 0xe6, 0x69, 0xe3, 0xbf, 0x3c, 0xf7, 0xab, 0x4f, 0xe6, 0xfc, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xb5, 0xfc, 0xcd, 0xef, 0x03, 0x00, 0x00,
}
