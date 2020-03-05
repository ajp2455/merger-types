// Code generated by protoc-gen-go. DO NOT EDIT.
// source: merger.proto

package merger

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type Params struct {
	Conns                int64    `protobuf:"varint,1,opt,name=Conns,json=conns,proto3" json:"conns"`
	Timeout              int64    `protobuf:"varint,2,opt,name=Timeout,json=timeout,proto3" json:"timeout"`
	Endtime              int64    `protobuf:"varint,3,opt,name=Endtime,json=endtime,proto3" json:"endtime"`
	Hostname             string   `protobuf:"bytes,4,opt,name=Hostname,json=hostname,proto3" json:"hostname"`
	Origin               string   `protobuf:"bytes,5,opt,name=Origin,json=origin,proto3" json:"origin"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c1fd482da7fe1f7, []int{0}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetConns() int64 {
	if m != nil {
		return m.Conns
	}
	return 0
}

func (m *Params) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Params) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *Params) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Params) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type Report struct {
	Info    *Params   `protobuf:"bytes,1,opt,name=Info,json=info,proto3" json:"Info,omitempty"`
	Clients []*Client `protobuf:"bytes,2,rep,name=Clients,json=clients,proto3" json:"clients"`
	// Points     [][3]int
	Points               []*Tuple          `protobuf:"bytes,3,rep,name=Points,json=points,proto3" json:"points"`
	Errors               []int64           `protobuf:"varint,4,rep,packed,name=Errors,json=errors,proto3" json:"errors"`
	Timeouts             []int64           `protobuf:"varint,5,rep,packed,name=Timeouts,json=timeouts,proto3" json:"timeouts"`
	ScriptData           map[string]string `protobuf:"bytes,6,rep,name=ScriptData,proto3" json:"ScriptData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c1fd482da7fe1f7, []int{1}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetInfo() *Params {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Report) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *Report) GetPoints() []*Tuple {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Report) GetErrors() []int64 {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Report) GetTimeouts() []int64 {
	if m != nil {
		return m.Timeouts
	}
	return nil
}

func (m *Report) GetScriptData() map[string]string {
	if m != nil {
		return m.ScriptData
	}
	return nil
}

type Client struct {
	Id                   int64     `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"id"`
	Origin               string    `protobuf:"bytes,2,opt,name=Origin,json=origin,proto3" json:"origin"`
	Windows              []*Window `protobuf:"bytes,3,rep,name=Windows,json=windows,proto3" json:"windows"`
	Errors               []int64   `protobuf:"varint,4,rep,packed,name=Errors,json=errors,proto3" json:"errors"`
	Established          []int64   `protobuf:"varint,5,rep,packed,name=Established,json=established,proto3" json:"established"`
	Responses            int64     `protobuf:"varint,6,opt,name=Responses,json=responses,proto3" json:"responses"`
	Timeouts             int64     `protobuf:"varint,7,opt,name=Timeouts,json=timeouts,proto3" json:"timeouts"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c1fd482da7fe1f7, []int{2}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Client) GetWindows() []*Window {
	if m != nil {
		return m.Windows
	}
	return nil
}

func (m *Client) GetErrors() []int64 {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Client) GetEstablished() []int64 {
	if m != nil {
		return m.Established
	}
	return nil
}

func (m *Client) GetResponses() int64 {
	if m != nil {
		return m.Responses
	}
	return 0
}

func (m *Client) GetTimeouts() int64 {
	if m != nil {
		return m.Timeouts
	}
	return 0
}

type Window struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"value"`
	Start                int64    `protobuf:"varint,2,opt,name=Start,json=start,proto3" json:"start"`
	End                  int64    `protobuf:"varint,3,opt,name=End,json=end,proto3" json:"end"`
	Iter                 int64    `protobuf:"varint,4,opt,name=Iter,json=iter,proto3" json:"iter"`
	Count                int64    `protobuf:"varint,5,opt,name=Count,json=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Window) Reset()         { *m = Window{} }
func (m *Window) String() string { return proto.CompactTextString(m) }
func (*Window) ProtoMessage()    {}
func (*Window) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c1fd482da7fe1f7, []int{3}
}

func (m *Window) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Window.Unmarshal(m, b)
}
func (m *Window) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Window.Marshal(b, m, deterministic)
}
func (m *Window) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Window.Merge(m, src)
}
func (m *Window) XXX_Size() int {
	return xxx_messageInfo_Window.Size(m)
}
func (m *Window) XXX_DiscardUnknown() {
	xxx_messageInfo_Window.DiscardUnknown(m)
}

var xxx_messageInfo_Window proto.InternalMessageInfo

func (m *Window) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Window) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Window) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Window) GetIter() int64 {
	if m != nil {
		return m.Iter
	}
	return 0
}

func (m *Window) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Tuple struct {
	Data                 []int64  `protobuf:"varint,1,rep,packed,name=Data,json=data,proto3" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tuple) Reset()         { *m = Tuple{} }
