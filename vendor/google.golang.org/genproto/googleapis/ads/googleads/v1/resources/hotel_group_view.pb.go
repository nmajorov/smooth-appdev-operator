// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/hotel_group_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A hotel group view.
type HotelGroupView struct {
	// The resource name of the hotel group view.
	// Hotel Group view resource names have the form:
	//
	// `customers/{customer_id}/hotelGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelGroupView) Reset()         { *m = HotelGroupView{} }
func (m *HotelGroupView) String() string { return proto.CompactTextString(m) }
func (*HotelGroupView) ProtoMessage()    {}
func (*HotelGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_group_view_63dd209963fe113e, []int{0}
}
func (m *HotelGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelGroupView.Unmarshal(m, b)
}
func (m *HotelGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelGroupView.Marshal(b, m, deterministic)
}
func (dst *HotelGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelGroupView.Merge(dst, src)
}
func (m *HotelGroupView) XXX_Size() int {
	return xxx_messageInfo_HotelGroupView.Size(m)
}
func (m *HotelGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelGroupView proto.InternalMessageInfo

func (m *HotelGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelGroupView)(nil), "google.ads.googleads.v1.resources.HotelGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/hotel_group_view.proto", fileDescriptor_hotel_group_view_63dd209963fe113e)
}

var fileDescriptor_hotel_group_view_63dd209963fe113e = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x69, 0x05, 0xc1, 0xa2, 0x0e, 0xe7, 0x22, 0xe2, 0xe0, 0x29, 0x07, 0x4e, 0x09, 0x41,
	0x04, 0x89, 0x53, 0x6e, 0xa9, 0x38, 0xc8, 0x71, 0x43, 0x07, 0x29, 0x94, 0x78, 0x09, 0x31, 0xd0,
	0xe6, 0x2b, 0x49, 0xaf, 0xb7, 0xfa, 0x2c, 0x8e, 0x3e, 0x8a, 0x8f, 0xe2, 0x53, 0x48, 0x1a, 0x13,
	0x70, 0xd1, 0xed, 0x4f, 0xf2, 0xfb, 0xff, 0xbe, 0x8f, 0xaf, 0xb8, 0x53, 0x00, 0xaa, 0x95, 0x98,
	0x0b, 0x87, 0x43, 0xf4, 0x69, 0x24, 0xd8, 0x4a, 0x07, 0x5b, 0xbb, 0x91, 0x0e, 0xbf, 0xc2, 0x20,
	0xdb, 0x46, 0x59, 0xd8, 0xf6, 0xcd, 0xa8, 0xe5, 0x0e, 0xf5, 0x16, 0x06, 0x98, 0xcd, 0x03, 0x8e,
	0xb8, 0x70, 0x28, 0x35, 0xd1, 0x48, 0x50, 0x6a, 0x9e, 0x9d, 0x47, 0x79, 0xaf, 0x31, 0x37, 0x06,
	0x06, 0x3e, 0x68, 0x30, 0x2e, 0x08, 0x2e, 0x6f, 0x8b, 0xe3, 0x07, 0xaf, 0x2e, 0xbd, 0xb9, 0xd2,
	0x72, 0x37, 0xbb, 0x2a, 0x8e, 0x62, 0xb9, 0x31, 0xbc, 0x93, 0xa7, 0xd9, 0x45, 0x76, 0x7d, 0xb0,
	0x3e, 0x8c, 0x8f, 0x4f, 0xbc, 0x93, 0xcb, 0xb7, 0xbc, 0x58, 0x6c, 0xa0, 0x43, 0xff, 0x8e, 0x5f,
	0x9e, 0xfc, 0xd6, 0xaf, 0xfc, 0xd4, 0x55, 0xf6, 0xfc, 0xf8, 0xd3, 0x54, 0xd0, 0x72, 0xa3, 0x10,
	0x58, 0x85, 0x95, 0x34, 0xd3, 0x4e, 0xf1, 0x04, 0xbd, 0x76, 0x7f, 0x5c, 0xe4, 0x3e, 0xa5, 0xf7,
	0x7c, 0xaf, 0x64, 0xec, 0x23, 0x9f, 0x97, 0x41, 0xc9, 0x84, 0x43, 0x21, 0xfa, 0x54, 0x11, 0xb4,
	0x8e, 0xe4, 0x67, 0x64, 0x6a, 0x26, 0x5c, 0x9d, 0x98, 0xba, 0x22, 0x75, 0x62, 0xbe, 0xf2, 0x45,
	0xf8, 0xa0, 0x94, 0x09, 0x47, 0x69, 0xa2, 0x28, 0xad, 0x08, 0xa5, 0x89, 0x7b, 0xd9, 0x9f, 0x96,
	0xbd, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd7, 0x25, 0xfa, 0xbd, 0x01, 0x00, 0x00,
}
