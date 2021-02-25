// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/node_execution.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A message used to fetch a single node execution entity.
type NodeExecutionGetRequest struct {
	// Uniquely identifies an individual node execution.
	Id                   *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NodeExecutionGetRequest) Reset()         { *m = NodeExecutionGetRequest{} }
func (m *NodeExecutionGetRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetRequest) ProtoMessage()    {}
func (*NodeExecutionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{0}
}

func (m *NodeExecutionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetRequest.Unmarshal(m, b)
}
func (m *NodeExecutionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetRequest.Marshal(b, m, deterministic)
}
func (m *NodeExecutionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetRequest.Merge(m, src)
}
func (m *NodeExecutionGetRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetRequest.Size(m)
}
func (m *NodeExecutionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetRequest proto.InternalMessageInfo

func (m *NodeExecutionGetRequest) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Represents a request structure to retrieve a list of node execution entities.
type NodeExecutionListRequest struct {
	// Indicates the workflow execution to filter by.
	WorkflowExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=workflow_execution_id,json=workflowExecutionId,proto3" json:"workflow_execution_id,omitempty"`
	// Indicates the number of resources to be returned.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the, server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering.
	// +optional
	SortBy *Sort `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	// Unique identifier of the parent node in the execution
	// +optional
	UniqueParentId       string   `protobuf:"bytes,6,opt,name=unique_parent_id,json=uniqueParentId,proto3" json:"unique_parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionListRequest) Reset()         { *m = NodeExecutionListRequest{} }
func (m *NodeExecutionListRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionListRequest) ProtoMessage()    {}
func (*NodeExecutionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{1}
}

func (m *NodeExecutionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionListRequest.Unmarshal(m, b)
}
func (m *NodeExecutionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionListRequest.Marshal(b, m, deterministic)
}
func (m *NodeExecutionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionListRequest.Merge(m, src)
}
func (m *NodeExecutionListRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionListRequest.Size(m)
}
func (m *NodeExecutionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionListRequest proto.InternalMessageInfo

func (m *NodeExecutionListRequest) GetWorkflowExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.WorkflowExecutionId
	}
	return nil
}

func (m *NodeExecutionListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *NodeExecutionListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NodeExecutionListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *NodeExecutionListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

func (m *NodeExecutionListRequest) GetUniqueParentId() string {
	if m != nil {
		return m.UniqueParentId
	}
	return ""
}

// Represents a request structure to retrieve a list of node execution entities launched by a specific task.
type NodeExecutionForTaskListRequest struct {
	// Indicates the node execution to filter by.
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=task_execution_id,json=taskExecutionId,proto3" json:"task_execution_id,omitempty"`
	// Indicates the number of resources to be returned.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the, server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionForTaskListRequest) Reset()         { *m = NodeExecutionForTaskListRequest{} }
func (m *NodeExecutionForTaskListRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionForTaskListRequest) ProtoMessage()    {}
func (*NodeExecutionForTaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{2}
}

func (m *NodeExecutionForTaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionForTaskListRequest.Unmarshal(m, b)
}
func (m *NodeExecutionForTaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionForTaskListRequest.Marshal(b, m, deterministic)
}
func (m *NodeExecutionForTaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionForTaskListRequest.Merge(m, src)
}
func (m *NodeExecutionForTaskListRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionForTaskListRequest.Size(m)
}
func (m *NodeExecutionForTaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionForTaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionForTaskListRequest proto.InternalMessageInfo

func (m *NodeExecutionForTaskListRequest) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecutionId
	}
	return nil
}

func (m *NodeExecutionForTaskListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *NodeExecutionForTaskListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NodeExecutionForTaskListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *NodeExecutionForTaskListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// Encapsulates all details for a single node execution entity.
// A node represents a component in the overall workflow graph. A node launch a task, multiple tasks, an entire nested
// sub-workflow, or even a separate child-workflow execution.
// The same task can be called repeatedly in a single workflow but each node is unique.
type NodeExecution struct {
	// Uniquely identifies an individual node execution.
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Path to remote data store where input blob is stored.
	InputUri string `protobuf:"bytes,2,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Computed results associated with this node execution.
	Closure *NodeExecutionClosure `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	// Metadata for Node Execution
	Metadata             *NodeExecutionMetaData `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NodeExecution) Reset()         { *m = NodeExecution{} }
func (m *NodeExecution) String() string { return proto.CompactTextString(m) }
func (*NodeExecution) ProtoMessage()    {}
func (*NodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{3}
}

func (m *NodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecution.Unmarshal(m, b)
}
func (m *NodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecution.Marshal(b, m, deterministic)
}
func (m *NodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecution.Merge(m, src)
}
func (m *NodeExecution) XXX_Size() int {
	return xxx_messageInfo_NodeExecution.Size(m)
}
func (m *NodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecution proto.InternalMessageInfo

func (m *NodeExecution) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecution) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

func (m *NodeExecution) GetClosure() *NodeExecutionClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

func (m *NodeExecution) GetMetadata() *NodeExecutionMetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Represents additional attributes related to a Node Execution
type NodeExecutionMetaData struct {
	// Node executions are grouped depending on retries of the parent
	// Retry group is unique within the context of a parent node.
	RetryGroup string `protobuf:"bytes,1,opt,name=retry_group,json=retryGroup,proto3" json:"retry_group,omitempty"`
	// Boolean flag indicating if the node has child nodes under it
	IsParentNode bool `protobuf:"varint,2,opt,name=is_parent_node,json=isParentNode,proto3" json:"is_parent_node,omitempty"`
	// Node id of the node in the original workflow
	// This maps to value of WorkflowTemplate.nodes[X].id
	SpecNodeId           string   `protobuf:"bytes,3,opt,name=spec_node_id,json=specNodeId,proto3" json:"spec_node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionMetaData) Reset()         { *m = NodeExecutionMetaData{} }
func (m *NodeExecutionMetaData) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionMetaData) ProtoMessage()    {}
func (*NodeExecutionMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{4}
}

func (m *NodeExecutionMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionMetaData.Unmarshal(m, b)
}
func (m *NodeExecutionMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionMetaData.Marshal(b, m, deterministic)
}
func (m *NodeExecutionMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionMetaData.Merge(m, src)
}
func (m *NodeExecutionMetaData) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionMetaData.Size(m)
}
func (m *NodeExecutionMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionMetaData proto.InternalMessageInfo

func (m *NodeExecutionMetaData) GetRetryGroup() string {
	if m != nil {
		return m.RetryGroup
	}
	return ""
}

func (m *NodeExecutionMetaData) GetIsParentNode() bool {
	if m != nil {
		return m.IsParentNode
	}
	return false
}

func (m *NodeExecutionMetaData) GetSpecNodeId() string {
	if m != nil {
		return m.SpecNodeId
	}
	return ""
}

// Request structure to retrieve a list of node execution entities.
type NodeExecutionList struct {
	NodeExecutions []*NodeExecution `protobuf:"bytes,1,rep,name=node_executions,json=nodeExecutions,proto3" json:"node_executions,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionList) Reset()         { *m = NodeExecutionList{} }
func (m *NodeExecutionList) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionList) ProtoMessage()    {}
func (*NodeExecutionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{5}
}

func (m *NodeExecutionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionList.Unmarshal(m, b)
}
func (m *NodeExecutionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionList.Marshal(b, m, deterministic)
}
func (m *NodeExecutionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionList.Merge(m, src)
}
func (m *NodeExecutionList) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionList.Size(m)
}
func (m *NodeExecutionList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionList proto.InternalMessageInfo

func (m *NodeExecutionList) GetNodeExecutions() []*NodeExecution {
	if m != nil {
		return m.NodeExecutions
	}
	return nil
}

func (m *NodeExecutionList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Container for node execution details and results.
type NodeExecutionClosure struct {
	// Only a node in a terminal state will have a non-empty output_result.
	//
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionClosure_OutputUri
	//	*NodeExecutionClosure_Error
	OutputResult isNodeExecutionClosure_OutputResult `protobuf_oneof:"output_result"`
	// The last recorded phase for this node execution.
	Phase core.NodeExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecution_Phase" json:"phase,omitempty"`
	// Time at which the node execution began running.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// The amount of time the node execution spent running.
	Duration *duration.Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time at which the node execution was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the node execution was last updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Store metadata for what the node launched.
	// for ex: if this is a workflow node, we store information for the launched workflow.
	//
	// Types that are valid to be assigned to TargetMetadata:
	//	*NodeExecutionClosure_WorkflowNodeMetadata
	//	*NodeExecutionClosure_TaskNodeMetadata
	TargetMetadata       isNodeExecutionClosure_TargetMetadata `protobuf_oneof:"target_metadata"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *NodeExecutionClosure) Reset()         { *m = NodeExecutionClosure{} }
func (m *NodeExecutionClosure) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionClosure) ProtoMessage()    {}
func (*NodeExecutionClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{6}
}

func (m *NodeExecutionClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionClosure.Unmarshal(m, b)
}
func (m *NodeExecutionClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionClosure.Marshal(b, m, deterministic)
}
func (m *NodeExecutionClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionClosure.Merge(m, src)
}
func (m *NodeExecutionClosure) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionClosure.Size(m)
}
func (m *NodeExecutionClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionClosure.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionClosure proto.InternalMessageInfo

type isNodeExecutionClosure_OutputResult interface {
	isNodeExecutionClosure_OutputResult()
}

type NodeExecutionClosure_OutputUri struct {
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionClosure_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionClosure_OutputUri) isNodeExecutionClosure_OutputResult() {}

func (*NodeExecutionClosure_Error) isNodeExecutionClosure_OutputResult() {}

func (m *NodeExecutionClosure) GetOutputResult() isNodeExecutionClosure_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionClosure) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionClosure_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionClosure) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionClosure_Error); ok {
		return x.Error
	}
	return nil
}

func (m *NodeExecutionClosure) GetPhase() core.NodeExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecution_UNDEFINED
}

func (m *NodeExecutionClosure) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *NodeExecutionClosure) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *NodeExecutionClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NodeExecutionClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type isNodeExecutionClosure_TargetMetadata interface {
	isNodeExecutionClosure_TargetMetadata()
}

type NodeExecutionClosure_WorkflowNodeMetadata struct {
	WorkflowNodeMetadata *WorkflowNodeMetadata `protobuf:"bytes,8,opt,name=workflow_node_metadata,json=workflowNodeMetadata,proto3,oneof"`
}

type NodeExecutionClosure_TaskNodeMetadata struct {
	TaskNodeMetadata *TaskNodeMetadata `protobuf:"bytes,9,opt,name=task_node_metadata,json=taskNodeMetadata,proto3,oneof"`
}

func (*NodeExecutionClosure_WorkflowNodeMetadata) isNodeExecutionClosure_TargetMetadata() {}

func (*NodeExecutionClosure_TaskNodeMetadata) isNodeExecutionClosure_TargetMetadata() {}

func (m *NodeExecutionClosure) GetTargetMetadata() isNodeExecutionClosure_TargetMetadata {
	if m != nil {
		return m.TargetMetadata
	}
	return nil
}

func (m *NodeExecutionClosure) GetWorkflowNodeMetadata() *WorkflowNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionClosure_WorkflowNodeMetadata); ok {
		return x.WorkflowNodeMetadata
	}
	return nil
}

func (m *NodeExecutionClosure) GetTaskNodeMetadata() *TaskNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionClosure_TaskNodeMetadata); ok {
		return x.TaskNodeMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeExecutionClosure) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeExecutionClosure_OutputUri)(nil),
		(*NodeExecutionClosure_Error)(nil),
		(*NodeExecutionClosure_WorkflowNodeMetadata)(nil),
		(*NodeExecutionClosure_TaskNodeMetadata)(nil),
	}
}

// Metadata for a WorkflowNode
type WorkflowNodeMetadata struct {
	ExecutionId          *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=executionId,proto3" json:"executionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowNodeMetadata) Reset()         { *m = WorkflowNodeMetadata{} }
func (m *WorkflowNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowNodeMetadata) ProtoMessage()    {}
func (*WorkflowNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{7}
}

func (m *WorkflowNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNodeMetadata.Unmarshal(m, b)
}
func (m *WorkflowNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNodeMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNodeMetadata.Merge(m, src)
}
func (m *WorkflowNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowNodeMetadata.Size(m)
}
func (m *WorkflowNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNodeMetadata proto.InternalMessageInfo

func (m *WorkflowNodeMetadata) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

// Metadata for the case in which the node is a TaskNode
type TaskNodeMetadata struct {
	// Captures the status of caching for this execution.
	CacheStatus core.CatalogCacheStatus `protobuf:"varint,1,opt,name=cache_status,json=cacheStatus,proto3,enum=flyteidl.core.CatalogCacheStatus" json:"cache_status,omitempty"`
	// This structure carries the catalog artifact information
	CatalogKey           *core.CatalogMetadata `protobuf:"bytes,2,opt,name=catalog_key,json=catalogKey,proto3" json:"catalog_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskNodeMetadata) Reset()         { *m = TaskNodeMetadata{} }
func (m *TaskNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskNodeMetadata) ProtoMessage()    {}
func (*TaskNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{8}
}

func (m *TaskNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNodeMetadata.Unmarshal(m, b)
}
func (m *TaskNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNodeMetadata.Marshal(b, m, deterministic)
}
func (m *TaskNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNodeMetadata.Merge(m, src)
}
func (m *TaskNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskNodeMetadata.Size(m)
}
func (m *TaskNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNodeMetadata proto.InternalMessageInfo

func (m *TaskNodeMetadata) GetCacheStatus() core.CatalogCacheStatus {
	if m != nil {
		return m.CacheStatus
	}
	return core.CatalogCacheStatus_CACHE_DISABLED
}

func (m *TaskNodeMetadata) GetCatalogKey() *core.CatalogMetadata {
	if m != nil {
		return m.CatalogKey
	}
	return nil
}

// Request structure to fetch inputs and output urls for a node execution.
type NodeExecutionGetDataRequest struct {
	// The identifier of the node execution for which to fetch inputs and outputs.
	Id                   *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NodeExecutionGetDataRequest) Reset()         { *m = NodeExecutionGetDataRequest{} }
func (m *NodeExecutionGetDataRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetDataRequest) ProtoMessage()    {}
func (*NodeExecutionGetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{9}
}

func (m *NodeExecutionGetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Unmarshal(m, b)
}
func (m *NodeExecutionGetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Marshal(b, m, deterministic)
}
func (m *NodeExecutionGetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetDataRequest.Merge(m, src)
}
func (m *NodeExecutionGetDataRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Size(m)
}
func (m *NodeExecutionGetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetDataRequest proto.InternalMessageInfo

func (m *NodeExecutionGetDataRequest) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Response structure for NodeExecutionGetDataRequest which contains inputs and outputs for a node execution.
type NodeExecutionGetDataResponse struct {
	// Signed url to fetch a core.LiteralMap of node execution inputs.
	Inputs *UrlBlob `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Signed url to fetch a core.LiteralMap of node execution outputs.
	Outputs *UrlBlob `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// Optional, full_inputs will only be populated if they are under a configured size threshold.
	FullInputs *core.LiteralMap `protobuf:"bytes,3,opt,name=full_inputs,json=fullInputs,proto3" json:"full_inputs,omitempty"`
	// Optional, full_outputs will only be populated if they are under a configured size threshold.
	FullOutputs          *core.LiteralMap `protobuf:"bytes,4,opt,name=full_outputs,json=fullOutputs,proto3" json:"full_outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeExecutionGetDataResponse) Reset()         { *m = NodeExecutionGetDataResponse{} }
func (m *NodeExecutionGetDataResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetDataResponse) ProtoMessage()    {}
func (*NodeExecutionGetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73b3eae493fd736, []int{10}
}

func (m *NodeExecutionGetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Unmarshal(m, b)
}
func (m *NodeExecutionGetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Marshal(b, m, deterministic)
}
func (m *NodeExecutionGetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetDataResponse.Merge(m, src)
}
func (m *NodeExecutionGetDataResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Size(m)
}
func (m *NodeExecutionGetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetDataResponse proto.InternalMessageInfo

func (m *NodeExecutionGetDataResponse) GetInputs() *UrlBlob {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *NodeExecutionGetDataResponse) GetOutputs() *UrlBlob {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *NodeExecutionGetDataResponse) GetFullInputs() *core.LiteralMap {
	if m != nil {
		return m.FullInputs
	}
	return nil
}

func (m *NodeExecutionGetDataResponse) GetFullOutputs() *core.LiteralMap {
	if m != nil {
		return m.FullOutputs
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeExecutionGetRequest)(nil), "flyteidl.admin.NodeExecutionGetRequest")
	proto.RegisterType((*NodeExecutionListRequest)(nil), "flyteidl.admin.NodeExecutionListRequest")
	proto.RegisterType((*NodeExecutionForTaskListRequest)(nil), "flyteidl.admin.NodeExecutionForTaskListRequest")
	proto.RegisterType((*NodeExecution)(nil), "flyteidl.admin.NodeExecution")
	proto.RegisterType((*NodeExecutionMetaData)(nil), "flyteidl.admin.NodeExecutionMetaData")
	proto.RegisterType((*NodeExecutionList)(nil), "flyteidl.admin.NodeExecutionList")
	proto.RegisterType((*NodeExecutionClosure)(nil), "flyteidl.admin.NodeExecutionClosure")
	proto.RegisterType((*WorkflowNodeMetadata)(nil), "flyteidl.admin.WorkflowNodeMetadata")
	proto.RegisterType((*TaskNodeMetadata)(nil), "flyteidl.admin.TaskNodeMetadata")
	proto.RegisterType((*NodeExecutionGetDataRequest)(nil), "flyteidl.admin.NodeExecutionGetDataRequest")
	proto.RegisterType((*NodeExecutionGetDataResponse)(nil), "flyteidl.admin.NodeExecutionGetDataResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/node_execution.proto", fileDescriptor_f73b3eae493fd736)
}

var fileDescriptor_f73b3eae493fd736 = []byte{
	// 978 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5d, 0x6f, 0xdc, 0x44,
	0x17, 0xae, 0xf3, 0xbd, 0x67, 0x93, 0x4d, 0x32, 0x6f, 0xfa, 0x76, 0x49, 0xda, 0x66, 0x31, 0x05,
	0x45, 0x48, 0xb5, 0xd5, 0xad, 0x8a, 0x00, 0x21, 0x50, 0x3e, 0xda, 0x26, 0x22, 0x81, 0x30, 0x6d,
	0x84, 0x84, 0x10, 0xd6, 0xac, 0x3d, 0xbb, 0x19, 0xc5, 0xeb, 0x71, 0x66, 0xc6, 0x0a, 0x7b, 0xc7,
	0xdf, 0xe8, 0xff, 0x43, 0x42, 0xe2, 0x37, 0x70, 0x81, 0x66, 0x3c, 0x76, 0xd6, 0x8e, 0x49, 0xa4,
	0x72, 0xc1, 0xe5, 0x9c, 0xf3, 0x3c, 0xcf, 0x99, 0xf3, 0xe1, 0x33, 0x86, 0x8f, 0x86, 0xf1, 0x44,
	0x51, 0x16, 0xc5, 0x3e, 0x89, 0xc6, 0x2c, 0xf1, 0x13, 0x1e, 0xd1, 0x80, 0xfe, 0x4a, 0xc3, 0x4c,
	0x31, 0x9e, 0x78, 0xa9, 0xe0, 0x8a, 0xa3, 0x4e, 0x01, 0xf2, 0x0c, 0x68, 0x73, 0xab, 0x46, 0x0a,
	0xf9, 0x78, 0x5c, 0x80, 0x37, 0x1f, 0x95, 0xce, 0x90, 0x0b, 0xea, 0xd7, 0xb4, 0xa6, 0xb8, 0xc6,
	0x1d, 0x12, 0x45, 0x62, 0x3e, 0xb2, 0xce, 0xc7, 0x55, 0x27, 0x8b, 0x68, 0xa2, 0xd8, 0x90, 0x51,
	0x61, 0xfd, 0x0f, 0xab, 0xfe, 0x98, 0x29, 0x2a, 0x48, 0x2c, 0xad, 0x77, 0x7b, 0xc4, 0xf9, 0x28,
	0xa6, 0xbe, 0x39, 0x0d, 0xb2, 0xa1, 0xaf, 0xd8, 0x98, 0x4a, 0x45, 0xc6, 0x69, 0x21, 0x5f, 0x07,
	0x44, 0x99, 0x20, 0xd7, 0x77, 0x73, 0x7f, 0x80, 0x07, 0xdf, 0xf1, 0x88, 0xbe, 0x2c, 0xae, 0xfc,
	0x9a, 0x2a, 0x4c, 0x2f, 0x33, 0x2a, 0x15, 0xfa, 0x0c, 0x66, 0x58, 0xd4, 0x75, 0x7a, 0xce, 0x4e,
	0xbb, 0xff, 0x89, 0x57, 0xd6, 0x43, 0x5f, 0xc3, 0xab, 0x70, 0x8e, 0xca, 0x3b, 0xe3, 0x19, 0x16,
	0xb9, 0xef, 0x66, 0xa0, 0x5b, 0xf1, 0x1f, 0x33, 0x59, 0x8a, 0xfe, 0x02, 0xf7, 0xaf, 0xb8, 0xb8,
	0x18, 0xc6, 0xfc, 0xea, 0xba, 0xe6, 0x41, 0x19, 0xe7, 0xd3, 0x5a, 0x9c, 0x1f, 0x2d, 0xb6, 0x29,
	0xd6, 0xff, 0xae, 0x6e, 0x3a, 0xd1, 0x06, 0xcc, 0xc7, 0x6c, 0xcc, 0x54, 0x77, 0xa6, 0xe7, 0xec,
	0xac, 0xe0, 0xfc, 0xa0, 0xad, 0x8a, 0x5f, 0xd0, 0xa4, 0x3b, 0xdb, 0x73, 0x76, 0x5a, 0x38, 0x3f,
	0xa0, 0x2e, 0x2c, 0x0e, 0x59, 0xac, 0xa8, 0x90, 0xdd, 0x39, 0x63, 0x2f, 0x8e, 0xe8, 0x29, 0x2c,
	0x4a, 0x2e, 0x54, 0x30, 0x98, 0x74, 0xe7, 0xcd, 0xbd, 0x36, 0xbc, 0xea, 0x3c, 0x78, 0x6f, 0xb8,
	0x50, 0x78, 0x41, 0x83, 0xf6, 0x26, 0x68, 0x07, 0xd6, 0xb2, 0x84, 0x5d, 0x66, 0x34, 0x48, 0x89,
	0xa0, 0x89, 0xd2, 0xf9, 0x2c, 0x18, 0xc5, 0x4e, 0x6e, 0x3f, 0x35, 0xe6, 0xa3, 0xc8, 0xfd, 0xd3,
	0x81, 0xed, 0x4a, 0x6d, 0x5e, 0x71, 0xf1, 0x96, 0xc8, 0x8b, 0xe9, 0x12, 0x61, 0x58, 0x57, 0x44,
	0x5e, 0x34, 0x95, 0xa7, 0xde, 0x06, 0x4d, 0x6d, 0x2a, 0xcd, 0xaa, 0xaa, 0x3a, 0xfe, 0x93, 0xb2,
	0xb8, 0x7f, 0x38, 0xb0, 0x52, 0x49, 0xf6, 0x7d, 0x47, 0x0a, 0x6d, 0x41, 0x8b, 0x25, 0x69, 0xa6,
	0x82, 0x4c, 0x30, 0x93, 0x42, 0x0b, 0x2f, 0x19, 0xc3, 0x99, 0x60, 0xe8, 0x6b, 0x58, 0x0c, 0x63,
	0x2e, 0x33, 0x41, 0x4d, 0x1e, 0xed, 0xfe, 0x93, 0xfa, 0xad, 0x2a, 0xd2, 0xfb, 0x39, 0x16, 0x17,
	0x24, 0xb4, 0x0b, 0x4b, 0x63, 0xaa, 0x48, 0x44, 0x14, 0x31, 0x09, 0xb7, 0xfb, 0x1f, 0xdf, 0x2a,
	0x70, 0x42, 0x15, 0x39, 0x20, 0x8a, 0xe0, 0x92, 0xe6, 0xfe, 0xe6, 0xc0, 0xfd, 0x46, 0x0c, 0xda,
	0x86, 0xb6, 0xa0, 0x4a, 0x4c, 0x82, 0x91, 0xe0, 0x59, 0x6a, 0x52, 0x6f, 0x61, 0x30, 0xa6, 0xd7,
	0xda, 0x82, 0x9e, 0x40, 0x87, 0xc9, 0x62, 0x6e, 0xf4, 0x2a, 0x32, 0xf9, 0x2d, 0xe1, 0x65, 0x26,
	0xf3, 0xa9, 0xd1, 0xba, 0xa8, 0x07, 0xcb, 0x32, 0xa5, 0xa1, 0x01, 0xe8, 0x71, 0xc8, 0x1b, 0x06,
	0xda, 0xa6, 0xfd, 0x47, 0x91, 0x7b, 0x09, 0xeb, 0x37, 0x3e, 0x3a, 0xf4, 0x0a, 0x56, 0xab, 0xdb,
	0x4d, 0x76, 0x9d, 0xde, 0xec, 0x4e, 0xbb, 0xff, 0xe8, 0xd6, 0x0c, 0x71, 0x27, 0x99, 0x3e, 0xca,
	0xeb, 0x41, 0x99, 0x99, 0x1a, 0x14, 0xf7, 0xf7, 0x39, 0xd8, 0x68, 0x2a, 0x2d, 0xda, 0x06, 0xe0,
	0x99, 0x2a, 0xfa, 0x65, 0x72, 0x3e, 0xbc, 0x87, 0x5b, 0xb9, 0x4d, 0xb7, 0xec, 0x05, 0xcc, 0x53,
	0x21, 0xb8, 0x30, 0x7a, 0x95, 0xdb, 0x98, 0x51, 0x28, 0x05, 0x5f, 0x6a, 0xd0, 0xe1, 0x3d, 0x9c,
	0xa3, 0xd1, 0xe7, 0x30, 0x9f, 0x9e, 0x13, 0x99, 0xf7, 0xb9, 0xd3, 0x77, 0x6f, 0x9b, 0x20, 0xef,
	0x54, 0x23, 0x71, 0x4e, 0x40, 0x5f, 0x00, 0x48, 0x45, 0x84, 0xa2, 0x51, 0x40, 0x94, 0xed, 0xf2,
	0xa6, 0x97, 0xef, 0x46, 0xaf, 0xd8, 0x8d, 0xde, 0xdb, 0x62, 0x79, 0xe2, 0x96, 0x45, 0xef, 0x2a,
	0xf4, 0x02, 0x96, 0x8a, 0x9d, 0x69, 0xa7, 0xfe, 0x83, 0x1b, 0xc4, 0x03, 0x0b, 0xc0, 0x25, 0x54,
	0x47, 0x0c, 0x05, 0x25, 0x36, 0xe2, 0xc2, 0xdd, 0x11, 0x2d, 0x7a, 0x57, 0x69, 0x6a, 0x96, 0x46,
	0x05, 0x75, 0xf1, 0x6e, 0xaa, 0x45, 0xef, 0x2a, 0xf4, 0x33, 0xfc, 0xbf, 0x5c, 0xaf, 0xa6, 0xf3,
	0xe5, 0x64, 0x2f, 0x35, 0x7f, 0x1a, 0xc5, 0x82, 0xd5, 0xb5, 0x3b, 0xb1, 0xd8, 0x43, 0x07, 0x6f,
	0x5c, 0x35, 0xd8, 0xd1, 0x29, 0x20, 0xb3, 0x99, 0xaa, 0xca, 0x2d, 0xa3, 0xdc, 0xab, 0x2b, 0xeb,
	0xdd, 0x54, 0x53, 0x5d, 0x53, 0x35, 0xdb, 0xde, 0x2a, 0xac, 0xd8, 0x49, 0x11, 0x54, 0x66, 0xb1,
	0xda, 0x5b, 0x87, 0x55, 0x45, 0xc4, 0x88, 0xaa, 0x52, 0xdf, 0x8d, 0x60, 0xa3, 0xe9, 0x96, 0xe8,
	0x18, 0xda, 0xf4, 0x7a, 0x5f, 0xbc, 0xc7, 0x03, 0x32, 0x4d, 0x77, 0xdf, 0x39, 0xb0, 0x56, 0xbf,
	0x32, 0x3a, 0x80, 0xe5, 0x90, 0x84, 0xe7, 0x34, 0x90, 0x8a, 0xa8, 0x4c, 0x9a, 0x18, 0x9d, 0xfe,
	0x87, 0xb5, 0x18, 0xfb, 0xf9, 0x83, 0xbe, 0xaf, 0x91, 0x6f, 0x0c, 0x10, 0xb7, 0xc3, 0xeb, 0x03,
	0xfa, 0x06, 0xda, 0xf6, 0xcd, 0x0f, 0x2e, 0xe8, 0xc4, 0xce, 0xfc, 0xe3, 0x66, 0x91, 0x22, 0x34,
	0x06, 0x4b, 0xf9, 0x96, 0x4e, 0xdc, 0x33, 0xd8, 0xaa, 0x3f, 0xd2, 0x66, 0x01, 0xfd, 0xcb, 0x87,
	0xfa, 0x2f, 0x07, 0x1e, 0x36, 0xeb, 0xca, 0x94, 0x27, 0x92, 0x22, 0x1f, 0x16, 0xcc, 0x96, 0x95,
	0x56, 0xfc, 0x41, 0xbd, 0xc7, 0x67, 0x22, 0xde, 0x8b, 0xf9, 0x00, 0x5b, 0x18, 0x7a, 0x06, 0x8b,
	0x79, 0x3b, 0xa5, 0xcd, 0xf2, 0x1f, 0x19, 0x05, 0x0e, 0x7d, 0x09, 0xed, 0x61, 0x16, 0xc7, 0x81,
	0x0d, 0x34, 0x6b, 0xbf, 0xb0, 0x6a, 0x16, 0xc7, 0xf9, 0x5f, 0xcf, 0x09, 0x49, 0x31, 0x68, 0xf4,
	0x51, 0x1e, 0xee, 0x2b, 0x58, 0x36, 0xdc, 0x22, 0xe6, 0xdc, 0x5d, 0x64, 0x13, 0xea, 0xfb, 0x1c,
	0xbd, 0xf7, 0xfc, 0xa7, 0x67, 0x23, 0xa6, 0xce, 0xb3, 0x81, 0x17, 0xf2, 0xb1, 0x1f, 0x4f, 0x86,
	0xca, 0x2f, 0xff, 0xb5, 0x46, 0x34, 0xf1, 0xd3, 0xc1, 0xd3, 0x11, 0xf7, 0xab, 0xff, 0x7d, 0x83,
	0x05, 0xf3, 0xfd, 0x3d, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x67, 0xaa, 0xa4, 0x45, 0x0a,
	0x00, 0x00,
}
