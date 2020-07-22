// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/hpo_job.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
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

type HyperparameterTuningObjective_HyperparameterTuningObjectiveType int32

const (
	HyperparameterTuningObjective_MINIMIZE HyperparameterTuningObjective_HyperparameterTuningObjectiveType = 0
	HyperparameterTuningObjective_MAXIMIZE HyperparameterTuningObjective_HyperparameterTuningObjectiveType = 1
)

var HyperparameterTuningObjective_HyperparameterTuningObjectiveType_name = map[int32]string{
	0: "MINIMIZE",
	1: "MAXIMIZE",
}

var HyperparameterTuningObjective_HyperparameterTuningObjectiveType_value = map[string]int32{
	"MINIMIZE": 0,
	"MAXIMIZE": 1,
}

func (x HyperparameterTuningObjective_HyperparameterTuningObjectiveType) String() string {
	return proto.EnumName(HyperparameterTuningObjective_HyperparameterTuningObjectiveType_name, int32(x))
}

func (HyperparameterTuningObjective_HyperparameterTuningObjectiveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{0, 0}
}

type HPOJobConfig_HyperparameterTuningStrategy int32

const (
	HPOJobConfig_BAYESIAN HPOJobConfig_HyperparameterTuningStrategy = 0
	HPOJobConfig_RANDOM   HPOJobConfig_HyperparameterTuningStrategy = 1
)

var HPOJobConfig_HyperparameterTuningStrategy_name = map[int32]string{
	0: "BAYESIAN",
	1: "RANDOM",
}

var HPOJobConfig_HyperparameterTuningStrategy_value = map[string]int32{
	"BAYESIAN": 0,
	"RANDOM":   1,
}

func (x HPOJobConfig_HyperparameterTuningStrategy) String() string {
	return proto.EnumName(HPOJobConfig_HyperparameterTuningStrategy_name, int32(x))
}

func (HPOJobConfig_HyperparameterTuningStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{2, 0}
}

type HPOJobConfig_TrainingJobEarlyStoppingType int32

const (
	HPOJobConfig_OFF  HPOJobConfig_TrainingJobEarlyStoppingType = 0
	HPOJobConfig_AUTO HPOJobConfig_TrainingJobEarlyStoppingType = 1
)

var HPOJobConfig_TrainingJobEarlyStoppingType_name = map[int32]string{
	0: "OFF",
	1: "AUTO",
}

var HPOJobConfig_TrainingJobEarlyStoppingType_value = map[string]int32{
	"OFF":  0,
	"AUTO": 1,
}

func (x HPOJobConfig_TrainingJobEarlyStoppingType) String() string {
	return proto.EnumName(HPOJobConfig_TrainingJobEarlyStoppingType_name, int32(x))
}

func (HPOJobConfig_TrainingJobEarlyStoppingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{2, 1}
}

type HyperparameterTuningObjective struct {
	ObjectiveType        HyperparameterTuningObjective_HyperparameterTuningObjectiveType `protobuf:"varint,1,opt,name=objective_type,json=objectiveType,proto3,enum=flyteidl.plugins.sagemaker.HyperparameterTuningObjective_HyperparameterTuningObjectiveType" json:"objective_type,omitempty"`
	MetricName           string                                                          `protobuf:"bytes,2,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *HyperparameterTuningObjective) Reset()         { *m = HyperparameterTuningObjective{} }
func (m *HyperparameterTuningObjective) String() string { return proto.CompactTextString(m) }
func (*HyperparameterTuningObjective) ProtoMessage()    {}
func (*HyperparameterTuningObjective) Descriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{0}
}

func (m *HyperparameterTuningObjective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HyperparameterTuningObjective.Unmarshal(m, b)
}
func (m *HyperparameterTuningObjective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HyperparameterTuningObjective.Marshal(b, m, deterministic)
}
func (m *HyperparameterTuningObjective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HyperparameterTuningObjective.Merge(m, src)
}
func (m *HyperparameterTuningObjective) XXX_Size() int {
	return xxx_messageInfo_HyperparameterTuningObjective.Size(m)
}
func (m *HyperparameterTuningObjective) XXX_DiscardUnknown() {
	xxx_messageInfo_HyperparameterTuningObjective.DiscardUnknown(m)
}

var xxx_messageInfo_HyperparameterTuningObjective proto.InternalMessageInfo

func (m *HyperparameterTuningObjective) GetObjectiveType() HyperparameterTuningObjective_HyperparameterTuningObjectiveType {
	if m != nil {
		return m.ObjectiveType
	}
	return HyperparameterTuningObjective_MINIMIZE
}

func (m *HyperparameterTuningObjective) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

type HPOJob struct {
	TrainingJob             *TrainingJob `protobuf:"bytes,1,opt,name=training_job,json=trainingJob,proto3" json:"training_job,omitempty"`
	MaxNumberOfTrainingJobs int64        `protobuf:"varint,2,opt,name=max_number_of_training_jobs,json=maxNumberOfTrainingJobs,proto3" json:"max_number_of_training_jobs,omitempty"`
	MaxParallelTrainingJobs int64        `protobuf:"varint,3,opt,name=max_parallel_training_jobs,json=maxParallelTrainingJobs,proto3" json:"max_parallel_training_jobs,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}     `json:"-"`
	XXX_unrecognized        []byte       `json:"-"`
	XXX_sizecache           int32        `json:"-"`
}

