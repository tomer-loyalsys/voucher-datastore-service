// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/voucherds.proto

package lsvoucherds

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// Ping rpc to check that handler respond
type PingReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{0}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

type PingRes struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRes) Reset()         { *m = PingRes{} }
func (m *PingRes) String() string { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()    {}
func (*PingRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{1}
}

func (m *PingRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRes.Unmarshal(m, b)
}
func (m *PingRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRes.Marshal(b, m, deterministic)
}
func (m *PingRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRes.Merge(m, src)
}
func (m *PingRes) XXX_Size() int {
	return xxx_messageInfo_PingRes.Size(m)
}
func (m *PingRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRes.DiscardUnknown(m)
}

var xxx_messageInfo_PingRes proto.InternalMessageInfo

func (m *PingRes) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Test rpc
type TestReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{2}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

type TestRes struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRes) Reset()         { *m = TestRes{} }
func (m *TestRes) String() string { return proto.CompactTextString(m) }
func (*TestRes) ProtoMessage()    {}
func (*TestRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{3}
}

func (m *TestRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRes.Unmarshal(m, b)
}
func (m *TestRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRes.Marshal(b, m, deterministic)
}
func (m *TestRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRes.Merge(m, src)
}
func (m *TestRes) XXX_Size() int {
	return xxx_messageInfo_TestRes.Size(m)
}
func (m *TestRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRes.DiscardUnknown(m)
}

var xxx_messageInfo_TestRes proto.InternalMessageInfo

func (m *TestRes) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

// Add upload vouchers rpc
type UploadToPoolReq struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	PoolId               string   `protobuf:"bytes,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
	UploadId             string   `protobuf:"bytes,4,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	Vouchers             []string `protobuf:"bytes,5,rep,name=vouchers,proto3" json:"vouchers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadToPoolReq) Reset()         { *m = UploadToPoolReq{} }
func (m *UploadToPoolReq) String() string { return proto.CompactTextString(m) }
func (*UploadToPoolReq) ProtoMessage()    {}
func (*UploadToPoolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{4}
}

func (m *UploadToPoolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadToPoolReq.Unmarshal(m, b)
}
func (m *UploadToPoolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadToPoolReq.Marshal(b, m, deterministic)
}
func (m *UploadToPoolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadToPoolReq.Merge(m, src)
}
func (m *UploadToPoolReq) XXX_Size() int {
	return xxx_messageInfo_UploadToPoolReq.Size(m)
}
func (m *UploadToPoolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadToPoolReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadToPoolReq proto.InternalMessageInfo

func (m *UploadToPoolReq) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UploadToPoolReq) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadToPoolReq) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *UploadToPoolReq) GetUploadId() string {
	if m != nil {
		return m.UploadId
	}
	return ""
}

func (m *UploadToPoolReq) GetVouchers() []string {
	if m != nil {
		return m.Vouchers
	}
	return nil
}

type UploadToPoolRes struct {
	TotalUpload          int64    `protobuf:"varint,1,opt,name=totalUpload,proto3" json:"totalUpload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadToPoolRes) Reset()         { *m = UploadToPoolRes{} }
func (m *UploadToPoolRes) String() string { return proto.CompactTextString(m) }
func (*UploadToPoolRes) ProtoMessage()    {}
func (*UploadToPoolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{5}
}

func (m *UploadToPoolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadToPoolRes.Unmarshal(m, b)
}
func (m *UploadToPoolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadToPoolRes.Marshal(b, m, deterministic)
}
func (m *UploadToPoolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadToPoolRes.Merge(m, src)
}
func (m *UploadToPoolRes) XXX_Size() int {
	return xxx_messageInfo_UploadToPoolRes.Size(m)
}
func (m *UploadToPoolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadToPoolRes.DiscardUnknown(m)
}

var xxx_messageInfo_UploadToPoolRes proto.InternalMessageInfo

func (m *UploadToPoolRes) GetTotalUpload() int64 {
	if m != nil {
		return m.TotalUpload
	}
	return 0
}

// Get pool status rpc
type GetPoolAvailabilityReq struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	PoolIds              []string `protobuf:"bytes,3,rep,name=poolIds,proto3" json:"poolIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolAvailabilityReq) Reset()         { *m = GetPoolAvailabilityReq{} }
func (m *GetPoolAvailabilityReq) String() string { return proto.CompactTextString(m) }
func (*GetPoolAvailabilityReq) ProtoMessage()    {}
func (*GetPoolAvailabilityReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{6}
}

func (m *GetPoolAvailabilityReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolAvailabilityReq.Unmarshal(m, b)
}
func (m *GetPoolAvailabilityReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolAvailabilityReq.Marshal(b, m, deterministic)
}
func (m *GetPoolAvailabilityReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolAvailabilityReq.Merge(m, src)
}
func (m *GetPoolAvailabilityReq) XXX_Size() int {
	return xxx_messageInfo_GetPoolAvailabilityReq.Size(m)
}
func (m *GetPoolAvailabilityReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolAvailabilityReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolAvailabilityReq proto.InternalMessageInfo

func (m *GetPoolAvailabilityReq) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *GetPoolAvailabilityReq) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GetPoolAvailabilityReq) GetPoolIds() []string {
	if m != nil {
		return m.PoolIds
	}
	return nil
}

