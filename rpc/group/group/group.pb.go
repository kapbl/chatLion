// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: group.proto

package group

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

// 获取群成员
type GetMembersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       string                 `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMembersRequest) Reset() {
	*x = GetMembersRequest{}
	mi := &file_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersRequest) ProtoMessage() {}

func (x *GetMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersRequest.ProtoReflect.Descriptor instead.
func (*GetMembersRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *GetMembersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetMembersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Members       []string               `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"` // 成员列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMembersResponse) Reset() {
	*x = GetMembersResponse{}
	mi := &file_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersResponse) ProtoMessage() {}

func (x *GetMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersResponse.ProtoReflect.Descriptor instead.
func (*GetMembersResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetMembersResponse) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

// 创建群组
type CreateGroupRequest struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	CreateGroupName        string                 `protobuf:"bytes,1,opt,name=create_group_name,json=createGroupName,proto3" json:"create_group_name,omitempty"`
	CreateGroupDescription string                 `protobuf:"bytes,2,opt,name=create_group_description,json=createGroupDescription,proto3" json:"create_group_description,omitempty"`
	CreateOwnerName        string                 `protobuf:"bytes,3,opt,name=create_owner_name,json=createOwnerName,proto3" json:"create_owner_name,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	mi := &file_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGroupRequest) GetCreateGroupName() string {
	if x != nil {
		return x.CreateGroupName
	}
	return ""
}

func (x *CreateGroupRequest) GetCreateGroupDescription() string {
	if x != nil {
		return x.CreateGroupDescription
	}
	return ""
}

func (x *CreateGroupRequest) GetCreateOwnerName() string {
	if x != nil {
		return x.CreateOwnerName
	}
	return ""
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CreateMessage string                 `protobuf:"bytes,1,opt,name=create_message,json=createMessage,proto3" json:"create_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	mi := &file_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGroupResponse) GetCreateMessage() string {
	if x != nil {
		return x.CreateMessage
	}
	return ""
}

// 加入群组
type JoinGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JoinGroupName string                 `protobuf:"bytes,1,opt,name=join_group_name,json=joinGroupName,proto3" json:"join_group_name,omitempty"`
	JoinerName    string                 `protobuf:"bytes,2,opt,name=joiner_name,json=joinerName,proto3" json:"joiner_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinGroupRequest) Reset() {
	*x = JoinGroupRequest{}
	mi := &file_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupRequest) ProtoMessage() {}

func (x *JoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *JoinGroupRequest) GetJoinGroupName() string {
	if x != nil {
		return x.JoinGroupName
	}
	return ""
}

func (x *JoinGroupRequest) GetJoinerName() string {
	if x != nil {
		return x.JoinerName
	}
	return ""
}

type JoinGroupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JoinMessage   string                 `protobuf:"bytes,1,opt,name=join_message,json=joinMessage,proto3" json:"join_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinGroupResponse) Reset() {
	*x = JoinGroupResponse{}
	mi := &file_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupResponse) ProtoMessage() {}

func (x *JoinGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *JoinGroupResponse) GetJoinMessage() string {
	if x != nil {
		return x.JoinMessage
	}
	return ""
}

// 删除群组
type DeleteGroupRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	DeleteGroupName string                 `protobuf:"bytes,1,opt,name=delete_group_name,json=deleteGroupName,proto3" json:"delete_group_name,omitempty"`
	UserName        string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"` // 是谁删除的这个群组？如果是群主：这个群直接删除，如果不是群主，离开这个群组
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	mi := &file_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteGroupRequest) GetDeleteGroupName() string {
	if x != nil {
		return x.DeleteGroupName
	}
	return ""
}