func (m *HPOJob) Reset()         { *m = HPOJob{} }
func (m *HPOJob) String() string { return proto.CompactTextString(m) }
func (*HPOJob) ProtoMessage()    {}
func (*HPOJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{1}
}

func (m *HPOJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPOJob.Unmarshal(m, b)
}
func (m *HPOJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPOJob.Marshal(b, m, deterministic)
}
func (m *HPOJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPOJob.Merge(m, src)
}
func (m *HPOJob) XXX_Size() int {
	return xxx_messageInfo_HPOJob.Size(m)
}
func (m *HPOJob) XXX_DiscardUnknown() {
	xxx_messageInfo_HPOJob.DiscardUnknown(m)
}

var xxx_messageInfo_HPOJob proto.InternalMessageInfo

func (m *HPOJob) GetTrainingJob() *TrainingJob {
	if m != nil {
		return m.TrainingJob
	}
	return nil
}

func (m *HPOJob) GetMaxNumberOfTrainingJobs() int64 {
	if m != nil {
		return m.MaxNumberOfTrainingJobs
	}
	return 0
}

func (m *HPOJob) GetMaxParallelTrainingJobs() int64 {
	if m != nil {
		return m.MaxParallelTrainingJobs
	}
	return 0
}

type HPOJobConfig struct {
	HyperparameterRanges         *ParameterRanges                          `protobuf:"bytes,1,opt,name=hyperparameter_ranges,json=hyperparameterRanges,proto3" json:"hyperparameter_ranges,omitempty"`
	TuningStrategy               HPOJobConfig_HyperparameterTuningStrategy `protobuf:"varint,2,opt,name=tuning_strategy,json=tuningStrategy,proto3,enum=flyteidl.plugins.sagemaker.HPOJobConfig_HyperparameterTuningStrategy" json:"tuning_strategy,omitempty"`
	TuningObjective              *HyperparameterTuningObjective            `protobuf:"bytes,3,opt,name=tuning_objective,json=tuningObjective,proto3" json:"tuning_objective,omitempty"`
	TrainingJobEarlyStoppingType HPOJobConfig_TrainingJobEarlyStoppingType `protobuf:"varint,4,opt,name=training_job_early_stopping_type,json=trainingJobEarlyStoppingType,proto3,enum=flyteidl.plugins.sagemaker.HPOJobConfig_TrainingJobEarlyStoppingType" json:"training_job_early_stopping_type,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                                  `json:"-"`
	XXX_unrecognized             []byte                                    `json:"-"`
	XXX_sizecache                int32                                     `json:"-"`
}

func (m *HPOJobConfig) Reset()         { *m = HPOJobConfig{} }
func (m *HPOJobConfig) String() string { return proto.CompactTextString(m) }
func (*HPOJobConfig) ProtoMessage()    {}
func (*HPOJobConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{2}
}

func (m *HPOJobConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPOJobConfig.Unmarshal(m, b)
}
func (m *HPOJobConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPOJobConfig.Marshal(b, m, deterministic)
}
func (m *HPOJobConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPOJobConfig.Merge(m, src)
}
func (m *HPOJobConfig) XXX_Size() int {
	return xxx_messageInfo_HPOJobConfig.Size(m)
}
func (m *HPOJobConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HPOJobConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HPOJobConfig proto.InternalMessageInfo

func (m *HPOJobConfig) GetHyperparameterRanges() *ParameterRanges {
	if m != nil {
		return m.HyperparameterRanges
	}
	return nil
}

func (m *HPOJobConfig) GetTuningStrategy() HPOJobConfig_HyperparameterTuningStrategy {
	if m != nil {
		return m.TuningStrategy
	}
	return HPOJobConfig_BAYESIAN
}

func (m *HPOJobConfig) GetTuningObjective() *HyperparameterTuningObjective {
	if m != nil {
		return m.TuningObjective
	}
	return nil
}

func (m *HPOJobConfig) GetTrainingJobEarlyStoppingType() HPOJobConfig_TrainingJobEarlyStoppingType {
	if m != nil {
		return m.TrainingJobEarlyStoppingType
	}
	return HPOJobConfig_OFF
}

type HPOJobCustom struct {
	HpoJobCore              *HPOJob            `protobuf:"bytes,1,opt,name=hpo_job_core,json=hpoJobCore,proto3" json:"hpo_job_core,omitempty"`
	TrainingJobTaskTemplate *core.TaskTemplate `protobuf:"bytes,2,opt,name=training_job_task_template,json=trainingJobTaskTemplate,proto3" json:"training_job_task_template,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}           `json:"-"`
	XXX_unrecognized        []byte             `json:"-"`
	XXX_sizecache           int32              `json:"-"`
}

