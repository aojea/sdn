// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/invocation.proto

package resultstore // import "google.golang.org/genproto/googleapis/devtools/resultstore/v2"

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

// An Invocation typically represents the result of running a tool. Each has a
// unique ID, typically generated by the server. Target resources under each
// Invocation contain the bulk of the data.
type Invocation struct {
	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Invocation. They must match
	// the resource name after proper encoding.
	Id *Invocation_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status of the invocation.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// When this invocation started and its duration.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Attributes of this invocation.
	InvocationAttributes *InvocationAttributes `protobuf:"bytes,5,opt,name=invocation_attributes,json=invocationAttributes,proto3" json:"invocation_attributes,omitempty"`
	// The workspace the tool was run in.
	WorkspaceInfo *WorkspaceInfo `protobuf:"bytes,6,opt,name=workspace_info,json=workspaceInfo,proto3" json:"workspace_info,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for invocation level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	// Use this field to specify build logs, and other invocation level logs.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	// Summary of aggregate coverage across all Actions in this Invocation.
	// the server populates this for you in the post-processing phase.
	CoverageSummaries    []*LanguageCoverageSummary `protobuf:"bytes,9,rep,name=coverage_summaries,json=coverageSummaries,proto3" json:"coverage_summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{0}
}
func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (dst *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(dst, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetId() *Invocation_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Invocation) GetStatusAttributes() *StatusAttributes {
	if m != nil {
		return m.StatusAttributes
	}
	return nil
}

func (m *Invocation) GetTiming() *Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *Invocation) GetInvocationAttributes() *InvocationAttributes {
	if m != nil {
		return m.InvocationAttributes
	}
	return nil
}

func (m *Invocation) GetWorkspaceInfo() *WorkspaceInfo {
	if m != nil {
		return m.WorkspaceInfo
	}
	return nil
}

func (m *Invocation) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Invocation) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Invocation) GetCoverageSummaries() []*LanguageCoverageSummary {
	if m != nil {
		return m.CoverageSummaries
	}
	return nil
}