func (m *Tuple) String() string { return proto.CompactTextString(m) }
func (*Tuple) ProtoMessage()    {}
func (*Tuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c1fd482da7fe1f7, []int{4}
}

func (m *Tuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tuple.Unmarshal(m, b)
}
func (m *Tuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tuple.Marshal(b, m, deterministic)
}
func (m *Tuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tuple.Merge(m, src)
}
func (m *Tuple) XXX_Size() int {
	return xxx_messageInfo_Tuple.Size(m)
}
func (m *Tuple) XXX_DiscardUnknown() {
	xxx_messageInfo_Tuple.DiscardUnknown(m)
}

var xxx_messageInfo_Tuple proto.InternalMessageInfo

func (m *Tuple) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "merger.Params")
	proto.RegisterType((*Report)(nil), "merger.Report")
	proto.RegisterMapType((map[string]string)(nil), "merger.Report.ScriptDataEntry")
	proto.RegisterType((*Client)(nil), "merger.Client")
	proto.RegisterType((*Window)(nil), "merger.Window")
	proto.RegisterType((*Tuple)(nil), "merger.Tuple")
}

func init() {
	proto.RegisterFile("merger.proto", fileDescriptor_4c1fd482da7fe1f7)
}

var fileDescriptor_4c1fd482da7fe1f7 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x4e, 0xdc, 0x3e,
	0x10, 0xc7, 0x95, 0x38, 0x71, 0xd8, 0x59, 0xfe, 0xfd, 0x0c, 0x3f, 0x35, 0xe2, 0x80, 0xb6, 0x56,
	0xa5, 0xa2, 0x4a, 0xd0, 0xaa, 0x5c, 0x0a, 0x12, 0x3d, 0x40, 0x57, 0x2a, 0xa7, 0x22, 0x83, 0xda,
	0x73, 0x20, 0x66, 0x49, 0xbb, 0x6b, 0xaf, 0x6c, 0x6f, 0x11, 0xe7, 0xf6, 0x29, 0x7a, 0xe8, 0x13,
	0xf5, 0x25, 0xfa, 0x00, 0x7d, 0x87, 0xca, 0x63, 0x07, 0xb2, 0x70, 0x40, 0x3d, 0xcd, 0xd8, 0xf3,
	0x99, 0xc8, 0xfe, 0x7e, 0x33, 0x86, 0xc5, 0x89, 0x34, 0x23, 0x69, 0x76, 0xa6, 0x46, 0x3b, 0xcd,
	0x68, 0x58, 0x6d, 0xac, 0xb9, 0x6a, 0x34, 0x92, 0xe6, 0x65, 0x08, 0xa1, 0xc8, 0xff, 0x24, 0x40,
	0x4f, 0x2a, 0x53, 0x4d, 0x2c, 0x7b, 0x0e, 0xf9, 0x91, 0x56, 0xca, 0x96, 0xc9, 0x20, 0xd9, 0x22,
	0x87, 0xff, 0xfd, 0xf8, 0xfe, 0x93, 0x2c, 0x7e, 0xb6, 0x5a, 0xed, 0xf3, 0x0b, 0xbf, 0xcf, 0x45,
	0x8e, 0x91, 0x6d, 0x43, 0x71, 0xd6, 0x4c, 0xa4, 0x9e, 0xb9, 0x32, 0x45, 0x74, 0xcd, 0xa3, 0xcb,
	0x01, 0x75, 0xa1, 0xc2, 0x45, 0x11, 0x33, 0x8f, 0x0f, 0x55, 0xed, 0x57, 0x25, 0x79, 0x80, 0xcb,
	0x50, 0xe1, 0xa2, 0x88, 0x19, 0x7b, 0x05, 0x0b, 0xef, 0xb5, 0x75, 0xaa, 0x9a, 0xc8, 0x32, 0x1b,
	0x24, 0x5b, 0xbd, 0xc3, 0x75, 0xcf, 0xaf, 0x04, 0xfe, 0x2a, 0x96, 0xb8, 0x58, 0x68, 0x53, 0xf6,
	0x02, 0xe8, 0x07, 0xd3, 0x8c, 0x1a, 0x55, 0xe6, 0xc8, 0x33, 0xcf, 0x2f, 0x05, 0x5e, 0x63, 0x81,
	0x0b, 0x1a, 0x93, 0x6f, 0x04, 0xa8, 0x90, 0x53, 0x6d, 0x1c, 0xe3, 0x90, 0x1d, 0xab, 0x4b, 0x8d,
	0xd7, 0xed, 0xbf, 0x5e, 0xde, 0x89, 0xa2, 0x05, 0x35, 0x44, 0xd6, 0xa8, 0x4b, 0xcd, 0x0e, 0xa0,
	0x38, 0x1a, 0x37, 0x52, 0x39, 0x5b, 0xa6, 0x03, 0xd2, 0xc5, 0xc2, 0xf6, 0xdc, 0x5d, 0x2e, 0x02,
	0xc9, 0x45, 0x11, 0x33, 0xb6, 0x07, 0xf4, 0x44, 0x37, 0xbe, 0x9b, 0x60, 0xf7, 0x52, 0xdb, 0x7d,
	0x36, 0x9b, 0x8e, 0xe5, 0xdc, 0x41, 0xa7, 0xc8, 0x71, 0x41, 0x43, 0xe2, 0x2f, 0x35, 0x34, 0x46,
	0x1b, 0x5b, 0x66, 0x03, 0xb2, 0x45, 0xe6, 0x58, 0x89, 0x05, 0x2e, 0x68, 0x48, 0xbc, 0x64, 0xd1,
	0x10, 0x5b, 0xe6, 0x48, 0x77, 0x25, 0x8b, 0x3e, 0x58, 0x2e, 0x16, 0xda, 0x94, 0xbd, 0x05, 0x38,
	0xbd, 0x30, 0xcd, 0xd4, 0xbd, 0xab, 0x5c, 0x55, 0x52, 0x3c, 0xdc, 0x66, 0x7b, 0xb8, 0xa0, 0xcf,
	0xce, 0x1d, 0x30, 0x54, 0xce, 0xdc, 0x88, 0x4e, 0xc7, 0xc6, 0x01, 0xac, 0xdc, 0x2b, 0xb3, 0x55,
	0x20, 0x5f, 0xe4, 0x0d, 0xaa, 0xd9, 0x13, 0x3e, 0x65, 0xeb, 0x90, 0x7f, 0xad, 0xc6, 0x33, 0x89,
	0x7f, 0x49, 0x4f, 0x84, 0xc5, 0x7e, 0xfa, 0x26, 0xe1, 0xbf, 0x53, 0xa0, 0x41, 0x40, 0xb6, 0x09,
	0xe9, 0x71, 0x1d, 0x7f, 0xb9, 0x65, 0x7f, 0xea, 0x5e, 0x38, 0x75, 0x53, 0x73, 0x91, 0x36, 0x75,
	0xc7, 0xdc, 0xf4, 0x31, 0x73, 0xbd, 0x5b, 0x9f, 0x1a, 0x55, 0xeb, 0xeb, 0x56, 0xef, 0x5b, 0xb7,
	0xc2, 0xf6, 0x9c, 0x5b, 0xd7, 0x81, 0xe4, 0xa2, 0x88, 0xd9, 0x3f, 0x49, 0xbe, 0x07, 0xfd, 0xa1,
	0x75, 0xd5, 0xf9, 0xb8, 0xb1, 0x57, 0xb2, 0x8e, 0xaa, 0x3f, 0xf1, 0x0d, 0x2c, 0x36, 0xdc, 0x55,
	0xb9, 0xe8, 0x77, 0x56, 0x6c, 0x17, 0x7a, 0x42, 0xda, 0xa9, 0x56, 0x56, 0xda, 0x92, 0xe2, 0xc5,
	0xff, 0xf7, 0x8d, 0xab, 0xa1, 0xd1, 0xb4, 0x35, 0x2e, 0x7a, 0xb7, 0xf9, 0x9c, 0xc5, 0x05, 0xf6,
	0x3c, 0x62, 0x31, 0xff, 0x95, 0x00, 0x0d, 0xd7, 0xf6, 0x93, 0xfd, 0x11, 0x8d, 0x78, 0x38, 0xd9,
	0xe8, 0x09, 0x8f, 0xde, 0x78, 0xf0, 0xd4, 0x55, 0xa6, 0x9d, 0xeb, 0x2e, 0x68, 0xfd, 0x3e, 0x17,
	0x39, 0x46, 0xf6, 0x14, 0xc8, 0x50, 0xd5, 0x71, 0x9e, 0x57, 0x3c, 0x06, 0xb7, 0xf3, 0xcc, 0x05,
	0x91, 0xaa, 0x66, 0xcf, 0x20, 0x3b, 0x76, 0xd2, 0xe0, 0x0c, 0x93, 0xc3, 0x55, 0xcf, 0xf4, 0xa3,
	0xb5, 0x4e, 0x1a, 0x2e, 0x32, 0x1f, 0xc2, 0xa3, 0x33, 0x53, 0x0e, 0x47, 0xf7, 0xfe, 0xa3, 0x33,
	0x53, 0x0e, 0x1f, 0x1d, 0x1f, 0xb7, 0x21, 0xc7, 0xa1, 0xf1, 0xdf, 0xc5, 0x9f, 0x36, 0x41, 0xc9,
	0xbb, 0xdf, 0xad, 0x2b, 0x57, 0x71, 0x91, 0xf9, 0x70, 0x4e, 0xf1, 0x79, 0xdb, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x51, 0x46, 0xe8, 0x0a, 0x0b, 0x05, 0x00, 0x00,
}
