// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: pb/faculty.proto

package grabber

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetFacultyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFacultyRequest) Reset() {
	*x = GetFacultyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFacultyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFacultyRequest) ProtoMessage() {}

func (x *GetFacultyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFacultyRequest.ProtoReflect.Descriptor instead.
func (*GetFacultyRequest) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{0}
}

func (x *GetFacultyRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Faculty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation string `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	Head         string `protobuf:"bytes,4,opt,name=head,proto3" json:"head,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Room         string `protobuf:"bytes,6,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *Faculty) Reset() {
	*x = Faculty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Faculty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Faculty) ProtoMessage() {}

func (x *Faculty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Faculty.ProtoReflect.Descriptor instead.
func (*Faculty) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{1}
}

func (x *Faculty) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Faculty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Faculty) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *Faculty) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

func (x *Faculty) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Faculty) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type ListFacultiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListFacultiesRequest) Reset() {
	*x = ListFacultiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFacultiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFacultiesRequest) ProtoMessage() {}

func (x *ListFacultiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFacultiesRequest.ProtoReflect.Descriptor instead.
func (*ListFacultiesRequest) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{2}
}

func (x *ListFacultiesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Meta_Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{3}
}

func (x *Meta) GetPage() *Meta_Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListFacultiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Faculty `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta *Meta      `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListFacultiesResponse) Reset() {
	*x = ListFacultiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFacultiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFacultiesResponse) ProtoMessage() {}

func (x *ListFacultiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFacultiesResponse.ProtoReflect.Descriptor instead.
func (*ListFacultiesResponse) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{4}
}

func (x *ListFacultiesResponse) GetData() []*Faculty {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListFacultiesResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Meta_Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	AllPages int32 `protobuf:"varint,2,opt,name=allPages,proto3" json:"allPages,omitempty"`
	Limit    int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Total    int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Meta_Page) Reset() {
	*x = Meta_Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_faculty_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta_Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta_Page) ProtoMessage() {}

func (x *Meta_Page) ProtoReflect() protoreflect.Message {
	mi := &file_pb_faculty_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta_Page.ProtoReflect.Descriptor instead.
func (*Meta_Page) Descriptor() ([]byte, []int) {
	return file_pb_faculty_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Meta_Page) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Meta_Page) GetAllPages() int32 {
	if x != nil {
		return x.AllPages
	}
	return 0
}

func (x *Meta_Page) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Meta_Page) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_pb_faculty_proto protoreflect.FileDescriptor

var file_pb_faculty_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x8f, 0x01, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x22, 0x2a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x92,
	0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x1a,
	0x62, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x6c, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x6c, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x60, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x61,
	0x62, 0x62, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x32, 0xa0, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61,
	0x63, 0x75, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_faculty_proto_rawDescOnce sync.Once
	file_pb_faculty_proto_rawDescData = file_pb_faculty_proto_rawDesc
)

func file_pb_faculty_proto_rawDescGZIP() []byte {
	file_pb_faculty_proto_rawDescOnce.Do(func() {
		file_pb_faculty_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_faculty_proto_rawDescData)
	})
	return file_pb_faculty_proto_rawDescData
}

var file_pb_faculty_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_faculty_proto_goTypes = []interface{}{
	(*GetFacultyRequest)(nil),     // 0: grabber.GetFacultyRequest
	(*Faculty)(nil),               // 1: grabber.Faculty
	(*ListFacultiesRequest)(nil),  // 2: grabber.ListFacultiesRequest
	(*Meta)(nil),                  // 3: grabber.Meta
	(*ListFacultiesResponse)(nil), // 4: grabber.ListFacultiesResponse
	(*Meta_Page)(nil),             // 5: grabber.Meta.Page
}
var file_pb_faculty_proto_depIdxs = []int32{
	5, // 0: grabber.Meta.page:type_name -> grabber.Meta.Page
	1, // 1: grabber.ListFacultiesResponse.data:type_name -> grabber.Faculty
	3, // 2: grabber.ListFacultiesResponse.meta:type_name -> grabber.Meta
	0, // 3: grabber.FacultyService.GetFaculty:input_type -> grabber.GetFacultyRequest
	2, // 4: grabber.FacultyService.ListFaculties:input_type -> grabber.ListFacultiesRequest
	1, // 5: grabber.FacultyService.GetFaculty:output_type -> grabber.Faculty
	4, // 6: grabber.FacultyService.ListFaculties:output_type -> grabber.ListFacultiesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_faculty_proto_init() }
func file_pb_faculty_proto_init() {
	if File_pb_faculty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_faculty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFacultyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_faculty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Faculty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_faculty_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFacultiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_faculty_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_faculty_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFacultiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_faculty_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta_Page); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_faculty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_faculty_proto_goTypes,
		DependencyIndexes: file_pb_faculty_proto_depIdxs,
		MessageInfos:      file_pb_faculty_proto_msgTypes,
	}.Build()
	File_pb_faculty_proto = out.File
	file_pb_faculty_proto_rawDesc = nil
	file_pb_faculty_proto_goTypes = nil
	file_pb_faculty_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FacultyServiceClient is the client API for FacultyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FacultyServiceClient interface {
	GetFaculty(ctx context.Context, in *GetFacultyRequest, opts ...grpc.CallOption) (*Faculty, error)
	ListFaculties(ctx context.Context, in *ListFacultiesRequest, opts ...grpc.CallOption) (*ListFacultiesResponse, error)
}

type facultyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFacultyServiceClient(cc grpc.ClientConnInterface) FacultyServiceClient {
	return &facultyServiceClient{cc}
}

func (c *facultyServiceClient) GetFaculty(ctx context.Context, in *GetFacultyRequest, opts ...grpc.CallOption) (*Faculty, error) {
	out := new(Faculty)
	err := c.cc.Invoke(ctx, "/grabber.FacultyService/GetFaculty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facultyServiceClient) ListFaculties(ctx context.Context, in *ListFacultiesRequest, opts ...grpc.CallOption) (*ListFacultiesResponse, error) {
	out := new(ListFacultiesResponse)
	err := c.cc.Invoke(ctx, "/grabber.FacultyService/ListFaculties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FacultyServiceServer is the server API for FacultyService service.
type FacultyServiceServer interface {
	GetFaculty(context.Context, *GetFacultyRequest) (*Faculty, error)
	ListFaculties(context.Context, *ListFacultiesRequest) (*ListFacultiesResponse, error)
}

// UnimplementedFacultyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFacultyServiceServer struct {
}

func (*UnimplementedFacultyServiceServer) GetFaculty(context.Context, *GetFacultyRequest) (*Faculty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFaculty not implemented")
}
func (*UnimplementedFacultyServiceServer) ListFaculties(context.Context, *ListFacultiesRequest) (*ListFacultiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFaculties not implemented")
}

func RegisterFacultyServiceServer(s *grpc.Server, srv FacultyServiceServer) {
	s.RegisterService(&_FacultyService_serviceDesc, srv)
}

func _FacultyService_GetFaculty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFacultyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacultyServiceServer).GetFaculty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grabber.FacultyService/GetFaculty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacultyServiceServer).GetFaculty(ctx, req.(*GetFacultyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacultyService_ListFaculties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFacultiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacultyServiceServer).ListFaculties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grabber.FacultyService/ListFaculties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacultyServiceServer).ListFaculties(ctx, req.(*ListFacultiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FacultyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grabber.FacultyService",
	HandlerType: (*FacultyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFaculty",
			Handler:    _FacultyService_GetFaculty_Handler,
		},
		{
			MethodName: "ListFaculties",
			Handler:    _FacultyService_ListFaculties_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/faculty.proto",
}