// The resource ID components that identify the Invocation.
type Invocation_Id struct {
	// The Invocation ID.
	InvocationId         string   `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation_Id) Reset()         { *m = Invocation_Id{} }
func (m *Invocation_Id) String() string { return proto.CompactTextString(m) }
func (*Invocation_Id) ProtoMessage()    {}
func (*Invocation_Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{0, 0}
}
func (m *Invocation_Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation_Id.Unmarshal(m, b)
}
func (m *Invocation_Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation_Id.Marshal(b, m, deterministic)
}
func (dst *Invocation_Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation_Id.Merge(dst, src)
}
func (m *Invocation_Id) XXX_Size() int {
	return xxx_messageInfo_Invocation_Id.Size(m)
}
func (m *Invocation_Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation_Id proto.InternalMessageInfo

func (m *Invocation_Id) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

// If known, represents the state of the user/build-system workspace.
type WorkspaceContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkspaceContext) Reset()         { *m = WorkspaceContext{} }
func (m *WorkspaceContext) String() string { return proto.CompactTextString(m) }
func (*WorkspaceContext) ProtoMessage()    {}
func (*WorkspaceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{1}
}
func (m *WorkspaceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceContext.Unmarshal(m, b)
}
func (m *WorkspaceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceContext.Marshal(b, m, deterministic)
}
func (dst *WorkspaceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceContext.Merge(dst, src)
}
func (m *WorkspaceContext) XXX_Size() int {
	return xxx_messageInfo_WorkspaceContext.Size(m)
}
func (m *WorkspaceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceContext.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceContext proto.InternalMessageInfo

// Describes the workspace under which the tool was invoked, this includes
// information that was fed into the command, the source code referenced, and
// the tool itself.
type WorkspaceInfo struct {
	// Data about the workspace that might be useful for debugging.
	WorkspaceContext *WorkspaceContext `protobuf:"bytes,1,opt,name=workspace_context,json=workspaceContext,proto3" json:"workspace_context,omitempty"`
	// Where the tool was invoked
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The client's working directory where the build/test was run from.
	WorkingDirectory string `protobuf:"bytes,4,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory,omitempty"`
	// Tools should set tool_tag to the name of the tool or use case.
	ToolTag string `protobuf:"bytes,5,opt,name=tool_tag,json=toolTag,proto3" json:"tool_tag,omitempty"`
	// The command lines invoked. The first command line is the one typed by the
	// user, then each one after that should be an expansion of the previous
	// command line.
	CommandLines         []*CommandLine `protobuf:"bytes,7,rep,name=command_lines,json=commandLines,proto3" json:"command_lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WorkspaceInfo) Reset()         { *m = WorkspaceInfo{} }
func (m *WorkspaceInfo) String() string { return proto.CompactTextString(m) }
func (*WorkspaceInfo) ProtoMessage()    {}
func (*WorkspaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{2}
}
func (m *WorkspaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceInfo.Unmarshal(m, b)
}
func (m *WorkspaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceInfo.Marshal(b, m, deterministic)
}
func (dst *WorkspaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceInfo.Merge(dst, src)
}
func (m *WorkspaceInfo) XXX_Size() int {
	return xxx_messageInfo_WorkspaceInfo.Size(m)
}
func (m *WorkspaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceInfo proto.InternalMessageInfo

func (m *WorkspaceInfo) GetWorkspaceContext() *WorkspaceContext {
	if m != nil {
		return m.WorkspaceContext
	}
	return nil
}

func (m *WorkspaceInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *WorkspaceInfo) GetWorkingDirectory() string {
	if m != nil {
		return m.WorkingDirectory
	}
	return ""
}

func (m *WorkspaceInfo) GetToolTag() string {
	if m != nil {
		return m.ToolTag
	}
	return ""
}

func (m *WorkspaceInfo) GetCommandLines() []*CommandLine {
	if m != nil {
		return m.CommandLines
	}
	return nil
}

// The command and arguments that produced this Invocation.
type CommandLine struct {
	// A label describing this command line.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The command-line tool that is run: argv[0].
	Tool string `protobuf:"bytes,2,opt,name=tool,proto3" json:"tool,omitempty"`
	// The arguments to the above tool: argv[1]...argv[N].
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// The actual command that was run with the tool.  (e.g. "build", or "test")
	// Omit if the tool doesn't accept a command.
	// This is a duplicate of one of the fields in args.
	Command              string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandLine) Reset()         { *m = CommandLine{} }
func (m *CommandLine) String() string { return proto.CompactTextString(m) }
func (*CommandLine) ProtoMessage()    {}
func (*CommandLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{3}
}
func (m *CommandLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLine.Unmarshal(m, b)
}
func (m *CommandLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLine.Marshal(b, m, deterministic)
}
func (dst *CommandLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLine.Merge(dst, src)
}
func (m *CommandLine) XXX_Size() int {
	return xxx_messageInfo_CommandLine.Size(m)
}
func (m *CommandLine) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLine.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLine proto.InternalMessageInfo

func (m *CommandLine) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *CommandLine) GetTool() string {
	if m != nil {
		return m.Tool
	}
	return ""
}

func (m *CommandLine) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CommandLine) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

// Attributes that apply to all invocations.
type InvocationAttributes struct {
	// The project ID this invocation is associated with. This must be
	// set in the CreateInvocation call, and can't be changed.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The list of users in the command chain.  The first user in this sequence
	// is the one who instigated the first command in the chain.
	Users []string `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	// Labels to categorize this invocation.
	// This is implemented as a set. All labels will be unique. Any duplicate
	// labels added will be ignored. Labels will be returned in lexicographical
	// order. Labels should be short, easy to read, and you
	// shouldn't have more than a handful.
	// Labels should match regex \w([- \w]*\w)?
	// Labels should not be used for unique properties such as unique IDs.
	// Use properties in cases that don't meet these conditions.
	Labels []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	// This field describes the overall context or purpose of this invocation.
	// It will be used in the UI to give users more information about
	// how or why this invocation was run.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// If this Invocation was run in the context of a larger Continuous
	// Integration build or other automated system, this field may contain more
	// information about the greater context.
	InvocationContexts   []*InvocationContext `protobuf:"bytes,6,rep,name=invocation_contexts,json=invocationContexts,proto3" json:"invocation_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InvocationAttributes) Reset()         { *m = InvocationAttributes{} }
func (m *InvocationAttributes) String() string { return proto.CompactTextString(m) }
func (*InvocationAttributes) ProtoMessage()    {}
func (*InvocationAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{4}
}
func (m *InvocationAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvocationAttributes.Unmarshal(m, b)
}
func (m *InvocationAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvocationAttributes.Marshal(b, m, deterministic)
}
func (dst *InvocationAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvocationAttributes.Merge(dst, src)
}
func (m *InvocationAttributes) XXX_Size() int {
	return xxx_messageInfo_InvocationAttributes.Size(m)
}
func (m *InvocationAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_InvocationAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_InvocationAttributes proto.InternalMessageInfo

func (m *InvocationAttributes) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *InvocationAttributes) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *InvocationAttributes) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *InvocationAttributes) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvocationAttributes) GetInvocationContexts() []*InvocationContext {
	if m != nil {
		return m.InvocationContexts
	}
	return nil
}