func (x *DeleteGroupRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type DeleteGroupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeleteMessage string                 `protobuf:"bytes,1,opt,name=delete_message,json=deleteMessage,proto3" json:"delete_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGroupResponse) Reset() {
	*x = DeleteGroupResponse{}
	mi := &file_group_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResponse) ProtoMessage() {}

func (x *DeleteGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteGroupResponse) GetDeleteMessage() string {
	if x != nil {
		return x.DeleteMessage
	}
	return ""
}

var File_group_proto protoreflect.FileDescriptor

const file_group_proto_rawDesc = "" +
	"\n" +
	"\vgroup.proto\x12\x05group\".\n" +
	"\x11GetMembersRequest\x12\x19\n" +
	"\bgroup_id\x18\x01 \x01(\tR\agroupId\".\n" +
	"\x12GetMembersResponse\x12\x18\n" +
	"\amembers\x18\x01 \x03(\tR\amembers\"\xa6\x01\n" +
	"\x12CreateGroupRequest\x12*\n" +
	"\x11create_group_name\x18\x01 \x01(\tR\x0fcreateGroupName\x128\n" +
	"\x18create_group_description\x18\x02 \x01(\tR\x16createGroupDescription\x12*\n" +
	"\x11create_owner_name\x18\x03 \x01(\tR\x0fcreateOwnerName\"<\n" +
	"\x13CreateGroupResponse\x12%\n" +
	"\x0ecreate_message\x18\x01 \x01(\tR\rcreateMessage\"[\n" +
	"\x10JoinGroupRequest\x12&\n" +
	"\x0fjoin_group_name\x18\x01 \x01(\tR\rjoinGroupName\x12\x1f\n" +
	"\vjoiner_name\x18\x02 \x01(\tR\n" +
	"joinerName\"6\n" +
	"\x11JoinGroupResponse\x12!\n" +
	"\fjoin_message\x18\x01 \x01(\tR\vjoinMessage\"]\n" +
	"\x12DeleteGroupRequest\x12*\n" +
	"\x11delete_group_name\x18\x01 \x01(\tR\x0fdeleteGroupName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\"<\n" +
	"\x13DeleteGroupResponse\x12%\n" +
	"\x0edelete_message\x18\x01 \x01(\tR\rdeleteMessage2\x9f\x02\n" +
	"\x05Group\x12J\n" +
	"\x13GetMembersByGroupID\x12\x18.group.GetMembersRequest\x1a\x19.group.GetMembersResponse\x12D\n" +
	"\vCreateGroup\x12\x19.group.CreateGroupRequest\x1a\x1a.group.CreateGroupResponse\x12>\n" +
	"\tJoinGroup\x12\x17.group.JoinGroupRequest\x1a\x18.group.JoinGroupResponse\x12D\n" +
	"\vDeleteGroup\x12\x19.group.DeleteGroupRequest\x1a\x1a.group.DeleteGroupResponseB\tZ\a./groupb\x06proto3"

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData []byte
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_group_proto_rawDesc), len(file_group_proto_rawDesc)))
	})
	return file_group_proto_rawDescData
}

var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_group_proto_goTypes = []any{
	(*GetMembersRequest)(nil),   // 0: group.GetMembersRequest
	(*GetMembersResponse)(nil),  // 1: group.GetMembersResponse
	(*CreateGroupRequest)(nil),  // 2: group.CreateGroupRequest
	(*CreateGroupResponse)(nil), // 3: group.CreateGroupResponse
	(*JoinGroupRequest)(nil),    // 4: group.JoinGroupRequest
	(*JoinGroupResponse)(nil),   // 5: group.JoinGroupResponse
	(*DeleteGroupRequest)(nil),  // 6: group.DeleteGroupRequest
	(*DeleteGroupResponse)(nil), // 7: group.DeleteGroupResponse
}
var file_group_proto_depIdxs = []int32{
	0, // 0: group.Group.GetMembersByGroupID:input_type -> group.GetMembersRequest
	2, // 1: group.Group.CreateGroup:input_type -> group.CreateGroupRequest
	4, // 2: group.Group.JoinGroup:input_type -> group.JoinGroupRequest
	6, // 3: group.Group.DeleteGroup:input_type -> group.DeleteGroupRequest
	1, // 4: group.Group.GetMembersByGroupID:output_type -> group.GetMembersResponse
	3, // 5: group.Group.CreateGroup:output_type -> group.CreateGroupResponse
	5, // 6: group.Group.JoinGroup:output_type -> group.JoinGroupResponse
	7, // 7: group.Group.DeleteGroup:output_type -> group.DeleteGroupResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_group_proto_rawDesc), len(file_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
