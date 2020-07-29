// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: proto/staff.proto

package staff

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type LeavingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line []*LeavingListLine `protobuf:"bytes,1,rep,name=line,proto3" json:"line,omitempty"`
}

func (x *LeavingList) Reset() {
	*x = LeavingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeavingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeavingList) ProtoMessage() {}

func (x *LeavingList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeavingList.ProtoReflect.Descriptor instead.
func (*LeavingList) Descriptor() ([]byte, []int) {
	return file_proto_staff_proto_rawDescGZIP(), []int{0}
}

func (x *LeavingList) GetLine() []*LeavingListLine {
	if x != nil {
		return x.Line
	}
	return nil
}

type LeavingListLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Employee *Employee `protobuf:"bytes,2,opt,name=employee,proto3" json:"employee,omitempty"`
	Action   []*Action `protobuf:"bytes,3,rep,name=action,proto3" json:"action,omitempty"`
	Marks    []*Marks  `protobuf:"bytes,4,rep,name=marks,proto3" json:"marks,omitempty"`
}

func (x *LeavingListLine) Reset() {
	*x = LeavingListLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeavingListLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeavingListLine) ProtoMessage() {}

func (x *LeavingListLine) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeavingListLine.ProtoReflect.Descriptor instead.
func (*LeavingListLine) Descriptor() ([]byte, []int) {
	return file_proto_staff_proto_rawDescGZIP(), []int{1}
}

func (x *LeavingListLine) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeavingListLine) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

func (x *LeavingListLine) GetAction() []*Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *LeavingListLine) GetMarks() []*Marks {
	if x != nil {
		return x.Marks
	}
	return nil
}

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname       string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	DateDismissal string `protobuf:"bytes,4,opt,name=dateDismissal,proto3" json:"dateDismissal,omitempty"`
	Subdivision   string `protobuf:"bytes,5,opt,name=subdivision,proto3" json:"subdivision,omitempty"`
	Position      string `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty"`
	MobilePhone   string `protobuf:"bytes,7,opt,name=mobilePhone,proto3" json:"mobilePhone,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_staff_proto_rawDescGZIP(), []int{2}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Employee) GetDateDismissal() string {
	if x != nil {
		return x.DateDismissal
	}
	return ""
}

func (x *Employee) GetSubdivision() string {
	if x != nil {
		return x.Subdivision
	}
	return ""
}

func (x *Employee) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Employee) GetMobilePhone() string {
	if x != nil {
		return x.MobilePhone
	}
	return ""
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle string   `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Value    string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Option   []string `protobuf:"bytes,5,rep,name=option,proto3" json:"option,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_proto_staff_proto_rawDescGZIP(), []int{3}
}

func (x *Action) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Action) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Action) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Action) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Action) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

type Marks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Mark     bool   `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	Disabled bool   `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *Marks) Reset() {
	*x = Marks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Marks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Marks) ProtoMessage() {}

func (x *Marks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Marks.ProtoReflect.Descriptor instead.
func (*Marks) Descriptor() ([]byte, []int) {
	return file_proto_staff_proto_rawDescGZIP(), []int{4}
}

func (x *Marks) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Marks) GetMark() bool {
	if x != nil {
		return x.Mark
	}
	return false
}

func (x *Marks) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

var File_proto_staff_proto protoreflect.FileDescriptor

var file_proto_staff_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x1f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22,
	0xce, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x61, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x22, 0x78, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x05, 0x4d, 0x61,
	0x72, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x12, 0x3d, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_staff_proto_rawDescOnce sync.Once
	file_proto_staff_proto_rawDescData = file_proto_staff_proto_rawDesc
)

func file_proto_staff_proto_rawDescGZIP() []byte {
	file_proto_staff_proto_rawDescOnce.Do(func() {
		file_proto_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_staff_proto_rawDescData)
	})
	return file_proto_staff_proto_rawDescData
}

var file_proto_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_staff_proto_goTypes = []interface{}{
	(*LeavingList)(nil),     // 0: LeavingList
	(*LeavingListLine)(nil), // 1: LeavingListLine
	(*Employee)(nil),        // 2: Employee
	(*Action)(nil),          // 3: Action
	(*Marks)(nil),           // 4: Marks
	(*empty.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_proto_staff_proto_depIdxs = []int32{
	1, // 0: LeavingList.line:type_name -> LeavingListLine
	2, // 1: LeavingListLine.employee:type_name -> Employee
	3, // 2: LeavingListLine.action:type_name -> Action
	4, // 3: LeavingListLine.marks:type_name -> Marks
	5, // 4: staff.getLeavingEmployees:input_type -> google.protobuf.Empty
	0, // 5: staff.getLeavingEmployees:output_type -> LeavingList
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_staff_proto_init() }
func file_proto_staff_proto_init() {
	if File_proto_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeavingList); i {
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
		file_proto_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeavingListLine); i {
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
		file_proto_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_proto_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_proto_staff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Marks); i {
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
			RawDescriptor: file_proto_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_staff_proto_goTypes,
		DependencyIndexes: file_proto_staff_proto_depIdxs,
		MessageInfos:      file_proto_staff_proto_msgTypes,
	}.Build()
	File_proto_staff_proto = out.File
	file_proto_staff_proto_rawDesc = nil
	file_proto_staff_proto_goTypes = nil
	file_proto_staff_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StaffClient is the client API for Staff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffClient interface {
	GetLeavingEmployees(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LeavingList, error)
}

type staffClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffClient(cc grpc.ClientConnInterface) StaffClient {
	return &staffClient{cc}
}

func (c *staffClient) GetLeavingEmployees(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LeavingList, error) {
	out := new(LeavingList)
	err := c.cc.Invoke(ctx, "/staff/getLeavingEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServer is the server API for Staff service.
type StaffServer interface {
	GetLeavingEmployees(context.Context, *empty.Empty) (*LeavingList, error)
}

// UnimplementedStaffServer can be embedded to have forward compatible implementations.
type UnimplementedStaffServer struct {
}

func (*UnimplementedStaffServer) GetLeavingEmployees(context.Context, *empty.Empty) (*LeavingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeavingEmployees not implemented")
}

func RegisterStaffServer(s *grpc.Server, srv StaffServer) {
	s.RegisterService(&_Staff_serviceDesc, srv)
}

func _Staff_GetLeavingEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).GetLeavingEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff/GetLeavingEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).GetLeavingEmployees(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Staff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "staff",
	HandlerType: (*StaffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getLeavingEmployees",
			Handler:    _Staff_GetLeavingEmployees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/staff.proto",
}