// Describes the invocation context which includes a display name and URL.
type InvocationContext struct {
	// A human readable name for the context under which this Invocation was run.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A URL pointing to a UI containing more information
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvocationContext) Reset()         { *m = InvocationContext{} }
func (m *InvocationContext) String() string { return proto.CompactTextString(m) }
func (*InvocationContext) ProtoMessage()    {}
func (*InvocationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_invocation_3b0cc1029ea97db2, []int{5}
}
func (m *InvocationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvocationContext.Unmarshal(m, b)
}
func (m *InvocationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvocationContext.Marshal(b, m, deterministic)
}
func (dst *InvocationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvocationContext.Merge(dst, src)
}
func (m *InvocationContext) XXX_Size() int {
	return xxx_messageInfo_InvocationContext.Size(m)
}
func (m *InvocationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_InvocationContext.DiscardUnknown(m)
}

var xxx_messageInfo_InvocationContext proto.InternalMessageInfo

func (m *InvocationContext) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *InvocationContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Invocation)(nil), "google.devtools.resultstore.v2.Invocation")
	proto.RegisterType((*Invocation_Id)(nil), "google.devtools.resultstore.v2.Invocation.Id")
	proto.RegisterType((*WorkspaceContext)(nil), "google.devtools.resultstore.v2.WorkspaceContext")
	proto.RegisterType((*WorkspaceInfo)(nil), "google.devtools.resultstore.v2.WorkspaceInfo")
	proto.RegisterType((*CommandLine)(nil), "google.devtools.resultstore.v2.CommandLine")
	proto.RegisterType((*InvocationAttributes)(nil), "google.devtools.resultstore.v2.InvocationAttributes")
	proto.RegisterType((*InvocationContext)(nil), "google.devtools.resultstore.v2.InvocationContext")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/invocation.proto", fileDescriptor_invocation_3b0cc1029ea97db2)
}

