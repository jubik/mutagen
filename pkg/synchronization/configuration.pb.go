// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/configuration.proto

package synchronization

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	behavior "github.com/havoc-io/mutagen/pkg/filesystem/behavior"
	sync "github.com/havoc-io/mutagen/pkg/sync"
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

// Configuration encodes session configuration parameters. It is used for create
// commands to specify configuration options, for loading global configuration
// options, and for storing a merged configuration inside sessions.
type Configuration struct {
	// SynchronizationMode specifies the synchronization mode that should be
	// used in synchronization.
	SynchronizationMode sync.SynchronizationMode `protobuf:"varint,11,opt,name=synchronizationMode,proto3,enum=sync.SynchronizationMode" json:"synchronizationMode,omitempty"`
	// MaximumEntryCount specifies the maximum number of filesystem entries that
	// endpoints will tolerate managing. A zero value indicates no limit.
	MaximumEntryCount uint64 `protobuf:"varint,12,opt,name=maximumEntryCount,proto3" json:"maximumEntryCount,omitempty"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. A zero value indicates no limit.
	MaximumStagingFileSize uint64 `protobuf:"varint,13,opt,name=maximumStagingFileSize,proto3" json:"maximumStagingFileSize,omitempty"`
	// ProbeMode specifies the filesystem probing mode.
	ProbeMode behavior.ProbeMode `protobuf:"varint,14,opt,name=probeMode,proto3,enum=behavior.ProbeMode" json:"probeMode,omitempty"`
	// ScanMode specifies the synchronization root scanning mode.
	ScanMode ScanMode `protobuf:"varint,15,opt,name=scanMode,proto3,enum=synchronization.ScanMode" json:"scanMode,omitempty"`
	// StageMode specifies the file staging mode.
	StageMode StageMode `protobuf:"varint,16,opt,name=stageMode,proto3,enum=synchronization.StageMode" json:"stageMode,omitempty"`
	// SymlinkMode specifies the symlink mode that should be used in
	// synchronization.
	SymlinkMode sync.SymlinkMode `protobuf:"varint,1,opt,name=symlinkMode,proto3,enum=sync.SymlinkMode" json:"symlinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode WatchMode `protobuf:"varint,21,opt,name=watchMode,proto3,enum=synchronization.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32 `protobuf:"varint,22,opt,name=watchPollingInterval,proto3" json:"watchPollingInterval,omitempty"`
	// DefaultIgnores specifies the ignore patterns brought in from the global
	// configuration.
	// DEPRECATED: This field is no longer used when loading from global
	// configuration. Instead, ignores provided by global configuration are
	// simply merged into the ignore list of the main configuration. However,
	// older sessions still use this field.
	DefaultIgnores []string `protobuf:"bytes,31,rep,name=defaultIgnores,proto3" json:"defaultIgnores,omitempty"`
	// Ignores specifies the ignore patterns brought in from the create request.
	Ignores []string `protobuf:"bytes,32,rep,name=ignores,proto3" json:"ignores,omitempty"`
	// IgnoreVCSMode specifies the VCS ignore mode that should be used in
	// synchronization.
	IgnoreVCSMode sync.IgnoreVCSMode `protobuf:"varint,33,opt,name=ignoreVCSMode,proto3,enum=sync.IgnoreVCSMode" json:"ignoreVCSMode,omitempty"`
	// DefaultFileMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultFileMode uint32 `protobuf:"varint,63,opt,name=defaultFileMode,proto3" json:"defaultFileMode,omitempty"`
	// DefaultDirectoryMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultDirectoryMode uint32 `protobuf:"varint,64,opt,name=defaultDirectoryMode,proto3" json:"defaultDirectoryMode,omitempty"`
	// DefaultOwner specifies the default owner identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultOwner string `protobuf:"bytes,65,opt,name=defaultOwner,proto3" json:"defaultOwner,omitempty"`
	// DefaultGroup specifies the default group identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultGroup         string   `protobuf:"bytes,66,opt,name=defaultGroup,proto3" json:"defaultGroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5db9b5485282e14, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetSynchronizationMode() sync.SynchronizationMode {
	if m != nil {
		return m.SynchronizationMode
	}
	return sync.SynchronizationMode_SynchronizationModeDefault
}

func (m *Configuration) GetMaximumEntryCount() uint64 {
	if m != nil {
		return m.MaximumEntryCount
	}
	return 0
}

