// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firecracker.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CreateVMRequest specifies creation parameters for a new FC instance
type CreateVMRequest struct {
	// VM identifier to assign
	VMID string `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	// Specifies the machine configuration for the VM
	MachineCfg *FirecrackerMachineConfiguration `protobuf:"bytes,2,opt,name=MachineCfg,json=machineCfg,proto3" json:"MachineCfg,omitempty"`
	// Specifies the file path where the kernel image is located
	KernelImagePath string `protobuf:"bytes,3,opt,name=KernelImagePath,json=kernelImagePath,proto3" json:"KernelImagePath,omitempty"`
	// Specifies the commandline arguments that should be passed to the kernel
	KernelArgs string `protobuf:"bytes,4,opt,name=KernelArgs,json=kernelArgs,proto3" json:"KernelArgs,omitempty"`
	// Specifies drive containing the rootfs of the VM
	RootDrive *FirecrackerRootDrive `protobuf:"bytes,5,opt,name=RootDrive,json=rootDrive,proto3" json:"RootDrive,omitempty"`
	// Specifies additional drives whose contents will be mounted inside the VM on boot.
	DriveMounts []*FirecrackerDriveMount `protobuf:"bytes,6,rep,name=DriveMounts,json=driveMounts,proto3" json:"DriveMounts,omitempty"`
	// Specifies the networking configuration for a VM
	NetworkInterfaces []*FirecrackerNetworkInterface `protobuf:"bytes,7,rep,name=NetworkInterfaces,json=networkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	// The number of dummy drives to reserve in advance before running FC instance.
	ContainerCount int32 `protobuf:"varint,8,opt,name=ContainerCount,json=containerCount,proto3" json:"ContainerCount,omitempty"`
	// Whether the VM should exit after all tasks running in it have been deleted.
	ExitAfterAllTasksDeleted bool          `protobuf:"varint,9,opt,name=ExitAfterAllTasksDeleted,json=exitAfterAllTasksDeleted,proto3" json:"ExitAfterAllTasksDeleted,omitempty"`
	JailerConfig             *JailerConfig `protobuf:"bytes,10,opt,name=JailerConfig,json=jailerConfig,proto3" json:"JailerConfig,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}      `json:"-"`
	XXX_unrecognized         []byte        `json:"-"`
	XXX_sizecache            int32         `json:"-"`
}

func (m *CreateVMRequest) Reset()         { *m = CreateVMRequest{} }
func (m *CreateVMRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMRequest) ProtoMessage()    {}
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{0}
}
func (m *CreateVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMRequest.Unmarshal(m, b)
}
func (m *CreateVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMRequest.Marshal(b, m, deterministic)
}
func (m *CreateVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMRequest.Merge(m, src)
}
func (m *CreateVMRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMRequest.Size(m)
}
func (m *CreateVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMRequest proto.InternalMessageInfo

func (m *CreateVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *CreateVMRequest) GetMachineCfg() *FirecrackerMachineConfiguration {
	if m != nil {
		return m.MachineCfg
	}
	return nil
}

func (m *CreateVMRequest) GetKernelImagePath() string {
	if m != nil {
		return m.KernelImagePath
	}
	return ""
}

func (m *CreateVMRequest) GetKernelArgs() string {
	if m != nil {
		return m.KernelArgs
	}
	return ""
}

func (m *CreateVMRequest) GetRootDrive() *FirecrackerRootDrive {
	if m != nil {
		return m.RootDrive
	}
	return nil
}

func (m *CreateVMRequest) GetDriveMounts() []*FirecrackerDriveMount {
	if m != nil {
		return m.DriveMounts
	}
	return nil
}

func (m *CreateVMRequest) GetNetworkInterfaces() []*FirecrackerNetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *CreateVMRequest) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

func (m *CreateVMRequest) GetExitAfterAllTasksDeleted() bool {
	if m != nil {
		return m.ExitAfterAllTasksDeleted
	}
	return false
}

func (m *CreateVMRequest) GetJailerConfig() *JailerConfig {
	if m != nil {
		return m.JailerConfig
	}
	return nil
}

type CreateVMResponse struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	SocketPath           string   `protobuf:"bytes,2,opt,name=SocketPath,json=socketPath,proto3" json:"SocketPath,omitempty"`
	LogFifoPath          string   `protobuf:"bytes,3,opt,name=LogFifoPath,json=logFifoPath,proto3" json:"LogFifoPath,omitempty"`
	MetricsFifoPath      string   `protobuf:"bytes,4,opt,name=MetricsFifoPath,json=metricsFifoPath,proto3" json:"MetricsFifoPath,omitempty"`
	CgroupPath           string   `protobuf:"bytes,5,opt,name=CgroupPath,json=cgroupPath,proto3" json:"CgroupPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMResponse) Reset()         { *m = CreateVMResponse{} }
func (m *CreateVMResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVMResponse) ProtoMessage()    {}
func (*CreateVMResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{1}
}
func (m *CreateVMResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMResponse.Unmarshal(m, b)
}
func (m *CreateVMResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMResponse.Marshal(b, m, deterministic)
}
func (m *CreateVMResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMResponse.Merge(m, src)
}
func (m *CreateVMResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVMResponse.Size(m)
}
func (m *CreateVMResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMResponse proto.InternalMessageInfo

func (m *CreateVMResponse) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *CreateVMResponse) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func (m *CreateVMResponse) GetLogFifoPath() string {
	if m != nil {
		return m.LogFifoPath
	}
	return ""
}

func (m *CreateVMResponse) GetMetricsFifoPath() string {
	if m != nil {
		return m.MetricsFifoPath
	}
	return ""
}

func (m *CreateVMResponse) GetCgroupPath() string {
	if m != nil {
		return m.CgroupPath
	}
	return ""
}

type StopVMRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	TimeoutSeconds       uint32   `protobuf:"varint,2,opt,name=TimeoutSeconds,json=timeoutSeconds,proto3" json:"TimeoutSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopVMRequest) Reset()         { *m = StopVMRequest{} }
func (m *StopVMRequest) String() string { return proto.CompactTextString(m) }
func (*StopVMRequest) ProtoMessage()    {}
func (*StopVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{2}
}
func (m *StopVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopVMRequest.Unmarshal(m, b)
}
func (m *StopVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopVMRequest.Marshal(b, m, deterministic)
}
func (m *StopVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopVMRequest.Merge(m, src)
}
func (m *StopVMRequest) XXX_Size() int {
	return xxx_messageInfo_StopVMRequest.Size(m)
}
func (m *StopVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopVMRequest proto.InternalMessageInfo

func (m *StopVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *StopVMRequest) GetTimeoutSeconds() uint32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

type GetVMInfoRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoRequest) Reset()         { *m = GetVMInfoRequest{} }
func (m *GetVMInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoRequest) ProtoMessage()    {}
func (*GetVMInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{3}
}
func (m *GetVMInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoRequest.Unmarshal(m, b)
}
func (m *GetVMInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetVMInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoRequest.Merge(m, src)
}
func (m *GetVMInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoRequest.Size(m)
}
func (m *GetVMInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoRequest proto.InternalMessageInfo

func (m *GetVMInfoRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMInfoResponse struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	SocketPath           string   `protobuf:"bytes,2,opt,name=SocketPath,json=socketPath,proto3" json:"SocketPath,omitempty"`
	LogFifoPath          string   `protobuf:"bytes,3,opt,name=LogFifoPath,json=logFifoPath,proto3" json:"LogFifoPath,omitempty"`
	MetricsFifoPath      string   `protobuf:"bytes,4,opt,name=MetricsFifoPath,json=metricsFifoPath,proto3" json:"MetricsFifoPath,omitempty"`
	CgroupPath           string   `protobuf:"bytes,5,opt,name=CgroupPath,json=cgroupPath,proto3" json:"CgroupPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoResponse) Reset()         { *m = GetVMInfoResponse{} }
func (m *GetVMInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoResponse) ProtoMessage()    {}
func (*GetVMInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{4}
}
func (m *GetVMInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoResponse.Unmarshal(m, b)
}
func (m *GetVMInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetVMInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoResponse.Merge(m, src)
}
func (m *GetVMInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoResponse.Size(m)
}
func (m *GetVMInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoResponse proto.InternalMessageInfo

func (m *GetVMInfoResponse) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *GetVMInfoResponse) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func (m *GetVMInfoResponse) GetLogFifoPath() string {
	if m != nil {
		return m.LogFifoPath
	}
	return ""
}

func (m *GetVMInfoResponse) GetMetricsFifoPath() string {
	if m != nil {
		return m.MetricsFifoPath
	}
	return ""
}

func (m *GetVMInfoResponse) GetCgroupPath() string {
	if m != nil {
		return m.CgroupPath
	}
	return ""
}

type SetVMMetadataRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetVMMetadataRequest) Reset()         { *m = SetVMMetadataRequest{} }
func (m *SetVMMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*SetVMMetadataRequest) ProtoMessage()    {}
func (*SetVMMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{5}
}
func (m *SetVMMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetVMMetadataRequest.Unmarshal(m, b)
}
func (m *SetVMMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetVMMetadataRequest.Marshal(b, m, deterministic)
}
func (m *SetVMMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVMMetadataRequest.Merge(m, src)
}
func (m *SetVMMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_SetVMMetadataRequest.Size(m)
}
func (m *SetVMMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVMMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVMMetadataRequest proto.InternalMessageInfo

func (m *SetVMMetadataRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *SetVMMetadataRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type UpdateVMMetadataRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVMMetadataRequest) Reset()         { *m = UpdateVMMetadataRequest{} }
func (m *UpdateVMMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVMMetadataRequest) ProtoMessage()    {}
func (*UpdateVMMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{6}
}
func (m *UpdateVMMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVMMetadataRequest.Unmarshal(m, b)
}
func (m *UpdateVMMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVMMetadataRequest.Marshal(b, m, deterministic)
}
func (m *UpdateVMMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVMMetadataRequest.Merge(m, src)
}
func (m *UpdateVMMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVMMetadataRequest.Size(m)
}
func (m *UpdateVMMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVMMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVMMetadataRequest proto.InternalMessageInfo

func (m *UpdateVMMetadataRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *UpdateVMMetadataRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type GetVMMetadataRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMMetadataRequest) Reset()         { *m = GetVMMetadataRequest{} }
func (m *GetVMMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMMetadataRequest) ProtoMessage()    {}
func (*GetVMMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{7}
}
func (m *GetVMMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMMetadataRequest.Unmarshal(m, b)
}
func (m *GetVMMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMMetadataRequest.Marshal(b, m, deterministic)
}
func (m *GetVMMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMMetadataRequest.Merge(m, src)
}
func (m *GetVMMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMMetadataRequest.Size(m)
}
func (m *GetVMMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMMetadataRequest proto.InternalMessageInfo

func (m *GetVMMetadataRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMMetadataResponse struct {
	Metadata             string   `protobuf:"bytes,1,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMMetadataResponse) Reset()         { *m = GetVMMetadataResponse{} }
func (m *GetVMMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*GetVMMetadataResponse) ProtoMessage()    {}
func (*GetVMMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{8}
}
func (m *GetVMMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMMetadataResponse.Unmarshal(m, b)
}
func (m *GetVMMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMMetadataResponse.Marshal(b, m, deterministic)
}
func (m *GetVMMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMMetadataResponse.Merge(m, src)
}
func (m *GetVMMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_GetVMMetadataResponse.Size(m)
}
func (m *GetVMMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMMetadataResponse proto.InternalMessageInfo

func (m *GetVMMetadataResponse) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type JailerConfig struct {
	NetNS string `protobuf:"bytes,1,opt,name=NetNS,json=netNS,proto3" json:"NetNS,omitempty"`
	// List of the physical numbers of the CPUs on which processes in that
	// cpuset are allowed to execute.  See List Format below for a description
	// of the format of cpus.
	//
	// The CPUs allowed to a cpuset may be changed by writing a new list to its
	// cpus file.
	// Taken from http://man7.org/linux/man-pages/man7/cpuset.7.html
	//
	// This is formatted as specified in the cpuset man page under "List Format"
	// http://man7.org/linux/man-pages/man7/cpuset.7.html
	CPUs string `protobuf:"bytes,2,opt,name=CPUs,json=cPUs,proto3" json:"CPUs,omitempty"`
	// List of memory nodes on which processes in this cpuset are allowed to
	// allocate memory.  See List Format below for a description of the format
	// of mems.
	// Taken from http://man7.org/linux/man-pages/man7/cpuset.7.html
	//
	// This is formatted as specified in the cpuset man page under "List Format"
	// http://man7.org/linux/man-pages/man7/cpuset.7.html
	Mems string `protobuf:"bytes,3,opt,name=Mems,json=mems,proto3" json:"Mems,omitempty"`
	UID  uint32 `protobuf:"varint,4,opt,name=UID,json=uID,proto3" json:"UID,omitempty"`
	GID  uint32 `protobuf:"varint,5,opt,name=GID,json=gID,proto3" json:"GID,omitempty"`
	// CgroupPath is used to dictate where the cgroup should be located
	// relative to the cgroup directory which is
	// /sys/fs/cgroup/cpu/<CgroupPath>/<vmID>
	// if no value was provided, then /firecracker-containerd will be used as
	// the default value
	CgroupPath           string   `protobuf:"bytes,6,opt,name=CgroupPath,json=cgroupPath,proto3" json:"CgroupPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JailerConfig) Reset()         { *m = JailerConfig{} }
func (m *JailerConfig) String() string { return proto.CompactTextString(m) }
func (*JailerConfig) ProtoMessage()    {}
func (*JailerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73317e9fb8da571, []int{9}
}
func (m *JailerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JailerConfig.Unmarshal(m, b)
}
func (m *JailerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JailerConfig.Marshal(b, m, deterministic)
}
func (m *JailerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JailerConfig.Merge(m, src)
}
func (m *JailerConfig) XXX_Size() int {
	return xxx_messageInfo_JailerConfig.Size(m)
}
func (m *JailerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_JailerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_JailerConfig proto.InternalMessageInfo

func (m *JailerConfig) GetNetNS() string {
	if m != nil {
		return m.NetNS
	}
	return ""
}

func (m *JailerConfig) GetCPUs() string {
	if m != nil {
		return m.CPUs
	}
	return ""
}

func (m *JailerConfig) GetMems() string {
	if m != nil {
		return m.Mems
	}
	return ""
}

func (m *JailerConfig) GetUID() uint32 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *JailerConfig) GetGID() uint32 {
	if m != nil {
		return m.GID
	}
	return 0
}

func (m *JailerConfig) GetCgroupPath() string {
	if m != nil {
		return m.CgroupPath
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVMRequest)(nil), "CreateVMRequest")
	proto.RegisterType((*CreateVMResponse)(nil), "CreateVMResponse")
	proto.RegisterType((*StopVMRequest)(nil), "StopVMRequest")
	proto.RegisterType((*GetVMInfoRequest)(nil), "GetVMInfoRequest")
	proto.RegisterType((*GetVMInfoResponse)(nil), "GetVMInfoResponse")
	proto.RegisterType((*SetVMMetadataRequest)(nil), "SetVMMetadataRequest")
	proto.RegisterType((*UpdateVMMetadataRequest)(nil), "UpdateVMMetadataRequest")
	proto.RegisterType((*GetVMMetadataRequest)(nil), "GetVMMetadataRequest")
	proto.RegisterType((*GetVMMetadataResponse)(nil), "GetVMMetadataResponse")
	proto.RegisterType((*JailerConfig)(nil), "JailerConfig")
}

func init() { proto.RegisterFile("firecracker.proto", fileDescriptor_a73317e9fb8da571) }

var fileDescriptor_a73317e9fb8da571 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xa4, 0x4d, 0xc6, 0x4d, 0xd2, 0xac, 0x5a, 0x58, 0x55, 0xa8, 0xb2, 0x7c, 0xa8,
	0x22, 0x0e, 0x95, 0x68, 0x2f, 0x88, 0x13, 0x25, 0xa6, 0x95, 0x5b, 0x5c, 0x55, 0x9b, 0xa6, 0x07,
	0x6e, 0x8b, 0x33, 0x76, 0x4d, 0xe2, 0xdd, 0xb0, 0xbb, 0x29, 0x70, 0xe6, 0x07, 0xf8, 0x16, 0xf8,
	0x41, 0xe4, 0x4d, 0x9a, 0x38, 0x11, 0x54, 0x48, 0x9c, 0x38, 0x65, 0xe7, 0xcd, 0x9b, 0x99, 0xa7,
	0x79, 0x13, 0x43, 0x27, 0xc9, 0x14, 0xc6, 0x8a, 0xc7, 0x23, 0x54, 0x47, 0x13, 0x25, 0x8d, 0xdc,
	0x77, 0xcd, 0xd7, 0x09, 0xea, 0x59, 0xe0, 0x7f, 0xab, 0x42, 0xbb, 0xa7, 0x90, 0x1b, 0xbc, 0x8d,
	0x18, 0x7e, 0x9a, 0xa2, 0x36, 0x84, 0x40, 0xf5, 0x36, 0x0a, 0x03, 0xea, 0x78, 0x4e, 0xb7, 0xc1,
	0xaa, 0xf7, 0x51, 0x18, 0x90, 0xd7, 0x00, 0x11, 0x8f, 0xef, 0x32, 0x81, 0xbd, 0x24, 0xa5, 0x1b,
	0x9e, 0xd3, 0x75, 0x8f, 0xbd, 0xa3, 0xb3, 0x65, 0xf3, 0x87, 0xac, 0x14, 0x49, 0x96, 0x4e, 0x15,
	0x37, 0x99, 0x14, 0x0c, 0xf2, 0x45, 0x0d, 0xe9, 0x42, 0xfb, 0x12, 0x95, 0xc0, 0x71, 0x98, 0xf3,
	0x14, 0xaf, 0xb9, 0xb9, 0xa3, 0x15, 0x3b, 0xa0, 0x3d, 0x5a, 0x85, 0xc9, 0x01, 0xc0, 0x8c, 0x79,
	0xaa, 0x52, 0x4d, 0xab, 0x96, 0x04, 0xa3, 0x05, 0x42, 0x4e, 0xa0, 0xc1, 0xa4, 0x34, 0x81, 0xca,
	0xee, 0x91, 0xd6, 0xac, 0x94, 0xbd, 0xb2, 0x94, 0x45, 0x92, 0x35, 0xd4, 0xc3, 0x93, 0xbc, 0x04,
	0xd7, 0x3e, 0x22, 0x39, 0x15, 0x46, 0xd3, 0x4d, 0xaf, 0xd2, 0x75, 0x8f, 0x9f, 0x94, 0xcb, 0x96,
	0x69, 0xe6, 0x0e, 0x97, 0x54, 0x72, 0x01, 0x9d, 0x2b, 0x34, 0x9f, 0xa5, 0x1a, 0x85, 0xc2, 0xa0,
	0x4a, 0x78, 0x8c, 0x9a, 0x6e, 0xd9, 0xfa, 0x67, 0xe5, 0xfa, 0x75, 0x12, 0xeb, 0x88, 0xf5, 0x32,
	0x72, 0x08, 0xad, 0x9e, 0x14, 0x86, 0x67, 0x02, 0x55, 0xaf, 0x68, 0x4f, 0xeb, 0x9e, 0xd3, 0xad,
	0xb1, 0x56, 0xbc, 0x82, 0x92, 0x57, 0x40, 0xdf, 0x7e, 0xc9, 0xcc, 0x69, 0x62, 0x50, 0x9d, 0x8e,
	0xc7, 0x37, 0x5c, 0x8f, 0x74, 0x80, 0x63, 0x34, 0x38, 0xa4, 0x0d, 0xcf, 0xe9, 0xd6, 0x19, 0xc5,
	0x3f, 0xe4, 0xc9, 0x0b, 0xd8, 0xbe, 0xe0, 0xd9, 0xb8, 0x68, 0x55, 0x78, 0x41, 0xc1, 0x6e, 0xa8,
	0x79, 0x54, 0x06, 0xd9, 0xf6, 0xc7, 0x52, 0xe4, 0xff, 0x70, 0x60, 0x67, 0x79, 0x05, 0x7a, 0x22,
	0x85, 0xc6, 0xdf, 0x9e, 0xc1, 0x01, 0x40, 0x5f, 0xc6, 0x23, 0x34, 0xd6, 0xbf, 0x8d, 0x99, 0x35,
	0x7a, 0x81, 0x10, 0x0f, 0xdc, 0x77, 0x32, 0x3d, 0xcb, 0x12, 0x59, 0x32, 0xd8, 0x1d, 0x2f, 0xa1,
	0xe2, 0x0c, 0x22, 0x34, 0x2a, 0x8b, 0xf5, 0x82, 0x35, 0x73, 0xb8, 0x9d, 0xaf, 0xc2, 0xc5, 0xac,
	0x5e, 0xaa, 0xe4, 0x74, 0x62, 0x49, 0xb5, 0xd9, 0xac, 0x78, 0x81, 0xf8, 0x97, 0xd0, 0xec, 0x1b,
	0x39, 0x79, 0xfc, 0x6e, 0x0f, 0xa1, 0x75, 0x93, 0xe5, 0x28, 0xa7, 0xa6, 0x8f, 0xb1, 0x14, 0x43,
	0x6d, 0x45, 0x37, 0x59, 0xcb, 0xac, 0xa0, 0xfe, 0x21, 0xec, 0x9c, 0xa3, 0xb9, 0x8d, 0x42, 0x91,
	0xc8, 0x47, 0xfa, 0xf9, 0x3f, 0x1d, 0xe8, 0x94, 0x88, 0xff, 0xc9, 0xaa, 0xce, 0x60, 0xb7, 0x5f,
	0x88, 0x8e, 0xd0, 0xf0, 0x21, 0x37, 0xfc, 0xb1, 0x8d, 0xed, 0x43, 0xfd, 0x81, 0x36, 0x57, 0x5d,
	0xcf, 0xe7, 0xb1, 0x1f, 0xc2, 0xd3, 0xc1, 0x64, 0x68, 0xcf, 0xe4, 0x5f, 0x5b, 0x3d, 0x87, 0xdd,
	0xf3, 0xbf, 0x94, 0xe4, 0x9f, 0xc0, 0xde, 0x1a, 0x77, 0xbe, 0xf7, 0xf2, 0x00, 0x67, 0x6d, 0xc0,
	0x77, 0x67, 0xf5, 0x7f, 0x40, 0x76, 0xa1, 0x76, 0x85, 0xe6, 0xaa, 0x3f, 0x67, 0xd6, 0x44, 0x11,
	0x14, 0xf3, 0x7a, 0xd7, 0x03, 0x3d, 0xd7, 0x57, 0x8d, 0xaf, 0x07, 0xba, 0xc0, 0x22, 0xcc, 0xf5,
	0xdc, 0x93, 0x6a, 0x8e, 0xb9, 0x26, 0x3b, 0x50, 0x19, 0x84, 0x81, 0x35, 0xa0, 0xc9, 0x2a, 0xd3,
	0x30, 0x28, 0x90, 0xf3, 0x30, 0xb0, 0xdb, 0x6e, 0xb2, 0x4a, 0x3a, 0xb3, 0xbc, 0x64, 0xc3, 0xe6,
	0xba, 0x0d, 0x6f, 0xb6, 0xde, 0xd7, 0xec, 0x57, 0xf7, 0xc3, 0xa6, 0xfd, 0x39, 0xf9, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xa5, 0x2f, 0xe8, 0x9e, 0x05, 0x00, 0x00,
}