var fileDescriptor_invocation_3b0cc1029ea97db2 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x71, 0x6b, 0x13, 0x31,
	0x14, 0xa7, 0xed, 0xda, 0xad, 0xaf, 0xab, 0x74, 0x71, 0xca, 0x59, 0x50, 0x6a, 0x15, 0xe9, 0x18,
	0x5e, 0xb5, 0x2a, 0x82, 0xa2, 0xa0, 0x13, 0x59, 0x61, 0xc8, 0xb8, 0x0d, 0x04, 0x41, 0x4a, 0x7a,
	0x97, 0xc6, 0xe8, 0x35, 0xa9, 0x49, 0xae, 0xb3, 0xdf, 0xc7, 0x0f, 0xe5, 0x57, 0xf0, 0x5b, 0x48,
	0x72, 0xb9, 0xf6, 0x56, 0xa7, 0xd7, 0xff, 0xf2, 0x7e, 0xf7, 0x7e, 0xbf, 0xbc, 0xbc, 0xfc, 0xf2,
	0x0e, 0xfa, 0x54, 0x08, 0x1a, 0x93, 0x7e, 0x44, 0xe6, 0x5a, 0x88, 0x58, 0xf5, 0x25, 0x51, 0x49,
	0xac, 0x95, 0x16, 0x92, 0xf4, 0xe7, 0x83, 0x3e, 0xe3, 0x73, 0x11, 0x62, 0xcd, 0x04, 0xf7, 0x67,
	0x52, 0x68, 0x81, 0xee, 0xa4, 0x04, 0x3f, 0x23, 0xf8, 0x39, 0x82, 0x3f, 0x1f, 0xb4, 0x0f, 0x0b,
	0x04, 0x43, 0x31, 0x9d, 0x66, 0x62, 0xed, 0x67, 0x85, 0xc9, 0x73, 0x22, 0x31, 0x25, 0x23, 0x95,
	0x4c, 0xa7, 0x58, 0x2e, 0x1c, 0xed, 0xa0, 0x80, 0x36, 0x61, 0x31, 0x49, 0x53, 0xbb, 0xbf, 0xaa,
	0x00, 0xc3, 0xe5, 0x19, 0x10, 0x82, 0x2d, 0x8e, 0xa7, 0xc4, 0x2b, 0x75, 0x4a, 0xbd, 0x7a, 0x60,
	0xd7, 0xe8, 0x15, 0x94, 0x59, 0xe4, 0x95, 0x3b, 0xa5, 0x5e, 0x63, 0xf0, 0xd0, 0xff, 0xff, 0xf1,
	0xfc, 0x95, 0x96, 0x3f, 0x8c, 0x82, 0x32, 0x8b, 0xd0, 0x67, 0xd8, 0x53, 0x1a, 0xeb, 0x44, 0x8d,
	0xb0, 0xd6, 0x92, 0x8d, 0x13, 0x4d, 0x94, 0x57, 0xb1, 0x6a, 0x8f, 0x8a, 0xd4, 0xce, 0x2c, 0xf1,
	0xcd, 0x92, 0x17, 0xb4, 0xd4, 0x1a, 0x82, 0x5e, 0x43, 0x4d, 0xb3, 0x29, 0xe3, 0xd4, 0xdb, 0xb2,
	0x9a, 0x0f, 0x8a, 0x34, 0xcf, 0x6d, 0x76, 0xe0, 0x58, 0x88, 0xc1, 0x8d, 0xd5, 0x1d, 0xe6, 0x4b,
	0xac, 0x5a, 0xb9, 0xa7, 0x9b, 0x1f, 0x38, 0x57, 0xe6, 0x3e, 0xbb, 0x02, 0x45, 0xe7, 0x70, 0xed,
	0x42, 0xc8, 0x6f, 0x6a, 0x86, 0x43, 0x32, 0x62, 0x7c, 0x22, 0xbc, 0xda, 0x66, 0x4d, 0xfd, 0x98,
	0xb1, 0x86, 0x7c, 0x22, 0x82, 0xe6, 0x45, 0x3e, 0x44, 0xc7, 0x00, 0x33, 0x29, 0x66, 0x44, 0x6a,
	0x46, 0x94, 0xb7, 0xdd, 0xa9, 0xf4, 0x1a, 0x83, 0x5e, 0x91, 0xe2, 0x69, 0xca, 0x58, 0x04, 0x39,
	0x2e, 0x7a, 0x01, 0x55, 0xe3, 0x0c, 0xe5, 0xed, 0x58, 0x91, 0xfb, 0x45, 0x22, 0xef, 0x59, 0x4c,
	0x82, 0x94, 0x82, 0x26, 0x80, 0xd6, 0xcc, 0x68, 0xaa, 0xa9, 0x5b, 0xa1, 0xe7, 0x45, 0x42, 0x27,
	0x98, 0xd3, 0x04, 0x53, 0x72, 0xe4, 0x14, 0xce, 0x52, 0x37, 0x07, 0x7b, 0xe1, 0x25, 0x80, 0x11,
	0xd5, 0x3e, 0x80, 0xf2, 0x30, 0x42, 0xf7, 0xa0, 0x99, 0xbb, 0x34, 0x16, 0x39, 0xbf, 0xee, 0xae,
	0xc0, 0x61, 0xd4, 0x45, 0xd0, 0x5a, 0x36, 0xee, 0x48, 0x70, 0x4d, 0x7e, 0xe8, 0xee, 0xcf, 0x32,
	0x34, 0x2f, 0x75, 0xd3, 0xd8, 0x73, 0x75, 0x29, 0x61, 0x9a, 0x66, 0xe5, 0x36, 0xb0, 0xe7, 0xba,
	0x7c, 0xd0, 0xba, 0x58, 0x43, 0x50, 0x1b, 0x76, 0xbe, 0x08, 0xa5, 0xed, 0xa3, 0xaa, 0xd8, 0x22,
	0x97, 0x31, 0x3a, 0x4c, 0xb7, 0x66, 0x9c, 0x8e, 0x22, 0x26, 0x49, 0xa8, 0x85, 0x5c, 0x58, 0x17,
	0xd7, 0x53, 0x21, 0xc6, 0xe9, 0xbb, 0x0c, 0x47, 0xb7, 0x60, 0xc7, 0xd4, 0x30, 0xd2, 0x98, 0x5a,
	0x6b, 0xd6, 0x83, 0x6d, 0x13, 0x9f, 0x63, 0x8a, 0x4e, 0xa1, 0x69, 0xa6, 0x06, 0xe6, 0xd1, 0x28,
	0x66, 0x7c, 0x69, 0x82, 0xc3, 0xa2, 0xf2, 0x8f, 0x52, 0xd2, 0x09, 0xe3, 0x24, 0xd8, 0x0d, 0x57,
	0x81, 0xea, 0x12, 0x68, 0xe4, 0x3e, 0xa2, 0x7d, 0xa8, 0xc6, 0x78, 0x4c, 0x62, 0xd7, 0xe6, 0x34,
	0x30, 0xb3, 0xc2, 0xc8, 0xda, 0xc9, 0x50, 0x0f, 0xec, 0xda, 0x60, 0x58, 0x52, 0xf3, 0xbe, 0x2b,
	0x06, 0x33, 0x6b, 0xe4, 0xc1, 0xb6, 0x13, 0x77, 0x87, 0xcb, 0xc2, 0xee, 0xef, 0x12, 0xec, 0x5f,
	0xf5, 0x7e, 0xd0, 0x6d, 0xeb, 0xe9, 0xaf, 0x24, 0xd4, 0xab, 0xcb, 0xad, 0x3b, 0x64, 0x18, 0x99,
	0x7a, 0x12, 0x45, 0xa4, 0xf2, 0xca, 0x76, 0x9b, 0x34, 0x40, 0x37, 0xa1, 0x66, 0x0b, 0xcb, 0x76,
	0x77, 0x11, 0xea, 0x40, 0x23, 0x22, 0x2a, 0x94, 0x6c, 0x66, 0x76, 0x71, 0x35, 0xe4, 0x21, 0x34,
	0x86, 0xeb, 0x39, 0x3b, 0x39, 0x13, 0x28, 0xaf, 0x66, 0xdb, 0xf8, 0x78, 0xf3, 0x09, 0x90, 0xd9,
	0x00, 0xb1, 0x75, 0x48, 0x75, 0x8f, 0x61, 0xef, 0xaf, 0x44, 0x74, 0x17, 0x76, 0x23, 0xa6, 0x66,
	0x31, 0x5e, 0x8c, 0x72, 0x63, 0xb7, 0xe1, 0xb0, 0x0f, 0xc6, 0x24, 0x2d, 0xa8, 0x24, 0x32, 0x6b,
	0xb2, 0x59, 0xbe, 0xfd, 0x0e, 0xdd, 0x50, 0x4c, 0x0b, 0xaa, 0x3a, 0x2d, 0x7d, 0x1a, 0xba, 0x0c,
	0x2a, 0x62, 0xcc, 0xa9, 0x2f, 0x24, 0xed, 0x53, 0xc2, 0xed, 0xd8, 0x77, 0x7f, 0x35, 0x3c, 0x63,
	0xea, 0x5f, 0x3f, 0x89, 0x97, 0xb9, 0x70, 0x5c, 0xb3, 0xac, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x9c, 0x60, 0x5f, 0x0e, 0x07, 0x00, 0x00,
}