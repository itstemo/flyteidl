// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/schedule.proto

package admin

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

type Schedule_FixedRateUnit int32

const (
	Schedule_MINUTE Schedule_FixedRateUnit = 0
	Schedule_HOUR   Schedule_FixedRateUnit = 1
	Schedule_DAY    Schedule_FixedRateUnit = 2
)

var Schedule_FixedRateUnit_name = map[int32]string{
	0: "MINUTE",
	1: "HOUR",
	2: "DAY",
}
var Schedule_FixedRateUnit_value = map[string]int32{
	"MINUTE": 0,
	"HOUR":   1,
	"DAY":    2,
}

func (x Schedule_FixedRateUnit) String() string {
	return proto.EnumName(Schedule_FixedRateUnit_name, int32(x))
}
func (Schedule_FixedRateUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_9da51ba222c61417, []int{0, 0}
}

type Schedule struct {
	// Types that are valid to be assigned to ScheduleExpression:
	//	*Schedule_CronExpression
	//	*Schedule_FixedRate_
	ScheduleExpression   isSchedule_ScheduleExpression `protobuf_oneof:"ScheduleExpression"`
	KickoffTimeInputArg  string                        `protobuf:"bytes,3,opt,name=kickoff_time_input_arg,json=kickoffTimeInputArg" json:"kickoff_time_input_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_9da51ba222c61417, []int{0}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (dst *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(dst, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

type isSchedule_ScheduleExpression interface {
	isSchedule_ScheduleExpression()
}

type Schedule_CronExpression struct {
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,oneof"`
}
type Schedule_FixedRate_ struct {
	FixedRate *Schedule_FixedRate `protobuf:"bytes,2,opt,name=fixed_rate,json=fixedRate,oneof"`
}

func (*Schedule_CronExpression) isSchedule_ScheduleExpression() {}
func (*Schedule_FixedRate_) isSchedule_ScheduleExpression()     {}

func (m *Schedule) GetScheduleExpression() isSchedule_ScheduleExpression {
	if m != nil {
		return m.ScheduleExpression
	}
	return nil
}

func (m *Schedule) GetCronExpression() string {
	if x, ok := m.GetScheduleExpression().(*Schedule_CronExpression); ok {
		return x.CronExpression
	}
	return ""
}

func (m *Schedule) GetFixedRate() *Schedule_FixedRate {
	if x, ok := m.GetScheduleExpression().(*Schedule_FixedRate_); ok {
		return x.FixedRate
	}
	return nil
}

func (m *Schedule) GetKickoffTimeInputArg() string {
	if m != nil {
		return m.KickoffTimeInputArg
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Schedule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Schedule_OneofMarshaler, _Schedule_OneofUnmarshaler, _Schedule_OneofSizer, []interface{}{
		(*Schedule_CronExpression)(nil),
		(*Schedule_FixedRate_)(nil),
	}
}

func _Schedule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.CronExpression)
	case *Schedule_FixedRate_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedRate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Schedule.ScheduleExpression has unexpected type %T", x)
	}
	return nil
}