type GetPoolAvailabilityRes struct {
	PoolAvailability     map[string]*GetPoolAvailabilityRes_GetPoolAvailability `protobuf:"bytes,1,rep,name=PoolAvailability,proto3" json:"PoolAvailability,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *GetPoolAvailabilityRes) Reset()         { *m = GetPoolAvailabilityRes{} }
func (m *GetPoolAvailabilityRes) String() string { return proto.CompactTextString(m) }
func (*GetPoolAvailabilityRes) ProtoMessage()    {}
func (*GetPoolAvailabilityRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{7}
}

func (m *GetPoolAvailabilityRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolAvailabilityRes.Unmarshal(m, b)
}
func (m *GetPoolAvailabilityRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolAvailabilityRes.Marshal(b, m, deterministic)
}
func (m *GetPoolAvailabilityRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolAvailabilityRes.Merge(m, src)
}
func (m *GetPoolAvailabilityRes) XXX_Size() int {
	return xxx_messageInfo_GetPoolAvailabilityRes.Size(m)
}
func (m *GetPoolAvailabilityRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolAvailabilityRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolAvailabilityRes proto.InternalMessageInfo

func (m *GetPoolAvailabilityRes) GetPoolAvailability() map[string]*GetPoolAvailabilityRes_GetPoolAvailability {
	if m != nil {
		return m.PoolAvailability
	}
	return nil
}

type GetPoolAvailabilityRes_GetPoolAvailability struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Available            int64    `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolAvailabilityRes_GetPoolAvailability) Reset() {
	*m = GetPoolAvailabilityRes_GetPoolAvailability{}
}
func (m *GetPoolAvailabilityRes_GetPoolAvailability) String() string {
	return proto.CompactTextString(m)
}
func (*GetPoolAvailabilityRes_GetPoolAvailability) ProtoMessage() {}
func (*GetPoolAvailabilityRes_GetPoolAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{7, 0}
}

func (m *GetPoolAvailabilityRes_GetPoolAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability.Unmarshal(m, b)
}
func (m *GetPoolAvailabilityRes_GetPoolAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability.Marshal(b, m, deterministic)
}
func (m *GetPoolAvailabilityRes_GetPoolAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability.Merge(m, src)
}
func (m *GetPoolAvailabilityRes_GetPoolAvailability) XXX_Size() int {
	return xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability.Size(m)
}
func (m *GetPoolAvailabilityRes_GetPoolAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolAvailabilityRes_GetPoolAvailability proto.InternalMessageInfo

func (m *GetPoolAvailabilityRes_GetPoolAvailability) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetPoolAvailabilityRes_GetPoolAvailability) GetAvailable() int64 {
	if m != nil {
		return m.Available
	}
	return 0
}