func (m *Configuration) GetMaximumStagingFileSize() uint64 {
	if m != nil {
		return m.MaximumStagingFileSize
	}
	return 0
}

func (m *Configuration) GetProbeMode() behavior.ProbeMode {
	if m != nil {
		return m.ProbeMode
	}
	return behavior.ProbeMode_ProbeModeDefault
}

func (m *Configuration) GetScanMode() ScanMode {
	if m != nil {
		return m.ScanMode
	}
	return ScanMode_ScanModeDefault
}

func (m *Configuration) GetStageMode() StageMode {
	if m != nil {
		return m.StageMode
	}
	return StageMode_StageModeDefault
}

func (m *Configuration) GetSymlinkMode() sync.SymlinkMode {
	if m != nil {
		return m.SymlinkMode
	}
	return sync.SymlinkMode_SymlinkModeDefault
}

func (m *Configuration) GetWatchMode() WatchMode {
	if m != nil {
		return m.WatchMode
	}
	return WatchMode_WatchModeDefault
}

func (m *Configuration) GetWatchPollingInterval() uint32 {
	if m != nil {
		return m.WatchPollingInterval
	}
	return 0
}

func (m *Configuration) GetDefaultIgnores() []string {
	if m != nil {
		return m.DefaultIgnores
	}
	return nil
}

func (m *Configuration) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Configuration) GetIgnoreVCSMode() sync.IgnoreVCSMode {
	if m != nil {
		return m.IgnoreVCSMode
	}
	return sync.IgnoreVCSMode_IgnoreVCSModeDefault
}

func (m *Configuration) GetDefaultFileMode() uint32 {
	if m != nil {
		return m.DefaultFileMode
	}
	return 0
}

func (m *Configuration) GetDefaultDirectoryMode() uint32 {
	if m != nil {
		return m.DefaultDirectoryMode
	}
	return 0
}

func (m *Configuration) GetDefaultOwner() string {
	if m != nil {
		return m.DefaultOwner
	}
	return ""
}

func (m *Configuration) GetDefaultGroup() string {
	if m != nil {
		return m.DefaultGroup
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "synchronization.Configuration")
}

func init() {
	proto.RegisterFile("synchronization/configuration.proto", fileDescriptor_e5db9b5485282e14)
}

