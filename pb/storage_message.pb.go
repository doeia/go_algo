// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: proto/storage_message.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Storage_Driver int32

const (
	Storage_UNKONWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

// Enum value maps for Storage_Driver.
var (
	Storage_Driver_name = map[int32]string{
		0: "UNKONWN",
		1: "HDD",
		2: "SSD",
	}
	Storage_Driver_value = map[string]int32{
		"UNKONWN": 0,
		"HDD":     1,
		"SSD":     2,
	}
)

func (x Storage_Driver) Enum() *Storage_Driver {
	p := new(Storage_Driver)
	*p = x
	return p
}

func (x Storage_Driver) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Storage_Driver) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_storage_message_proto_enumTypes[0].Descriptor()
}

func (Storage_Driver) Type() protoreflect.EnumType {
	return &file_proto_storage_message_proto_enumTypes[0]
}

func (x Storage_Driver) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Storage_Driver.Descriptor instead.
func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return file_proto_storage_message_proto_rawDescGZIP(), []int{0, 0}
}

type Storage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        Storage_Driver         `protobuf:"varint,1,opt,name=driver,proto3,enum=Storage_Driver" json:"driver,omitempty"`
	Memory        *Memory                `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Storage) Reset() {
	*x = Storage{}
	mi := &file_proto_storage_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_proto_storage_message_proto_rawDescGZIP(), []int{0}
}

func (x *Storage) GetDriver() Storage_Driver {
	if x != nil {
		return x.Driver
	}
	return Storage_UNKONWN
}

func (x *Storage) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

var File_proto_storage_message_proto protoreflect.FileDescriptor

var file_proto_storage_message_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x27,
	0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4f,
	0x4e, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x44, 0x44, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x53, 0x44, 0x10, 0x02, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_storage_message_proto_rawDescOnce sync.Once
	file_proto_storage_message_proto_rawDescData []byte
)

func file_proto_storage_message_proto_rawDescGZIP() []byte {
	file_proto_storage_message_proto_rawDescOnce.Do(func() {
		file_proto_storage_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_storage_message_proto_rawDesc), len(file_proto_storage_message_proto_rawDesc)))
	})
	return file_proto_storage_message_proto_rawDescData
}

var file_proto_storage_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_storage_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_storage_message_proto_goTypes = []any{
	(Storage_Driver)(0), // 0: Storage.Driver
	(*Storage)(nil),     // 1: Storage
	(*Memory)(nil),      // 2: Memory
}
var file_proto_storage_message_proto_depIdxs = []int32{
	0, // 0: Storage.driver:type_name -> Storage.Driver
	2, // 1: Storage.memory:type_name -> Memory
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_storage_message_proto_init() }
func file_proto_storage_message_proto_init() {
	if File_proto_storage_message_proto != nil {
		return
	}
	file_proto_memery_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_storage_message_proto_rawDesc), len(file_proto_storage_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_storage_message_proto_goTypes,
		DependencyIndexes: file_proto_storage_message_proto_depIdxs,
		EnumInfos:         file_proto_storage_message_proto_enumTypes,
		MessageInfos:      file_proto_storage_message_proto_msgTypes,
	}.Build()
	File_proto_storage_message_proto = out.File
	file_proto_storage_message_proto_goTypes = nil
	file_proto_storage_message_proto_depIdxs = nil
}
