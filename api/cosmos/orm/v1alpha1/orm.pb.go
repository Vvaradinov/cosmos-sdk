// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/orm/v1alpha1/orm.proto

package ormv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TableDescriptor describes an ORM table.
type TableDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// primary_key defines the primary key for the table.
	PrimaryKey *PrimaryKeyDescriptor `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	// index defines one or more secondary indexes.
	Index []*SecondaryIndexDescriptor `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	// id is a non-zero integer ID that must be unique within the
	// tables and singletons in this file. It may be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TableDescriptor) Reset() {
	*x = TableDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableDescriptor) ProtoMessage() {}

func (x *TableDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableDescriptor.ProtoReflect.Descriptor instead.
func (*TableDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{0}
}

func (x *TableDescriptor) GetPrimaryKey() *PrimaryKeyDescriptor {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *TableDescriptor) GetIndex() []*SecondaryIndexDescriptor {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *TableDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// PrimaryKeyDescriptor describes a table primary key.
type PrimaryKeyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fields is a comma-separated list of fields in the primary key. Spaces are
	// not allowed. Supported field types, their encodings, and any applicable constraints
	// are described below.
	//   - uint32, uint64 are encoded as big-endian fixed width bytes and support
	//   sorted iteration.
	//   - string's are encoded as raw bytes in terminal key segments and null-terminated
	//   in non-terminal segments. Null characters are thus forbidden in strings.
	//   string fields support sorted iteration.
	//   - bytes are encoded as raw bytes in terminal segments and length-prefixed
	//   with a single byte in non-terminal segments. Because of this byte arrays
	//   longer than 255 bytes are unsupported and bytes fields should not
	//   be assumed to be lexically sorted. If you have a byte array longer than
	//   255 bytes that you'd like to index, you should consider hashing it first.
	//   - int32, sint32, int64, sint64 are encoding as fixed width bytes with
	//   an encoding that enables sorted iteration.
	//   - google.protobuf.Timestamp and google.protobuf.Duration are encoded
	//   as 12 bytes using an encoding that enables sorted iteration.
	//   - enum fields are encoded using varint encoding and do not support sorted
	//   iteration.
	//   - bool fields are encoded as a single byte 0 or 1.
	//
	// All other fields types are unsupported in keys including repeated and
	// oneof fields.
	//
	// Primary keys are prefixed by the varint encoded table id and the byte 0x0
	// plus any additional prefix specified by the schema.
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// auto_increment specifies that the primary key is generated by an
	// auto-incrementing integer. If this is set to true fields must only
	// contain one field of that is of type uint64.
	AutoIncrement bool `protobuf:"varint,2,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	// references specifies that this primary key references the primary key
	// of another table. See the documentation for the SecondaryIndexDescriptor.references
	// field for more details. An additional constraint placed on primary keys
	// which reference another table is that those references cannot be circular.
	References string `protobuf:"bytes,3,opt,name=references,proto3" json:"references,omitempty"`
}

func (x *PrimaryKeyDescriptor) Reset() {
	*x = PrimaryKeyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryKeyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryKeyDescriptor) ProtoMessage() {}

func (x *PrimaryKeyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryKeyDescriptor.ProtoReflect.Descriptor instead.
func (*PrimaryKeyDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{1}
}

func (x *PrimaryKeyDescriptor) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *PrimaryKeyDescriptor) GetAutoIncrement() bool {
	if x != nil {
		return x.AutoIncrement
	}
	return false
}

func (x *PrimaryKeyDescriptor) GetReferences() string {
	if x != nil {
		return x.References
	}
	return ""
}

// PrimaryKeyDescriptor describes a table secondary index.
type SecondaryIndexDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fields is a comma-separated list of fields in the index. The supported
	// field types are the same as those for PrimaryKeyDescriptor.fields.
	// Index keys are prefixed by the varint encoded table id and the varint
	// encoded index id plus any additional prefix specified by the schema.
	//
	// In addition the the field segments, non-unique index keys are suffixed with
	// any additional primary key fields not present in the index fields so that the
	// primary key can be reconstructed. Unique indexes instead of being suffixed
	// store the remaining primary key fields in the value..
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	// id is a non-zero integer ID that must be unique within the indexes for this
	// table and less than 32768. It may be deprecated in the future when this can
	// be auto-generated.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// unique specifies that this an unique index.
	Unique bool `protobuf:"varint,3,opt,name=unique,proto3" json:"unique,omitempty"`
	// references specifies that this index references another table defined in the same
	// proto file. Currently references are not support to tables with composite
	// primary keys, therefore fields must reference one field and its type must
	// be the same type as the primary key field of the referenced table.
	// References to tables in defined by different proto files are not supported
	// to avoid tight coupling of dependencies. Beyond validating that the reference
	// is valid key constraints are currently not enforced, but references should
	// be used by clients to perform automatic joins.
	References string `protobuf:"bytes,4,opt,name=references,proto3" json:"references,omitempty"`
}

func (x *SecondaryIndexDescriptor) Reset() {
	*x = SecondaryIndexDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondaryIndexDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondaryIndexDescriptor) ProtoMessage() {}

func (x *SecondaryIndexDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondaryIndexDescriptor.ProtoReflect.Descriptor instead.
func (*SecondaryIndexDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{2}
}

func (x *SecondaryIndexDescriptor) GetFields() string {
	if x != nil {
		return x.Fields
	}
	return ""
}

func (x *SecondaryIndexDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecondaryIndexDescriptor) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *SecondaryIndexDescriptor) GetReferences() string {
	if x != nil {
		return x.References
	}
	return ""
}

// TableDescriptor describes an ORM singleton table which has at most one instance.
type SingletonDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a non-zero integer ID that must be unique within the
	// tables and singletons in this file. It may be deprecated in the future when this
	// can be auto-generated.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingletonDescriptor) Reset() {
	*x = SingletonDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingletonDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingletonDescriptor) ProtoMessage() {}

func (x *SingletonDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_orm_v1alpha1_orm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingletonDescriptor.ProtoReflect.Descriptor instead.
func (*SingletonDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_orm_v1alpha1_orm_proto_rawDescGZIP(), []int{3}
}

func (x *SingletonDescriptor) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var file_cosmos_orm_v1alpha1_orm_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         104503792,
		Name:          "cosmos.orm.v1alpha1.module_file_id",
		Tag:           "varint,104503792,opt,name=module_file_id",
		Filename:      "cosmos/orm/v1alpha1/orm.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*TableDescriptor)(nil),
		Field:         104503790,
		Name:          "cosmos.orm.v1alpha1.table",
		Tag:           "bytes,104503790,opt,name=table",
		Filename:      "cosmos/orm/v1alpha1/orm.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*SingletonDescriptor)(nil),
		Field:         104503791,
		Name:          "cosmos.orm.v1alpha1.singleton",
		Tag:           "bytes,104503791,opt,name=singleton",
		Filename:      "cosmos/orm/v1alpha1/orm.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// module_file_id assigns an ID to a .proto file which should be unique
	// within the scope of a module's .proto files. It will be used as the default
	// value in assigning IDs to FileDescriptor's used in the ORM's schema and
	// can be overridden manually in case there is any conflict or the value is
	// missing.
	//
	// optional uint32 module_file_id = 104503792;
	E_ModuleFileId = &file_cosmos_orm_v1alpha1_orm_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// table specifies that this message will be used as an ORM table. It cannot
	// be used together with the singleton option.
	//
	// optional cosmos.orm.v1alpha1.TableDescriptor table = 104503790;
	E_Table = &file_cosmos_orm_v1alpha1_orm_proto_extTypes[1]
	// singleton specifies that this message will be used as an ORM singleton. It cannot
	// be used together with the table option.
	//
	// optional cosmos.orm.v1alpha1.SingletonDescriptor singleton = 104503791;
	E_Singleton = &file_cosmos_orm_v1alpha1_orm_proto_extTypes[2]
)

var File_cosmos_orm_v1alpha1_orm_proto protoreflect.FileDescriptor

var file_cosmos_orm_v1alpha1_orm_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x14, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x7a, 0x0a, 0x18, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x25,
	0x0a, 0x13, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x45, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf0, 0xb3, 0xea, 0x31, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x3a, 0x5e, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xee, 0xb3, 0xea, 0x31, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x6a, 0x0a, 0x09,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xef, 0xb3, 0xea, 0x31, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x74, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x42, 0xd3, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x4f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6f, 0x72, 0x6d, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4f, 0x58, 0xaa, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x4f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x4f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x4f, 0x72,
	0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a,
	0x3a, 0x4f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_orm_v1alpha1_orm_proto_rawDescOnce sync.Once
	file_cosmos_orm_v1alpha1_orm_proto_rawDescData = file_cosmos_orm_v1alpha1_orm_proto_rawDesc
)

func file_cosmos_orm_v1alpha1_orm_proto_rawDescGZIP() []byte {
	file_cosmos_orm_v1alpha1_orm_proto_rawDescOnce.Do(func() {
		file_cosmos_orm_v1alpha1_orm_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_orm_v1alpha1_orm_proto_rawDescData)
	})
	return file_cosmos_orm_v1alpha1_orm_proto_rawDescData
}

var file_cosmos_orm_v1alpha1_orm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_orm_v1alpha1_orm_proto_goTypes = []interface{}{
	(*TableDescriptor)(nil),             // 0: cosmos.orm.v1alpha1.TableDescriptor
	(*PrimaryKeyDescriptor)(nil),        // 1: cosmos.orm.v1alpha1.PrimaryKeyDescriptor
	(*SecondaryIndexDescriptor)(nil),    // 2: cosmos.orm.v1alpha1.SecondaryIndexDescriptor
	(*SingletonDescriptor)(nil),         // 3: cosmos.orm.v1alpha1.SingletonDescriptor
	(*descriptorpb.FileOptions)(nil),    // 4: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
}
var file_cosmos_orm_v1alpha1_orm_proto_depIdxs = []int32{
	1, // 0: cosmos.orm.v1alpha1.TableDescriptor.primary_key:type_name -> cosmos.orm.v1alpha1.PrimaryKeyDescriptor
	2, // 1: cosmos.orm.v1alpha1.TableDescriptor.index:type_name -> cosmos.orm.v1alpha1.SecondaryIndexDescriptor
	4, // 2: cosmos.orm.v1alpha1.module_file_id:extendee -> google.protobuf.FileOptions
	5, // 3: cosmos.orm.v1alpha1.table:extendee -> google.protobuf.MessageOptions
	5, // 4: cosmos.orm.v1alpha1.singleton:extendee -> google.protobuf.MessageOptions
	0, // 5: cosmos.orm.v1alpha1.table:type_name -> cosmos.orm.v1alpha1.TableDescriptor
	3, // 6: cosmos.orm.v1alpha1.singleton:type_name -> cosmos.orm.v1alpha1.SingletonDescriptor
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	5, // [5:7] is the sub-list for extension type_name
	2, // [2:5] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_orm_v1alpha1_orm_proto_init() }
func file_cosmos_orm_v1alpha1_orm_proto_init() {
	if File_cosmos_orm_v1alpha1_orm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_orm_v1alpha1_orm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableDescriptor); i {
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
		file_cosmos_orm_v1alpha1_orm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryKeyDescriptor); i {
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
		file_cosmos_orm_v1alpha1_orm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondaryIndexDescriptor); i {
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
		file_cosmos_orm_v1alpha1_orm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingletonDescriptor); i {
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
			RawDescriptor: file_cosmos_orm_v1alpha1_orm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_orm_v1alpha1_orm_proto_goTypes,
		DependencyIndexes: file_cosmos_orm_v1alpha1_orm_proto_depIdxs,
		MessageInfos:      file_cosmos_orm_v1alpha1_orm_proto_msgTypes,
		ExtensionInfos:    file_cosmos_orm_v1alpha1_orm_proto_extTypes,
	}.Build()
	File_cosmos_orm_v1alpha1_orm_proto = out.File
	file_cosmos_orm_v1alpha1_orm_proto_rawDesc = nil
	file_cosmos_orm_v1alpha1_orm_proto_goTypes = nil
	file_cosmos_orm_v1alpha1_orm_proto_depIdxs = nil
}