func (m *HPOJobCustom) Reset()         { *m = HPOJobCustom{} }
func (m *HPOJobCustom) String() string { return proto.CompactTextString(m) }
func (*HPOJobCustom) ProtoMessage()    {}
func (*HPOJobCustom) Descriptor() ([]byte, []int) {
	return fileDescriptor_50903433b4c088ac, []int{3}
}

func (m *HPOJobCustom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPOJobCustom.Unmarshal(m, b)
}
func (m *HPOJobCustom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPOJobCustom.Marshal(b, m, deterministic)
}
func (m *HPOJobCustom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPOJobCustom.Merge(m, src)
}
func (m *HPOJobCustom) XXX_Size() int {
	return xxx_messageInfo_HPOJobCustom.Size(m)
}
func (m *HPOJobCustom) XXX_DiscardUnknown() {
	xxx_messageInfo_HPOJobCustom.DiscardUnknown(m)
}

var xxx_messageInfo_HPOJobCustom proto.InternalMessageInfo

func (m *HPOJobCustom) GetHpoJobCore() *HPOJob {
	if m != nil {
		return m.HpoJobCore
	}
	return nil
}

func (m *HPOJobCustom) GetTrainingJobTaskTemplate() *core.TaskTemplate {
	if m != nil {
		return m.TrainingJobTaskTemplate
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.HyperparameterTuningObjective_HyperparameterTuningObjectiveType", HyperparameterTuningObjective_HyperparameterTuningObjectiveType_name, HyperparameterTuningObjective_HyperparameterTuningObjectiveType_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.HPOJobConfig_HyperparameterTuningStrategy", HPOJobConfig_HyperparameterTuningStrategy_name, HPOJobConfig_HyperparameterTuningStrategy_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.HPOJobConfig_TrainingJobEarlyStoppingType", HPOJobConfig_TrainingJobEarlyStoppingType_name, HPOJobConfig_TrainingJobEarlyStoppingType_value)
	proto.RegisterType((*HyperparameterTuningObjective)(nil), "flyteidl.plugins.sagemaker.HyperparameterTuningObjective")
	proto.RegisterType((*HPOJob)(nil), "flyteidl.plugins.sagemaker.HPOJob")
	proto.RegisterType((*HPOJobConfig)(nil), "flyteidl.plugins.sagemaker.HPOJobConfig")
	proto.RegisterType((*HPOJobCustom)(nil), "flyteidl.plugins.sagemaker.HPOJobCustom")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/hpo_job.proto", fileDescriptor_50903433b4c088ac)
}

var fileDescriptor_50903433b4c088ac = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x4f, 0xd4, 0x4c,
	0x14, 0xa5, 0x40, 0xf8, 0xf8, 0xee, 0xae, 0xeb, 0x66, 0xa2, 0x01, 0x17, 0x8c, 0xd8, 0x17, 0x49,
	0x0c, 0x6d, 0x58, 0x62, 0xa2, 0xd1, 0xc4, 0x2c, 0xb2, 0x84, 0x25, 0x61, 0x97, 0x94, 0x9a, 0x20,
	0x3e, 0x8c, 0xd3, 0x65, 0xb6, 0x5b, 0x68, 0x3b, 0x93, 0xe9, 0xac, 0xa1, 0xaf, 0x3e, 0xf8, 0xe6,
	0x5f, 0xf1, 0x4f, 0xf8, 0x03, 0xfc, 0x4b, 0xa6, 0xd3, 0xd9, 0x6e, 0x97, 0x68, 0x89, 0x3e, 0xde,
	0x99, 0x7b, 0xce, 0xed, 0x39, 0x3d, 0x73, 0x61, 0x7b, 0x14, 0xa6, 0x92, 0x06, 0x97, 0xa1, 0xcd,
	0xc3, 0x89, 0x1f, 0xc4, 0x89, 0x9d, 0x10, 0x9f, 0x46, 0xe4, 0x9a, 0x0a, 0x7b, 0xcc, 0x19, 0xbe,
	0x62, 0x9e, 0xc5, 0x05, 0x93, 0x0c, 0xb5, 0xa6, 0x9d, 0x96, 0xee, 0xb4, 0x8a, 0xce, 0xd6, 0x6e,
	0x05, 0x0b, 0x27, 0x82, 0x44, 0x54, 0x52, 0x81, 0x05, 0x89, 0x7d, 0x9a, 0xe4, 0x74, 0xad, 0x9d,
	0x0a, 0x88, 0x14, 0x24, 0x88, 0x83, 0xd8, 0x9f, 0x4d, 0x6f, 0x3d, 0x2a, 0xda, 0x87, 0x4c, 0x50,
	0x5b, 0x92, 0xe4, 0x5a, 0x33, 0x99, 0x5f, 0x17, 0xe1, 0xf1, 0x51, 0xca, 0xa9, 0x28, 0x26, 0xb9,
	0x93, 0x0c, 0x3d, 0xf0, 0xae, 0xe8, 0x50, 0x06, 0x9f, 0x29, 0xfa, 0x62, 0x40, 0x83, 0x4d, 0x2b,
	0x2c, 0x53, 0x4e, 0xd7, 0x8d, 0x2d, 0x63, 0xbb, 0xd1, 0xfe, 0x68, 0xfd, 0x59, 0x94, 0x55, 0xc9,
	0x59, 0x7d, 0xeb, 0xa6, 0x9c, 0x3a, 0xf7, 0x58, 0xb9, 0x44, 0x4f, 0xa0, 0x16, 0x51, 0x29, 0x82,
	0x21, 0x8e, 0x49, 0x44, 0xd7, 0x17, 0xb7, 0x8c, 0xed, 0xff, 0x1d, 0xc8, 0x8f, 0xfa, 0x24, 0xa2,
	0xe6, 0x5b, 0x78, 0x7a, 0x27, 0x29, 0xaa, 0xc3, 0xea, 0x49, 0xaf, 0xdf, 0x3b, 0xe9, 0x5d, 0x74,
	0x9b, 0x0b, 0xaa, 0xea, 0x9c, 0xe7, 0x95, 0x61, 0xfe, 0x34, 0x60, 0xe5, 0xe8, 0x74, 0x70, 0xcc,
	0x3c, 0x74, 0x0c, 0xf5, 0xb2, 0x89, 0x4a, 0x6e, 0xad, 0xfd, 0xac, 0x4a, 0xae, 0xab, 0xfb, 0x8f,
	0x99, 0xe7, 0xd4, 0xe4, 0xac, 0x40, 0x6f, 0x60, 0x23, 0x22, 0x37, 0x38, 0x9e, 0x44, 0x1e, 0x15,
	0x98, 0x8d, 0x70, 0x99, 0x39, 0x51, 0x42, 0x96, 0x9c, 0xb5, 0x88, 0xdc, 0xf4, 0x55, 0xc7, 0x60,
	0x54, 0x62, 0x4a, 0xd0, 0x6b, 0x68, 0x65, 0xe8, 0x4c, 0x54, 0x18, 0xd2, 0xf0, 0x16, 0x78, 0xa9,
	0x00, 0x9f, 0xea, 0x86, 0x32, 0xd8, 0xfc, 0xb1, 0x0c, 0xf5, 0x5c, 0xd1, 0x3b, 0x16, 0x8f, 0x02,
	0x1f, 0x7d, 0x82, 0x87, 0xe3, 0x39, 0x8f, 0x74, 0xa8, 0xb4, 0xc0, 0xe7, 0x55, 0x02, 0x4f, 0xa7,
	0x18, 0x47, 0x41, 0x9c, 0x07, 0xf3, 0x4c, 0xf9, 0x29, 0x8a, 0xe1, 0xbe, 0x54, 0xbe, 0xe3, 0x44,
	0x0a, 0x22, 0xa9, 0x9f, 0x2a, 0x85, 0x8d, 0x76, 0xb7, 0x32, 0x2b, 0xa5, 0x8f, 0xfc, 0x6d, 0x34,
	0xce, 0x34, 0x99, 0xd3, 0x90, 0x73, 0x35, 0xba, 0x84, 0xa6, 0x9e, 0x57, 0xc4, 0x45, 0xb9, 0x52,
	0x6b, 0xbf, 0xfa, 0xe7, 0x70, 0x3a, 0x5a, 0xc2, 0xec, 0x05, 0x7c, 0x33, 0x60, 0xab, 0xec, 0x3c,
	0xa6, 0x44, 0x84, 0x29, 0x4e, 0x24, 0xe3, 0x3c, 0x3b, 0x52, 0x6f, 0x62, 0xf9, 0x2f, 0x75, 0x96,
	0x7e, 0x55, 0x37, 0xa3, 0x3b, 0xd3, 0x6c, 0x2a, 0xfd, 0x9b, 0xb2, 0xe2, 0xd6, 0x7c, 0x09, 0x9b,
	0x55, 0x2e, 0x65, 0xc1, 0xde, 0xef, 0x7c, 0xe8, 0x9e, 0xf5, 0x3a, 0xfd, 0xe6, 0x02, 0x02, 0x58,
	0x71, 0x3a, 0xfd, 0x83, 0xc1, 0x49, 0xd3, 0x30, 0x77, 0x61, 0xb3, 0x6a, 0x2e, 0xfa, 0x0f, 0x96,
	0x06, 0x87, 0x87, 0xcd, 0x05, 0xb4, 0x0a, 0xcb, 0x9d, 0xf7, 0xee, 0xa0, 0x69, 0x98, 0xdf, 0x8d,
	0x22, 0x45, 0x93, 0x44, 0xb2, 0x08, 0x1d, 0x40, 0x5d, 0xef, 0x36, 0x9c, 0x6d, 0x13, 0x1d, 0x1e,
	0xf3, 0x6e, 0xe1, 0x0e, 0x8c, 0x39, 0x53, 0x06, 0x08, 0x8a, 0xce, 0xa1, 0x35, 0x67, 0x69, 0xb6,
	0x93, 0xb0, 0xa4, 0x11, 0x0f, 0x89, 0xcc, 0xdf, 0x77, 0xad, 0xbd, 0x31, 0xe3, 0xcc, 0x26, 0x59,
	0x2e, 0x49, 0xae, 0x5d, 0xdd, 0xe2, 0xac, 0x95, 0x2c, 0x2a, 0x5f, 0xec, 0xbf, 0xb8, 0xd8, 0xf3,
	0x03, 0x39, 0x9e, 0x78, 0xd6, 0x90, 0x45, 0x76, 0x98, 0x8e, 0xa4, 0x5d, 0xac, 0x3f, 0x9f, 0xc6,
	0x36, 0xf7, 0x76, 0x7c, 0x66, 0xdf, 0x5e, 0xa0, 0xde, 0x8a, 0xda, 0x87, 0x7b, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x95, 0x92, 0x8c, 0xd4, 0x05, 0x00, 0x00,
}