func _Schedule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Schedule)
	switch tag {
	case 1: // ScheduleExpression.cron_expression
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ScheduleExpression = &Schedule_CronExpression{x}
		return true, err
	case 2: // ScheduleExpression.fixed_rate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Schedule_FixedRate)
		err := b.DecodeMessage(msg)
		m.ScheduleExpression = &Schedule_FixedRate_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Schedule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CronExpression)))
		n += len(x.CronExpression)
	case *Schedule_FixedRate_:
		s := proto.Size(x.FixedRate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Schedule_FixedRate struct {
	Value                uint32                 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Unit                 Schedule_FixedRateUnit `protobuf:"varint,2,opt,name=unit,enum=admin.Schedule_FixedRateUnit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Schedule_FixedRate) Reset()         { *m = Schedule_FixedRate{} }
func (m *Schedule_FixedRate) String() string { return proto.CompactTextString(m) }
func (*Schedule_FixedRate) ProtoMessage()    {}
func (*Schedule_FixedRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_9da51ba222c61417, []int{0, 0}
}
func (m *Schedule_FixedRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule_FixedRate.Unmarshal(m, b)
}
func (m *Schedule_FixedRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule_FixedRate.Marshal(b, m, deterministic)
}
func (dst *Schedule_FixedRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule_FixedRate.Merge(dst, src)
}
func (m *Schedule_FixedRate) XXX_Size() int {
	return xxx_messageInfo_Schedule_FixedRate.Size(m)
}
func (m *Schedule_FixedRate) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule_FixedRate.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule_FixedRate proto.InternalMessageInfo

func (m *Schedule_FixedRate) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Schedule_FixedRate) GetUnit() Schedule_FixedRateUnit {
	if m != nil {
		return m.Unit
	}
	return Schedule_MINUTE
}

func init() {
	proto.RegisterType((*Schedule)(nil), "admin.Schedule")
	proto.RegisterType((*Schedule_FixedRate)(nil), "admin.Schedule.FixedRate")
	proto.RegisterEnum("admin.Schedule_FixedRateUnit", Schedule_FixedRateUnit_name, Schedule_FixedRateUnit_value)
}

func init() { proto.RegisterFile("admin/schedule.proto", fileDescriptor_schedule_9da51ba222c61417) }

var fileDescriptor_schedule_9da51ba222c61417 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4f, 0x83, 0x40,
	0x14, 0xc4, 0x81, 0xfe, 0xe5, 0x99, 0x56, 0xb2, 0x12, 0x83, 0x26, 0x26, 0x4d, 0x4f, 0xf5, 0x82,
	0xb1, 0xbd, 0x79, 0x6b, 0x63, 0x0d, 0x3d, 0xa8, 0xc9, 0x0a, 0x07, 0xbd, 0x10, 0x84, 0x47, 0xdd,
	0xb4, 0x2c, 0x64, 0x59, 0x4c, 0x3f, 0x9f, 0x9f, 0xcc, 0xb0, 0x2d, 0x35, 0x1e, 0x7a, 0x7c, 0x33,
	0x93, 0xdf, 0x9b, 0x0c, 0xd8, 0x51, 0x92, 0x31, 0x7e, 0x57, 0xc6, 0x5f, 0x98, 0x54, 0x5b, 0x74,
	0x0b, 0x91, 0xcb, 0x9c, 0x74, 0x94, 0x3a, 0xfe, 0x31, 0xa0, 0xff, 0x76, 0x70, 0xc8, 0x2d, 0x9c,
	0xc7, 0x22, 0xe7, 0x21, 0xee, 0x0a, 0x81, 0x65, 0xc9, 0x72, 0xee, 0xe8, 0x23, 0x7d, 0x62, 0x7a,
	0x1a, 0x1d, 0xd6, 0xc6, 0xf2, 0xa8, 0x93, 0x07, 0x80, 0x94, 0xed, 0x30, 0x09, 0x45, 0x24, 0xd1,
	0x31, 0x46, 0xfa, 0xe4, 0x6c, 0x7a, 0xe5, 0x2a, 0xa6, 0xdb, 0xf0, 0xdc, 0xa7, 0x3a, 0x41, 0x23,
	0x89, 0x9e, 0x46, 0xcd, 0xb4, 0x39, 0xc8, 0x0c, 0x2e, 0x37, 0x2c, 0xde, 0xe4, 0x69, 0x1a, 0x4a,
	0x96, 0x61, 0xc8, 0x78, 0x51, 0xc9, 0x30, 0x12, 0x6b, 0xa7, 0x55, 0x7f, 0xa3, 0x17, 0x07, 0xd7,
	0x67, 0x19, 0xae, 0x6a, 0x6f, 0x2e, 0xd6, 0xd7, 0x3e, 0x98, 0x47, 0x1c, 0xb1, 0xa1, 0xf3, 0x1d,
	0x6d, 0x2b, 0x54, 0xf5, 0x06, 0x74, 0x7f, 0x90, 0x7b, 0x68, 0x57, 0x9c, 0x49, 0xd5, 0x66, 0x38,
	0xbd, 0x39, 0xd9, 0x26, 0xe0, 0x4c, 0x52, 0x15, 0x1d, 0xbb, 0x30, 0xf8, 0x27, 0x13, 0x80, 0xee,
	0xf3, 0xea, 0x25, 0xf0, 0x97, 0x96, 0x46, 0xfa, 0xd0, 0xf6, 0x5e, 0x03, 0x6a, 0xe9, 0xa4, 0x07,
	0xad, 0xc7, 0xf9, 0xbb, 0x65, 0x2c, 0x6c, 0x20, 0x0d, 0xef, 0x6f, 0x8c, 0x45, 0xef, 0x63, 0xbf,
	0xe6, 0x67, 0x57, 0x6d, 0x3b, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xd8, 0xcd, 0xa0, 0x73,
	0x01, 0x00, 0x00,
}