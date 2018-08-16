// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/errors.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

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

// Defines a generic error type that dictates the behavior of the retry strategy.
type ContainerError_ErrorType int32

const (
	ContainerError_NON_RECOVERABLE ContainerError_ErrorType = 0
	ContainerError_RECOVERABLE     ContainerError_ErrorType = 1
)

var ContainerError_ErrorType_name = map[int32]string{
	0: "NON_RECOVERABLE",
	1: "RECOVERABLE",
}
var ContainerError_ErrorType_value = map[string]int32{
	"NON_RECOVERABLE": 0,
	"RECOVERABLE":     1,
}

func (x ContainerError_ErrorType) String() string {
	return proto.EnumName(ContainerError_ErrorType_name, int32(x))
}
func (ContainerError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_errors_1309c4bab349edca, []int{0, 0}
}

// Error message to propagate detailed errors from container executions to the execution
// engine.
type ContainerError struct {
	// A simplified code for errors, so that we can provide a glossary of all possible errors.
	ErrorCode string `protobuf:"bytes,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// A detailed error message.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// An abstract error type for this error. Defaults to Non_Recoverable if not
	// specified.
	ErrorType            ContainerError_ErrorType `protobuf:"varint,3,opt,name=error_type,json=errorType,proto3,enum=flyteidl.core.ContainerError_ErrorType" json:"error_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ContainerError) Reset()         { *m = ContainerError{} }
func (m *ContainerError) String() string { return proto.CompactTextString(m) }
func (*ContainerError) ProtoMessage()    {}
func (*ContainerError) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_1309c4bab349edca, []int{0}
}
func (m *ContainerError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerError.Unmarshal(m, b)
}
func (m *ContainerError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerError.Marshal(b, m, deterministic)
}
func (dst *ContainerError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerError.Merge(dst, src)
}
func (m *ContainerError) XXX_Size() int {
	return xxx_messageInfo_ContainerError.Size(m)
}
func (m *ContainerError) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerError.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerError proto.InternalMessageInfo

func (m *ContainerError) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *ContainerError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *ContainerError) GetErrorType() ContainerError_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return ContainerError_NON_RECOVERABLE
}

// Defines the errors.pb file format the container can produce to communicate
// failure reasons to the execution engine.
type ErrorDocument struct {
	// The error raised during execution.
	Error                *ContainerError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ErrorDocument) Reset()         { *m = ErrorDocument{} }
func (m *ErrorDocument) String() string { return proto.CompactTextString(m) }
func (*ErrorDocument) ProtoMessage()    {}
func (*ErrorDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_errors_1309c4bab349edca, []int{1}
}
func (m *ErrorDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDocument.Unmarshal(m, b)
}
func (m *ErrorDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDocument.Marshal(b, m, deterministic)
}
func (dst *ErrorDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDocument.Merge(dst, src)
}
func (m *ErrorDocument) XXX_Size() int {
	return xxx_messageInfo_ErrorDocument.Size(m)
}
func (m *ErrorDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDocument.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDocument proto.InternalMessageInfo

func (m *ErrorDocument) GetError() *ContainerError {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerError)(nil), "flyteidl.core.ContainerError")
	proto.RegisterType((*ErrorDocument)(nil), "flyteidl.core.ErrorDocument")
	proto.RegisterEnum("flyteidl.core.ContainerError_ErrorType", ContainerError_ErrorType_name, ContainerError_ErrorType_value)
}

func init() { proto.RegisterFile("flyteidl/core/errors.proto", fileDescriptor_errors_1309c4bab349edca) }

var fileDescriptor_errors_1309c4bab349edca = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0xc9, 0xe9, 0x81, 0xe4, 0x94, 0xce, 0x33,
	0x72, 0xf1, 0x39, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0xb9, 0x82, 0x14, 0x0a, 0xc9,
	0x72, 0x71, 0x81, 0x75, 0xc4, 0x27, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x71, 0x82, 0x45, 0x9c, 0xf3, 0x53, 0x52, 0x85, 0x94, 0xb9, 0x78, 0x21, 0xd2, 0xb9, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x15, 0x3c, 0x60, 0x41, 0x5f, 0x88, 0x98, 0x90, 0x1b,
	0xcc, 0x8c, 0x92, 0xca, 0x82, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x75, 0x3d, 0x14,
	0xab, 0xf5, 0x50, 0xad, 0xd5, 0x03, 0x93, 0x21, 0x95, 0x05, 0xa9, 0x50, 0xcb, 0x40, 0x4c, 0x25,
	0x43, 0x2e, 0x4e, 0xb8, 0xb8, 0x90, 0x30, 0x17, 0xbf, 0x9f, 0xbf, 0x5f, 0x7c, 0x90, 0xab, 0xb3,
	0x7f, 0x98, 0x6b, 0x90, 0xa3, 0x93, 0x8f, 0xab, 0x00, 0x83, 0x10, 0x3f, 0x17, 0x37, 0xb2, 0x00,
	0xa3, 0x92, 0x0b, 0x17, 0x2f, 0x58, 0x8b, 0x4b, 0x7e, 0x72, 0x69, 0x6e, 0x6a, 0x5e, 0x89, 0x90,
	0x31, 0x17, 0x2b, 0xd8, 0x40, 0xb0, 0x57, 0xb8, 0x8d, 0x64, 0xf1, 0x3a, 0x23, 0x08, 0xa2, 0xd6,
	0xc9, 0x28, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xa7,
	0x32, 0xad, 0x44, 0x1f, 0x1e, 0xa8, 0xe9, 0xa9, 0x79, 0xfa, 0x05, 0x49, 0xba, 0xe9, 0xf9, 0xfa,
	0x28, 0xe1, 0x9c, 0xc4, 0x06, 0x0e, 0x61, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xe6,
	0xea, 0x6f, 0x7f, 0x01, 0x00, 0x00,
}