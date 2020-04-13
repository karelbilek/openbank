// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/fx/all.proto

package fx

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// FX holds all the details about a foreign exchange object.
type FX struct {
	// BankID is an identifier for the bank on this fx transaction.
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,2,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"from_currency_code,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode string `protobuf:"bytes,3,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"to_currency_code,omitempty"`
	// Rate is the exchange rate of the foreign exchange.
	Rate string `protobuf:"bytes,4,opt,name=Rate,json=rate,proto3" json:"rate,omitempty"`
	// InverseRate is the inverse of the exchange rate of the foreign exchange.
	InverseRate string `protobuf:"bytes,5,opt,name=InverseRate,json=inverse_rate,proto3" json:"inverse_rate,omitempty"`
	// EffectiveDate is the effective date of the foreign exchange quote.
	EffectiveDate        string   `protobuf:"bytes,6,opt,name=EffectiveDate,json=effective_date,proto3" json:"effective_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FX) Reset()         { *m = FX{} }
func (m *FX) String() string { return proto.CompactTextString(m) }
func (*FX) ProtoMessage()    {}
func (*FX) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{0}
}

func (m *FX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FX.Unmarshal(m, b)
}
func (m *FX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FX.Marshal(b, m, deterministic)
}
func (m *FX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FX.Merge(m, src)
}
func (m *FX) XXX_Size() int {
	return xxx_messageInfo_FX.Size(m)
}
func (m *FX) XXX_DiscardUnknown() {
	xxx_messageInfo_FX.DiscardUnknown(m)
}

var xxx_messageInfo_FX proto.InternalMessageInfo

func (m *FX) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *FX) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *FX) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func (m *FX) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *FX) GetInverseRate() string {
	if m != nil {
		return m.InverseRate
	}
	return ""
}

func (m *FX) GetEffectiveDate() string {
	if m != nil {
		return m.EffectiveDate
	}
	return ""
}