var fileDescriptor_e5db9b5485282e14 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x15, 0x81, 0x06, 0xf5, 0xd6, 0x96, 0x79, 0x30, 0x42, 0x5f, 0x16, 0x06, 0x42, 0x79,
	0x80, 0x58, 0xdb, 0x04, 0x82, 0x27, 0x60, 0xe5, 0x8f, 0x2a, 0x84, 0x98, 0x1c, 0x09, 0x24, 0x5e,
	0x26, 0x37, 0x75, 0x13, 0x6b, 0x89, 0x5d, 0x39, 0x4e, 0x47, 0xf6, 0x99, 0xf8, 0x90, 0x28, 0xb7,
	0xc9, 0xea, 0xa6, 0x2d, 0x6f, 0xf5, 0x39, 0xbf, 0x53, 0xdf, 0xeb, 0xd3, 0xa2, 0x67, 0x79, 0x29,
	0xa3, 0x44, 0x2b, 0x29, 0x6e, 0x98, 0x11, 0x4a, 0x92, 0x48, 0xc9, 0xa9, 0x88, 0x0b, 0x0d, 0xa7,
	0x60, 0xa6, 0x95, 0x51, 0xb8, 0xdf, 0x82, 0x06, 0xcf, 0xa7, 0x22, 0xe5, 0x79, 0x99, 0x1b, 0x9e,
	0x91, 0x31, 0x4f, 0xd8, 0x5c, 0x28, 0x4d, 0x66, 0x5a, 0x8d, 0xf9, 0x65, 0xa6, 0x26, 0x7c, 0x11,
	0x1b, 0x1c, 0xb5, 0xbf, 0x3b, 0x8f, 0x98, 0xb4, 0x01, 0x6f, 0x0d, 0x30, 0x2c, 0xe6, 0xff, 0x25,
	0xae, 0x99, 0x89, 0x12, 0x9b, 0x18, 0x54, 0x04, 0x11, 0xb1, 0x54, 0x9a, 0x5f, 0xce, 0xa3, 0xdc,
	0xf6, 0x60, 0x6e, 0x62, 0x09, 0x8f, 0x41, 0xc8, 0xcb, 0x2c, 0x15, 0xf2, 0xca, 0x22, 0x8f, 0xff,
	0xee, 0xa0, 0xee, 0xd0, 0xde, 0x1c, 0x7f, 0x43, 0x07, 0xad, 0xbb, 0xbf, 0xab, 0x09, 0x77, 0x77,
	0x3d, 0xc7, 0xef, 0x9d, 0x3e, 0x09, 0x2a, 0x2f, 0x08, 0xd7, 0x01, 0xba, 0x29, 0x85, 0x5f, 0xa2,
	0xfd, 0x8c, 0xfd, 0x11, 0x59, 0x91, 0x7d, 0x96, 0x46, 0x97, 0x43, 0x55, 0x48, 0xe3, 0xee, 0x79,
	0x8e, 0x7f, 0x97, 0xae, 0x1b, 0xf8, 0x0d, 0x3a, 0xac, 0xc5, 0xd0, 0xb0, 0x58, 0xc8, 0xf8, 0x8b,
	0x48, 0x79, 0x28, 0x6e, 0xb8, 0xdb, 0x85, 0xc8, 0x16, 0x17, 0x9f, 0xa0, 0x0e, 0x74, 0x00, 0x83,
	0xf6, 0x60, 0xd0, 0x83, 0xa0, 0xa9, 0x27, 0xb8, 0x68, 0x2c, 0xba, 0xa4, 0xf0, 0x6b, 0x74, 0xbf,
	0x2a, 0x05, 0x12, 0x7d, 0x6b, 0x35, 0x6b, 0x81, 0x20, 0xac, 0x01, 0x7a, 0x8b, 0xe2, 0xb7, 0xa8,
	0x03, 0x55, 0x41, 0xee, 0x01, 0xe4, 0x06, 0xeb, 0xb9, 0x86, 0xa0, 0x4b, 0x18, 0x9f, 0xa1, 0xdd,
	0xfa, 0xf9, 0x21, 0xeb, 0x40, 0x76, 0xbf, 0x79, 0xce, 0x5b, 0x83, 0xda, 0x54, 0x75, 0x1d, 0xf4,
	0x0e, 0x91, 0x47, 0x5b, 0xae, 0xfb, 0xd5, 0x10, 0x74, 0x09, 0xe3, 0x53, 0xf4, 0x10, 0x0e, 0x17,
	0x2a, 0x4d, 0x85, 0x8c, 0x47, 0xd2, 0x70, 0x3d, 0x67, 0xa9, 0x7b, 0xe8, 0x39, 0x7e, 0x97, 0x6e,
	0xf4, 0xf0, 0x0b, 0xd4, 0x9b, 0xf0, 0x29, 0x2b, 0x52, 0x33, 0x82, 0x5f, 0x55, 0xee, 0x1e, 0x79,
	0x77, 0xfc, 0x0e, 0x6d, 0xa9, 0xd8, 0x45, 0xf7, 0x44, 0x0d, 0x78, 0x00, 0x34, 0x47, 0xfc, 0x0e,
	0x75, 0x17, 0x1f, 0x7f, 0x0e, 0x43, 0x98, 0xf9, 0x69, 0x5d, 0x06, 0xac, 0x39, 0xb2, 0x2d, 0xba,
	0x4a, 0x62, 0x1f, 0xf5, 0xeb, 0x6b, 0xaa, 0x5a, 0x21, 0xfc, 0x1e, 0x66, 0x6d, 0xcb, 0xd5, 0x6a,
	0xb5, 0xf4, 0x49, 0x68, 0x1e, 0x19, 0xa5, 0x4b, 0xc0, 0x3f, 0x2c, 0x56, 0xdb, 0xe4, 0xe1, 0x63,
	0xb4, 0x57, 0xeb, 0x3f, 0xae, 0x25, 0xd7, 0xee, 0x47, 0xcf, 0xf1, 0x3b, 0x74, 0x45, 0xb3, 0x98,
	0xaf, 0x5a, 0x15, 0x33, 0xf7, 0x7c, 0x85, 0x01, 0xed, 0xfc, 0xe4, 0x37, 0x89, 0x85, 0x49, 0x8a,
	0x71, 0x10, 0xa9, 0x8c, 0x24, 0x6c, 0xae, 0xa2, 0x57, 0x42, 0x91, 0xac, 0xa8, 0x7a, 0x96, 0x64,
	0x76, 0x15, 0x93, 0x56, 0x3d, 0xe3, 0x1d, 0xf8, 0xa3, 0x9d, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x3c, 0xda, 0xdc, 0x16, 0x71, 0x04, 0x00, 0x00,
}