// Delete all pool uploads vouchers rpc
type DeletePoolReq struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	PoolId               string   `protobuf:"bytes,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePoolReq) Reset()         { *m = DeletePoolReq{} }
func (m *DeletePoolReq) String() string { return proto.CompactTextString(m) }
func (*DeletePoolReq) ProtoMessage()    {}
func (*DeletePoolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{8}
}

func (m *DeletePoolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePoolReq.Unmarshal(m, b)
}
func (m *DeletePoolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePoolReq.Marshal(b, m, deterministic)
}
func (m *DeletePoolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePoolReq.Merge(m, src)
}
func (m *DeletePoolReq) XXX_Size() int {
	return xxx_messageInfo_DeletePoolReq.Size(m)
}
func (m *DeletePoolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePoolReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePoolReq proto.InternalMessageInfo

func (m *DeletePoolReq) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *DeletePoolReq) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *DeletePoolReq) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type DeletePoolRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePoolRes) Reset()         { *m = DeletePoolRes{} }
func (m *DeletePoolRes) String() string { return proto.CompactTextString(m) }
func (*DeletePoolRes) ProtoMessage()    {}
func (*DeletePoolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{9}
}

func (m *DeletePoolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePoolRes.Unmarshal(m, b)
}
func (m *DeletePoolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePoolRes.Marshal(b, m, deterministic)
}
func (m *DeletePoolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePoolRes.Merge(m, src)
}
func (m *DeletePoolRes) XXX_Size() int {
	return xxx_messageInfo_DeletePoolRes.Size(m)
}
func (m *DeletePoolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePoolRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePoolRes proto.InternalMessageInfo

// pop voucher from pool list rpc
type PopFromPoolReq struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	PoolId               string   `protobuf:"bytes,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopFromPoolReq) Reset()         { *m = PopFromPoolReq{} }
func (m *PopFromPoolReq) String() string { return proto.CompactTextString(m) }
func (*PopFromPoolReq) ProtoMessage()    {}
func (*PopFromPoolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{10}
}

func (m *PopFromPoolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopFromPoolReq.Unmarshal(m, b)
}
func (m *PopFromPoolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopFromPoolReq.Marshal(b, m, deterministic)
}
func (m *PopFromPoolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopFromPoolReq.Merge(m, src)
}
func (m *PopFromPoolReq) XXX_Size() int {
	return xxx_messageInfo_PopFromPoolReq.Size(m)
}
func (m *PopFromPoolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PopFromPoolReq.DiscardUnknown(m)
}

var xxx_messageInfo_PopFromPoolReq proto.InternalMessageInfo