// CreateFXRequest is a request envelope to create a foreign exchange quote.
type CreateFXRequest struct {
	// FX is the foreign exchange information to create.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"fx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFXRequest) Reset()         { *m = CreateFXRequest{} }
func (m *CreateFXRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFXRequest) ProtoMessage()    {}
func (*CreateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{1}
}

func (m *CreateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFXRequest.Unmarshal(m, b)
}
func (m *CreateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFXRequest.Marshal(b, m, deterministic)
}
func (m *CreateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFXRequest.Merge(m, src)
}
func (m *CreateFXRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFXRequest.Size(m)
}
func (m *CreateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFXRequest proto.InternalMessageInfo

func (m *CreateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// UpdateFXRequest is a request envelope to update a foreign exchange quote.
type UpdateFXRequest struct {
	// FX is the foreign exchange information to update.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"fx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFXRequest) Reset()         { *m = UpdateFXRequest{} }
func (m *UpdateFXRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFXRequest) ProtoMessage()    {}
func (*UpdateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{2}
}

func (m *UpdateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFXRequest.Unmarshal(m, b)
}
func (m *UpdateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFXRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFXRequest.Merge(m, src)
}
func (m *UpdateFXRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFXRequest.Size(m)
}
func (m *UpdateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFXRequest proto.InternalMessageInfo

func (m *UpdateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// GetCurrentFXRateRequest is a request envelope to get a foreign exchange rate.
type GetCurrentFXRateRequest struct {
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,1,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"from_currency_code,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode       string   `protobuf:"bytes,2,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"to_currency_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentFXRateRequest) Reset()         { *m = GetCurrentFXRateRequest{} }
func (m *GetCurrentFXRateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentFXRateRequest) ProtoMessage()    {}
func (*GetCurrentFXRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{3}
}

func (m *GetCurrentFXRateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentFXRateRequest.Unmarshal(m, b)
}
func (m *GetCurrentFXRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentFXRateRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentFXRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentFXRateRequest.Merge(m, src)
}
func (m *GetCurrentFXRateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentFXRateRequest.Size(m)
}
func (m *GetCurrentFXRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentFXRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentFXRateRequest proto.InternalMessageInfo

func (m *GetCurrentFXRateRequest) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *GetCurrentFXRateRequest) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func init() {
	proto.RegisterType((*FX)(nil), "fx.FX")
	proto.RegisterType((*CreateFXRequest)(nil), "fx.CreateFXRequest")
	proto.RegisterType((*UpdateFXRequest)(nil), "fx.UpdateFXRequest")
	proto.RegisterType((*GetCurrentFXRateRequest)(nil), "fx.GetCurrentFXRateRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/fx/all.proto", fileDescriptor_e3bdb6bf3dbfaace)
}

var fileDescriptor_e3bdb6bf3dbfaace = []byte{
	// 1174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xde, 0x5d, 0xa7, 0x6e, 0x3a, 0xfd, 0x20, 0x1a, 0x50, 0x89, 0xdc, 0x52, 0x46, 0xa9, 0x44,
	0x43, 0x68, 0x66, 0x1d, 0x37, 0x95, 0xaa, 0x48, 0x55, 0x71, 0xd2, 0xb8, 0x24, 0x02, 0x11, 0x99,
	0x52, 0x59, 0x95, 0x50, 0x34, 0xde, 0x7d, 0xd7, 0x1e, 0x62, 0xcf, 0x6c, 0x67, 0x66, 0x6d, 0x87,
	0x52, 0x89, 0x2f, 0x89, 0x72, 0x2b, 0xe1, 0xc4, 0x19, 0x7e, 0x00, 0x37, 0xb8, 0x71, 0xe2, 0x5e,
	0x89, 0x0b, 0x3f, 0x80, 0x1b, 0x77, 0x84, 0xc4, 0x05, 0xcd, 0xee, 0xda, 0x4d, 0xec, 0xa4, 0x40,
	0x4e, 0x1e, 0x7b, 0x9e, 0xe7, 0x99, 0xf7, 0x7d, 0xde, 0x0f, 0xa3, 0xab, 0x2d, 0x6e, 0xda, 0x49,
	0x93, 0x06, 0xb2, 0xeb, 0xcb, 0x18, 0x44, 0x93, 0x89, 0x9d, 0x67, 0x87, 0xde, 0x92, 0x1f, 0x0d,
	0x7c, 0xd6, 0xe9, 0xd0, 0x58, 0x49, 0x23, 0xb1, 0x17, 0x0d, 0x4a, 0x17, 0x5b, 0x52, 0xb6, 0x3a,
	0xe0, 0xb3, 0x98, 0xfb, 0x4c, 0x08, 0x69, 0x98, 0xe1, 0x52, 0xe8, 0x0c, 0x51, 0xba, 0x9a, 0x7e,
	0x04, 0x8b, 0x2d, 0x10, 0x8b, 0xba, 0xcf, 0x5a, 0x2d, 0x50, 0xbe, 0x8c, 0x53, 0xc4, 0x21, 0xe8,
	0x0b, 0xb9, 0x56, 0xfa, 0xad, 0x99, 0x44, 0x3e, 0x74, 0x63, 0xb3, 0x9b, 0x5d, 0xce, 0x7d, 0xe7,
	0x21, 0xaf, 0xd6, 0xc0, 0x97, 0x51, 0x71, 0x95, 0x89, 0x9d, 0x8d, 0xdb, 0xb3, 0x2e, 0x71, 0xe7,
	0x4f, 0xad, 0xa2, 0x69, 0x67, 0xd6, 0x99, 0x77, 0xca, 0xce, 0x96, 0x53, 0x3f, 0x69, 0xa3, 0xdc,
	0xe6, 0x21, 0xbe, 0x81, 0x66, 0x6a, 0x4a, 0x76, 0xd7, 0x12, 0xa5, 0x40, 0x04, 0xbb, 0x6b, 0x32,
	0x84, 0x59, 0x6f, 0x02, 0x8e, 0x23, 0x25, 0xbb, 0xdb, 0x41, 0x0e, 0xda, 0x0e, 0x64, 0x08, 0x78,
	0x19, 0x9d, 0xbb, 0x2b, 0x0f, 0xf0, 0x0a, 0x13, 0xbc, 0x19, 0x23, 0xc7, 0x58, 0x97, 0xd0, 0x54,
	0x9d, 0x19, 0x98, 0x9d, 0x9a, 0xc0, 0x4e, 0x29, 0x66, 0x00, 0x2f, 0xa2, 0xd3, 0x1b, 0xa2, 0x07,
	0x4a, 0x43, 0x0a, 0x3b, 0x31, 0x01, 0x3b, 0xc3, 0xb3, 0xeb, 0xed, 0x14, 0xbe, 0x84, 0xce, 0xae,
	0x47, 0x11, 0x04, 0x86, 0xf7, 0xe0, 0xb6, 0x25, 0x14, 0x27, 0x08, 0xe7, 0x60, 0x08, 0xd8, 0x0e,
	0x99, 0x81, 0x95, 0xe2, 0xb4, 0x33, 0xe3, 0xcc, 0x3a, 0x73, 0x37, 0xd1, 0x0b, 0x6b, 0x0a, 0x98,
	0x81, 0x5a, 0xa3, 0x0e, 0x0f, 0x12, 0xd0, 0x06, 0xcf, 0x59, 0xdf, 0x52, 0xb7, 0x4e, 0x57, 0x8a,
	0x34, 0x1a, 0xd0, 0x5a, 0xe3, 0x80, 0x94, 0x17, 0x0d, 0xf6, 0xd3, 0xdf, 0x8f, 0xc3, 0x63, 0xd3,
	0xbf, 0x76, 0xd1, 0xcb, 0x77, 0xc0, 0x64, 0xfe, 0x99, 0x5a, 0xc3, 0x66, 0x3b, 0xd4, 0x39, 0xac,
	0x26, 0xee, 0x31, 0x6b, 0xe2, 0xfd, 0x7b, 0x4d, 0x86, 0x31, 0x55, 0x7e, 0x38, 0x89, 0x4e, 0xd5,
	0x1a, 0xef, 0x81, 0xea, 0xf1, 0x00, 0xf0, 0xe7, 0x05, 0x34, 0x33, 0x1e, 0x21, 0xbe, 0x60, 0xd3,
	0x3a, 0x22, 0xee, 0x52, 0x9e, 0xf3, 0xdc, 0x2f, 0xde, 0x5e, 0xf5, 0xab, 0xac, 0x07, 0x17, 0xeb,
	0x60, 0x14, 0x87, 0x1e, 0x10, 0x18, 0x04, 0x6d, 0x26, 0x5a, 0x40, 0x6c, 0xdd, 0x48, 0x13, 0x4c,
	0x1f, 0x40, 0x10, 0xd3, 0x97, 0x24, 0x0f, 0x86, 0x83, 0x2e, 0x5d, 0x1f, 0xc1, 0x4d, 0x7b, 0x9c,
	0x62, 0xd3, 0x25, 0x6c, 0x88, 0xde, 0x25, 0x46, 0x12, 0x26, 0xa4, 0x69, 0x83, 0xda, 0xbc, 0x85,
	0x0a, 0x95, 0x72, 0x19, 0xdf, 0x40, 0x97, 0xf2, 0x50, 0x08, 0x0c, 0x20, 0x48, 0x0c, 0x84, 0x44,
	0x27, 0x41, 0x00, 0x5a, 0x47, 0x49, 0xa7, 0xb3, 0x4b, 0xf1, 0x79, 0xf4, 0x52, 0x09, 0x5f, 0xf6,
	0x43, 0x88, 0xb8, 0xe0, 0xd9, 0x6c, 0x45, 0x83, 0x5a, 0x63, 0x73, 0x09, 0x15, 0x96, 0xcb, 0xcb,
	0x78, 0x01, 0xcd, 0xd7, 0xc1, 0x24, 0x4a, 0x40, 0x48, 0xfa, 0x6d, 0x1b, 0x5e, 0x1b, 0x88, 0x02,
	0x2d, 0x13, 0x15, 0x00, 0xe1, 0x9a, 0x08, 0x69, 0x48, 0x24, 0x13, 0x11, 0xd2, 0xe6, 0xeb, 0xe8,
	0x0a, 0x2a, 0xbe, 0x5b, 0x4d, 0x4c, 0xbb, 0x82, 0x5f, 0x41, 0x17, 0xda, 0xc6, 0xc4, 0x7a, 0xc5,
	0xf7, 0x59, 0x62, 0xda, 0xb4, 0x29, 0x76, 0xa8, 0x91, 0x7e, 0x34, 0xa0, 0x0a, 0x58, 0xf8, 0xd9,
	0xaf, 0xbf, 0x7f, 0xe3, 0x5d, 0xc5, 0x0b, 0xf9, 0x52, 0x78, 0x38, 0x5e, 0xdb, 0x47, 0xfe, 0xc3,
	0x83, 0x45, 0x7b, 0xf4, 0xd8, 0x73, 0x9e, 0x78, 0x69, 0xbd, 0xf0, 0xb7, 0x1e, 0x9a, 0x1e, 0xb6,
	0x29, 0x7e, 0xd1, 0x1a, 0x3c, 0xd6, 0xb4, 0x23, 0xd7, 0xff, 0x70, 0xf7, 0xaa, 0x4f, 0xdd, 0xd4,
	0xf5, 0x57, 0x33, 0x0c, 0x61, 0x24, 0x92, 0x0a, 0x78, 0x4b, 0x3c, 0xb3, 0xf2, 0x41, 0x22, 0x0d,
	0x94, 0x96, 0x33, 0x80, 0x26, 0x8c, 0x08, 0xe8, 0x1f, 0x81, 0x22, 0x4c, 0x84, 0x44, 0xa5, 0x7e,
	0x68, 0xc2, 0x0d, 0xdd, 0xbc, 0x63, 0x6d, 0x5e, 0xc2, 0x6f, 0xa2, 0xd7, 0x6a, 0x39, 0x61, 0x7d,
	0x48, 0x08, 0x52, 0xbd, 0xff, 0x68, 0x77, 0x73, 0x01, 0xcd, 0x8f, 0xbc, 0xbb, 0x84, 0x2e, 0x1e,
	0xe1, 0x5d, 0x5f, 0x71, 0x03, 0xa9, 0x79, 0xa7, 0x57, 0xdc, 0x85, 0xb9, 0x62, 0xe6, 0xdf, 0x3e,
	0x6f, 0xfe, 0x76, 0xd1, 0xf4, 0x70, 0x06, 0x33, 0x6f, 0xc6, 0x26, 0xb2, 0x74, 0x9e, 0x66, 0x7b,
	0x92, 0x0e, 0xf7, 0x24, 0x5d, 0xb7, 0x7b, 0x72, 0xee, 0x67, 0x77, 0xaf, 0xfa, 0x7d, 0xee, 0x55,
	0xc6, 0x39, 0xda, 0x2b, 0x92, 0x01, 0xf4, 0x91, 0x88, 0x4d, 0xdf, 0xfa, 0xb2, 0x8c, 0xe7, 0x0f,
	0xf1, 0x65, 0xbf, 0x1f, 0x24, 0x49, 0x85, 0x42, 0x7a, 0xac, 0xfc, 0x4b, 0x63, 0xf9, 0x97, 0x0a,
	0x8f, 0x3d, 0x67, 0xf5, 0xc7, 0xe2, 0x5e, 0xf5, 0xcb, 0x22, 0x2a, 0x54, 0x68, 0x19, 0xdf, 0x43,
	0xc5, 0x5a, 0x83, 0x54, 0xb7, 0x36, 0xf0, 0xfa, 0x96, 0x92, 0x3d, 0x1e, 0x82, 0xce, 0xeb, 0x92,
	0x57, 0x92, 0x85, 0x44, 0xc6, 0xa0, 0xb2, 0x3f, 0x12, 0x22, 0xb3, 0xf6, 0x9e, 0xc8, 0x69, 0xd8,
	0xef, 0xb4, 0x72, 0x62, 0x89, 0x96, 0x69, 0x79, 0xc1, 0xf5, 0x2a, 0x33, 0x2c, 0x8e, 0x3b, 0x3c,
	0x48, 0x99, 0xfe, 0x87, 0x5a, 0x8a, 0x95, 0x89, 0x5f, 0xea, 0xdb, 0x76, 0x86, 0xca, 0xb8, 0x81,
	0xee, 0x1d, 0x36, 0x43, 0xd9, 0x58, 0x36, 0x65, 0xb8, 0x6b, 0xe7, 0xa8, 0xcb, 0x3a, 0x91, 0x54,
	0x5d, 0x66, 0x6c, 0xcf, 0x48, 0x45, 0x42, 0x09, 0xd9, 0x70, 0x75, 0x99, 0x09, 0xda, 0xf9, 0xf0,
	0xc7, 0x10, 0xd8, 0xeb, 0x9c, 0x4b, 0xeb, 0x6f, 0xdb, 0x07, 0x96, 0xf0, 0x3a, 0x5a, 0x3b, 0xfa,
	0x81, 0x91, 0x50, 0x20, 0x85, 0x61, 0x5c, 0xe8, 0xf4, 0x36, 0xd1, 0xa0, 0xae, 0xa4, 0x66, 0x84,
	0x20, 0x0c, 0x67, 0x1d, 0x4d, 0xeb, 0x5b, 0x56, 0xed, 0x1a, 0xde, 0x40, 0x77, 0x26, 0xd5, 0x2c,
	0xfe, 0x99, 0x54, 0x9b, 0xf5, 0x80, 0xc4, 0xa0, 0xba, 0x5c, 0x6b, 0x6e, 0x5d, 0x93, 0x84, 0xa5,
	0x55, 0x3d, 0xb0, 0x1e, 0x68, 0xfd, 0xff, 0x2f, 0x91, 0x7a, 0x0d, 0x15, 0xae, 0x97, 0xcb, 0xf8,
	0x16, 0xba, 0x79, 0x90, 0xc2, 0x04, 0x49, 0xc4, 0xc8, 0x01, 0x50, 0x4a, 0x2a, 0x22, 0x83, 0x20,
	0x51, 0xd6, 0xae, 0x4c, 0x51, 0x83, 0xea, 0x81, 0x22, 0x9a, 0x87, 0x40, 0xef, 0x7f, 0xe1, 0xa1,
	0x4f, 0xbd, 0x51, 0x4b, 0xfd, 0xe9, 0x4e, 0x17, 0xf0, 0xc7, 0xd5, 0x3c, 0x46, 0x49, 0xa2, 0xc1,
	0xe8, 0x7d, 0x6d, 0x03, 0x50, 0xa0, 0x8d, 0xe2, 0xa9, 0xb4, 0xcd, 0x25, 0x31, 0x6d, 0xeb, 0x4a,
	0x90, 0x0e, 0xb0, 0x4d, 0x5d, 0x53, 0x72, 0xb7, 0x0d, 0xfb, 0x2f, 0x6c, 0xda, 0xb1, 0x92, 0xa9,
	0x60, 0x24, 0x3b, 0x1d, 0xd9, 0xcf, 0x92, 0xb7, 0xef, 0x49, 0xc5, 0x3f, 0xca, 0x10, 0x76, 0x77,
	0x91, 0xa8, 0x23, 0xfb, 0x74, 0x7e, 0xaa, 0x32, 0x6d, 0xdb, 0xd5, 0x4a, 0xac, 0x9c, 0xb2, 0x27,
	0x23, 0x77, 0x40, 0xac, 0x7e, 0x80, 0xde, 0x78, 0xee, 0xaa, 0xc4, 0x67, 0xee, 0x71, 0xbb, 0x8f,
	0x06, 0x24, 0x64, 0x86, 0x21, 0xfa, 0xfc, 0xd9, 0xc0, 0xe7, 0xde, 0x61, 0x82, 0xb5, 0x60, 0x84,
	0x3f, 0x6b, 0x64, 0x8a, 0x61, 0x31, 0xa7, 0xd1, 0xe0, 0x2d, 0x77, 0xcb, 0xbd, 0xef, 0x45, 0x83,
	0x4f, 0x5c, 0xe7, 0xb1, 0xeb, 0x3c, 0x71, 0x9d, 0x9f, 0x5c, 0xe7, 0x37, 0xd7, 0xf9, 0xcb, 0x75,
	0x9e, 0x7a, 0x4e, 0xb3, 0x98, 0xae, 0x84, 0x6b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x62,
	0x7f, 0x43, 0xc8, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FXServiceClient is the client API for FXService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FXServiceClient interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	// specified currency code.
	GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type fXServiceClient struct {
	cc *grpc.ClientConn
}

func NewFXServiceClient(cc *grpc.ClientConn) FXServiceClient {
	return &fXServiceClient{cc}
}

func (c *fXServiceClient) GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/GetCurrentFXRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/CreateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/fx.FXService/UpdateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FXServiceServer is the server API for FXService service.
type FXServiceServer interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	// specified currency code.
	GetCurrentFXRate(context.Context, *GetCurrentFXRateRequest) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(context.Context, *CreateFXRequest) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(context.Context, *UpdateFXRequest) (*empty.Empty, error)
}

// UnimplementedFXServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFXServiceServer struct {
}

func (*UnimplementedFXServiceServer) GetCurrentFXRate(ctx context.Context, req *GetCurrentFXRateRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentFXRate not implemented")
}
func (*UnimplementedFXServiceServer) CreateFX(ctx context.Context, req *CreateFXRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFX not implemented")
}
func (*UnimplementedFXServiceServer) UpdateFX(ctx context.Context, req *UpdateFXRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFX not implemented")
}

func RegisterFXServiceServer(s *grpc.Server, srv FXServiceServer) {
	s.RegisterService(&_FXService_serviceDesc, srv)
}

func _FXService_GetCurrentFXRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentFXRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/GetCurrentFXRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, req.(*GetCurrentFXRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_CreateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).CreateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/CreateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).CreateFX(ctx, req.(*CreateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_UpdateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).UpdateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/UpdateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).UpdateFX(ctx, req.(*UpdateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FXService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fx.FXService",
	HandlerType: (*FXServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentFXRate",
			Handler:    _FXService_GetCurrentFXRate_Handler,
		},
		{
			MethodName: "CreateFX",
			Handler:    _FXService_CreateFX_Handler,
		},
		{
			MethodName: "UpdateFX",
			Handler:    _FXService_UpdateFX_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/fx/all.proto",
}