func (m *PopFromPoolReq) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *PopFromPoolReq) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *PopFromPoolReq) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type PopFromPoolRes struct {
	Voucher              string   `protobuf:"bytes,1,opt,name=voucher,proto3" json:"voucher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopFromPoolRes) Reset()         { *m = PopFromPoolRes{} }
func (m *PopFromPoolRes) String() string { return proto.CompactTextString(m) }
func (*PopFromPoolRes) ProtoMessage()    {}
func (*PopFromPoolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e572ab4fba54a2c, []int{11}
}

func (m *PopFromPoolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopFromPoolRes.Unmarshal(m, b)
}
func (m *PopFromPoolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopFromPoolRes.Marshal(b, m, deterministic)
}
func (m *PopFromPoolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopFromPoolRes.Merge(m, src)
}
func (m *PopFromPoolRes) XXX_Size() int {
	return xxx_messageInfo_PopFromPoolRes.Size(m)
}
func (m *PopFromPoolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PopFromPoolRes.DiscardUnknown(m)
}

var xxx_messageInfo_PopFromPoolRes proto.InternalMessageInfo

func (m *PopFromPoolRes) GetVoucher() string {
	if m != nil {
		return m.Voucher
	}
	return ""
}

func init() {
	proto.RegisterType((*PingReq)(nil), "lsvoucherds.PingReq")
	proto.RegisterType((*PingRes)(nil), "lsvoucherds.PingRes")
	proto.RegisterType((*TestReq)(nil), "lsvoucherds.TestReq")
	proto.RegisterType((*TestRes)(nil), "lsvoucherds.TestRes")
	proto.RegisterType((*UploadToPoolReq)(nil), "lsvoucherds.UploadToPoolReq")
	proto.RegisterType((*UploadToPoolRes)(nil), "lsvoucherds.UploadToPoolRes")
	proto.RegisterType((*GetPoolAvailabilityReq)(nil), "lsvoucherds.GetPoolAvailabilityReq")
	proto.RegisterType((*GetPoolAvailabilityRes)(nil), "lsvoucherds.GetPoolAvailabilityRes")
	proto.RegisterMapType((map[string]*GetPoolAvailabilityRes_GetPoolAvailability)(nil), "lsvoucherds.GetPoolAvailabilityRes.PoolAvailabilityEntry")
	proto.RegisterType((*GetPoolAvailabilityRes_GetPoolAvailability)(nil), "lsvoucherds.GetPoolAvailabilityRes.GetPoolAvailability")
	proto.RegisterType((*DeletePoolReq)(nil), "lsvoucherds.DeletePoolReq")
	proto.RegisterType((*DeletePoolRes)(nil), "lsvoucherds.DeletePoolRes")
	proto.RegisterType((*PopFromPoolReq)(nil), "lsvoucherds.PopFromPoolReq")
	proto.RegisterType((*PopFromPoolRes)(nil), "lsvoucherds.PopFromPoolRes")
}

func init() { proto.RegisterFile("grpc/proto/voucherds.proto", fileDescriptor_1e572ab4fba54a2c) }

var fileDescriptor_1e572ab4fba54a2c = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5f, 0x6f, 0xd3, 0x30,
	0x10, 0x6f, 0x9a, 0x76, 0x5d, 0x2f, 0xc0, 0x26, 0x33, 0x26, 0xcb, 0x4c, 0xa8, 0xf2, 0x5e, 0x2a,
	0x1e, 0x3a, 0xa9, 0x43, 0xe2, 0xcf, 0x1b, 0x12, 0x6c, 0xca, 0x03, 0xa8, 0x8a, 0x06, 0xaf, 0x90,
	0x35, 0x56, 0x09, 0x78, 0x75, 0x6a, 0xbb, 0x95, 0x2a, 0xf1, 0x45, 0xf8, 0x26, 0x7c, 0x04, 0x3e,
	0x16, 0xb2, 0x9d, 0xb4, 0x49, 0xc8, 0xc6, 0xa4, 0x69, 0x6f, 0xfe, 0xfd, 0xee, 0x77, 0x77, 0xbe,
	0xf3, 0x9d, 0x81, 0xcc, 0x64, 0x36, 0x3d, 0xc9, 0xa4, 0xd0, 0xe2, 0x64, 0x25, 0x96, 0xd3, 0x6f,
	0x4c, 0x26, 0x6a, 0x64, 0x31, 0x0a, 0xb8, 0xda, 0x50, 0xb4, 0x0f, 0xbd, 0x49, 0x3a, 0x9f, 0x45,
	0x6c, 0x41, 0x8f, 0x8b, 0xa3, 0x42, 0x18, 0x7a, 0x2b, 0x26, 0x55, 0x2a, 0xe6, 0xd8, 0x1b, 0x78,
	0xc3, 0x7e, 0x54, 0x40, 0xa3, 0xbf, 0x60, 0x4a, 0xe7, 0x7a, 0x77, 0xb4, 0xfa, 0x54, 0x7d, 0x8e,
	0x79, 0x9a, 0x58, 0xfd, 0x6e, 0x54, 0x40, 0xfa, 0xcb, 0x83, 0xbd, 0x4f, 0x19, 0x17, 0x71, 0x72,
	0x21, 0x26, 0x42, 0xf0, 0x88, 0x2d, 0xd0, 0x21, 0xec, 0x48, 0x36, 0xdb, 0x06, 0xcf, 0x11, 0x7a,
	0x06, 0x30, 0x5d, 0x2a, 0x2d, 0xae, 0x98, 0x0c, 0x13, 0xdc, 0xb6, 0xb6, 0x12, 0x63, 0xfc, 0x32,
	0x21, 0x78, 0x98, 0x60, 0xdf, 0xf9, 0x39, 0x84, 0x08, 0xec, 0x2e, 0x6d, 0x8a, 0x30, 0xc1, 0x1d,
	0x6b, 0xd9, 0x60, 0x63, 0xcb, 0x8b, 0x55, 0xb8, 0x3b, 0xf0, 0x8d, 0xad, 0xc0, 0xf4, 0xb4, 0x7e,
	0x35, 0x85, 0x06, 0x10, 0x68, 0xa1, 0x63, 0xee, 0x78, 0x7b, 0x3f, 0x3f, 0x2a, 0x53, 0xf4, 0x3b,
	0x1c, 0x9e, 0x33, 0x6d, 0xf4, 0x6f, 0x57, 0x71, 0xca, 0xe3, 0xcb, 0x94, 0xa7, 0x7a, 0x7d, 0x97,
	0xb2, 0x30, 0xf4, 0x5c, 0x21, 0x0a, 0xfb, 0xf6, 0x86, 0x05, 0xa4, 0x7f, 0xda, 0xd7, 0x24, 0x53,
	0x88, 0xc1, 0x7e, 0x9d, 0xc6, 0xde, 0xc0, 0x1f, 0x06, 0xe3, 0xd7, 0xa3, 0xd2, 0xfb, 0x8e, 0x9a,
	0xdd, 0x47, 0x75, 0xee, 0xfd, 0x5c, 0xcb, 0x75, 0xf4, 0x4f, 0x48, 0x12, 0xc2, 0xe3, 0x86, 0x08,
	0xe8, 0x00, 0xba, 0xb6, 0x27, 0x79, 0x83, 0x1c, 0x40, 0x47, 0xd0, 0x8f, 0x9d, 0x8a, 0x33, 0x5b,
	0xa7, 0x1f, 0x6d, 0x09, 0xf2, 0x13, 0x9e, 0x34, 0x66, 0x45, 0xfb, 0xe0, 0xff, 0x60, 0xeb, 0xbc,
	0x69, 0xe6, 0x88, 0x3e, 0x40, 0x77, 0x15, 0xf3, 0xa5, 0x0b, 0x12, 0x8c, 0x5f, 0xde, 0xa6, 0xa2,
	0x26, 0xda, 0x45, 0x79, 0xd3, 0x7e, 0xe5, 0xd1, 0x2f, 0xf0, 0xf0, 0x1d, 0xe3, 0x4c, 0xb3, 0x7b,
	0x1a, 0x42, 0xba, 0x57, 0x4d, 0xa0, 0xe8, 0x57, 0x78, 0x34, 0x11, 0xd9, 0x99, 0x14, 0x57, 0xf7,
	0x95, 0xf2, 0x79, 0x2d, 0x83, 0xdb, 0x5b, 0xd7, 0xa8, 0xcd, 0xde, 0x3a, 0x38, 0xfe, 0xed, 0x43,
	0xe7, 0x5c, 0x66, 0x53, 0xf4, 0x02, 0x3a, 0x66, 0xcb, 0xd1, 0x41, 0xa5, 0xa9, 0xf9, 0x1f, 0x40,
	0x9a, 0x58, 0x45, 0x5b, 0xc6, 0xcb, 0xec, 0x7a, 0xcd, 0x2b, 0xff, 0x09, 0x48, 0x13, 0x6b, 0xbc,
	0x3e, 0xc2, 0x83, 0xf2, 0x82, 0xa1, 0xa3, 0x8a, 0xae, 0xf6, 0x2d, 0x90, 0x9b, 0xac, 0x8a, 0xb6,
	0x86, 0x1e, 0x8a, 0x9b, 0xa7, 0xf1, 0xf8, 0xff, 0xf3, 0xb1, 0x20, 0xb7, 0x10, 0x99, 0x2b, 0x9f,
	0x01, 0x6c, 0x9f, 0x11, 0x91, 0x8a, 0x53, 0x65, 0x80, 0xc8, 0xf5, 0x36, 0x13, 0x27, 0x84, 0xa0,
	0xf4, 0x36, 0xe8, 0x69, 0xb5, 0xaf, 0x95, 0xb9, 0x20, 0x37, 0x18, 0x15, 0x6d, 0x5d, 0xee, 0xd8,
	0x6f, 0xfb, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x2c, 0xd6, 0x41, 0xd4, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcClient is the client API for Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
	Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRes, error)
	UploadToPool(ctx context.Context, opts ...grpc.CallOption) (Grpc_UploadToPoolClient, error)
	GetPoolAvailability(ctx context.Context, in *GetPoolAvailabilityReq, opts ...grpc.CallOption) (*GetPoolAvailabilityRes, error)
	DeletePool(ctx context.Context, in *DeletePoolReq, opts ...grpc.CallOption) (*DeletePoolRes, error)
	PopFromPool(ctx context.Context, in *PopFromPoolReq, opts ...grpc.CallOption) (*PopFromPoolRes, error)
}

type grpcClient struct {
	cc *grpc.ClientConn
}

func NewGrpcClient(cc *grpc.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := c.cc.Invoke(ctx, "/lsvoucherds.Grpc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRes, error) {
	out := new(TestRes)
	err := c.cc.Invoke(ctx, "/lsvoucherds.Grpc/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) UploadToPool(ctx context.Context, opts ...grpc.CallOption) (Grpc_UploadToPoolClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Grpc_serviceDesc.Streams[0], "/lsvoucherds.Grpc/UploadToPool", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcUploadToPoolClient{stream}
	return x, nil
}

type Grpc_UploadToPoolClient interface {
	Send(*UploadToPoolReq) error
	CloseAndRecv() (*UploadToPoolRes, error)
	grpc.ClientStream
}

type grpcUploadToPoolClient struct {
	grpc.ClientStream
}

func (x *grpcUploadToPoolClient) Send(m *UploadToPoolReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcUploadToPoolClient) CloseAndRecv() (*UploadToPoolRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadToPoolRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcClient) GetPoolAvailability(ctx context.Context, in *GetPoolAvailabilityReq, opts ...grpc.CallOption) (*GetPoolAvailabilityRes, error) {
	out := new(GetPoolAvailabilityRes)
	err := c.cc.Invoke(ctx, "/lsvoucherds.Grpc/GetPoolAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) DeletePool(ctx context.Context, in *DeletePoolReq, opts ...grpc.CallOption) (*DeletePoolRes, error) {
	out := new(DeletePoolRes)
	err := c.cc.Invoke(ctx, "/lsvoucherds.Grpc/DeletePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) PopFromPool(ctx context.Context, in *PopFromPoolReq, opts ...grpc.CallOption) (*PopFromPoolRes, error) {
	out := new(PopFromPoolRes)
	err := c.cc.Invoke(ctx, "/lsvoucherds.Grpc/PopFromPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
type GrpcServer interface {
	Ping(context.Context, *PingReq) (*PingRes, error)
	Test(context.Context, *TestReq) (*TestRes, error)
	UploadToPool(Grpc_UploadToPoolServer) error
	GetPoolAvailability(context.Context, *GetPoolAvailabilityReq) (*GetPoolAvailabilityRes, error)
	DeletePool(context.Context, *DeletePoolReq) (*DeletePoolRes, error)
	PopFromPool(context.Context, *PopFromPoolReq) (*PopFromPoolRes, error)
}

func RegisterGrpcServer(s *grpc.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsvoucherds.Grpc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsvoucherds.Grpc/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Test(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_UploadToPool_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcServer).UploadToPool(&grpcUploadToPoolServer{stream})
}

type Grpc_UploadToPoolServer interface {
	SendAndClose(*UploadToPoolRes) error
	Recv() (*UploadToPoolReq, error)
	grpc.ServerStream
}

type grpcUploadToPoolServer struct {
	grpc.ServerStream
}

func (x *grpcUploadToPoolServer) SendAndClose(m *UploadToPoolRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcUploadToPoolServer) Recv() (*UploadToPoolReq, error) {
	m := new(UploadToPoolReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Grpc_GetPoolAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolAvailabilityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetPoolAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsvoucherds.Grpc/GetPoolAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetPoolAvailability(ctx, req.(*GetPoolAvailabilityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_DeletePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePoolReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).DeletePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsvoucherds.Grpc/DeletePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).DeletePool(ctx, req.(*DeletePoolReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_PopFromPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopFromPoolReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).PopFromPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lsvoucherds.Grpc/PopFromPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).PopFromPool(ctx, req.(*PopFromPoolReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lsvoucherds.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Grpc_Ping_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _Grpc_Test_Handler,
		},
		{
			MethodName: "GetPoolAvailability",
			Handler:    _Grpc_GetPoolAvailability_Handler,
		},
		{
			MethodName: "DeletePool",
			Handler:    _Grpc_DeletePool_Handler,
		},
		{
			MethodName: "PopFromPool",
			Handler:    _Grpc_PopFromPool_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadToPool",
			Handler:       _Grpc_UploadToPool_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/proto/voucherds.proto",
}
