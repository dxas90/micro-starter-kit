// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/account/proto/user/user_service.proto

package user

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	entities "github.com/xmlking/micro-starter-kit/srv/account/proto/entities"
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

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// with Workaround in .override.go
type ExistRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExistRequest) Reset()         { *m = ExistRequest{} }
func (m *ExistRequest) String() string { return proto.CompactTextString(m) }
func (*ExistRequest) ProtoMessage()    {}
func (*ExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{0}
}

func (m *ExistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistRequest.Unmarshal(m, b)
}
func (m *ExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistRequest.Marshal(b, m, deterministic)
}
func (m *ExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistRequest.Merge(m, src)
}
func (m *ExistRequest) XXX_Size() int {
	return xxx_messageInfo_ExistRequest.Size(m)
}
func (m *ExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistRequest proto.InternalMessageInfo

func (m *ExistRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ExistRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ExistRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *ExistRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *ExistRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type ExistResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistResponse) Reset()         { *m = ExistResponse{} }
func (m *ExistResponse) String() string { return proto.CompactTextString(m) }
func (*ExistResponse) ProtoMessage()    {}
func (*ExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{1}
}

func (m *ExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistResponse.Unmarshal(m, b)
}
func (m *ExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistResponse.Marshal(b, m, deterministic)
}
func (m *ExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistResponse.Merge(m, src)
}
func (m *ExistResponse) XXX_Size() int {
	return xxx_messageInfo_ExistResponse.Size(m)
}
func (m *ExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExistResponse proto.InternalMessageInfo

func (m *ExistResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type ListRequest struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ListRequest) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListRequest) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *ListRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ListRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *ListRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *ListRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type ListResponse struct {
	Results              []*entities.User `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetResults() []*entities.User {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *GetRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *GetRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *GetRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type GetResponse struct {
	Result               *entities.User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetResult() *entities.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateRequest struct {
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{6}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *CreateRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *CreateRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *CreateRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type CreateResponse struct {
	Result               *entities.User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{7}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetResult() *entities.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UpdateRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *UpdateRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UpdateRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *UpdateRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type UpdateResponse struct {
	Result               *entities.User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetResult() *entities.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeleteRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *DeleteRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *DeleteRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *DeleteRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type DeleteResponse struct {
	Result               *entities.User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b5b5d017608cf5, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetResult() *entities.User {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ExistRequest)(nil), "micro.account.user.v1.ExistRequest")
	proto.RegisterType((*ExistResponse)(nil), "micro.account.user.v1.ExistResponse")
	proto.RegisterType((*ListRequest)(nil), "micro.account.user.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "micro.account.user.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "micro.account.user.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "micro.account.user.v1.GetResponse")
	proto.RegisterType((*CreateRequest)(nil), "micro.account.user.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "micro.account.user.v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "micro.account.user.v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "micro.account.user.v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "micro.account.user.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "micro.account.user.v1.DeleteResponse")
}

func init() {
	proto.RegisterFile("srv/account/proto/user/user_service.proto", fileDescriptor_70b5b5d017608cf5)
}

var fileDescriptor_70b5b5d017608cf5 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x55, 0xd2, 0xb4, 0xeb, 0x7e, 0x5d, 0xd1, 0x64, 0x01, 0xaa, 0xca, 0xbf, 0xd1, 0x0d, 0xa9,
	0x4c, 0x6b, 0xba, 0x75, 0x02, 0x34, 0x26, 0x2e, 0x61, 0x68, 0x42, 0x4c, 0x03, 0x65, 0xea, 0x0e,
	0x43, 0x50, 0xbc, 0xd6, 0xeb, 0xac, 0xa5, 0x49, 0xb0, 0x9d, 0x32, 0x40, 0x48, 0x5c, 0xf8, 0x3e,
	0x88, 0xcf, 0xc3, 0x1d, 0xc1, 0x27, 0x40, 0x3d, 0xa1, 0xc4, 0x49, 0xd7, 0xa0, 0x66, 0x2b, 0xd0,
	0x71, 0xea, 0x25, 0x72, 0xa2, 0xf7, 0x5e, 0x9c, 0xdf, 0xfb, 0x3d, 0xdb, 0x81, 0xdb, 0x9c, 0x75,
	0xab, 0xb8, 0xd9, 0x74, 0x3c, 0x5b, 0x54, 0x5d, 0xe6, 0x08, 0xa7, 0xea, 0x71, 0xc2, 0x82, 0x4b,
	0x83, 0x13, 0xd6, 0xa5, 0x4d, 0xa2, 0x07, 0xcf, 0xd1, 0xa5, 0x0e, 0x6d, 0x32, 0x47, 0x0f, 0xc1,
	0xba, 0x8f, 0xd0, 0xbb, 0x2b, 0xc5, 0xeb, 0x6d, 0xc7, 0x69, 0x5b, 0x44, 0x92, 0xf7, 0xbd, 0x83,
	0xea, 0x1b, 0x86, 0x5d, 0x97, 0x30, 0x2e, 0x69, 0xc5, 0x21, 0x6f, 0x20, 0xb6, 0xa0, 0x82, 0x12,
	0xde, 0x1f, 0x44, 0x50, 0x71, 0x48, 0x59, 0xab, 0xe1, 0x62, 0x26, 0xde, 0x86, 0xd0, 0x2e, 0xb6,
	0x68, 0x0b, 0x0b, 0xd2, 0x1f, 0x48, 0x68, 0xe9, 0xbb, 0x0a, 0x33, 0x8f, 0x8e, 0x29, 0x17, 0x26,
	0x79, 0xed, 0x11, 0x2e, 0xd0, 0x5d, 0x50, 0x69, 0xab, 0xa0, 0xcc, 0x29, 0xe5, 0x5c, 0xed, 0xaa,
	0x2e, 0xe7, 0xa4, 0x47, 0x73, 0xd2, 0x77, 0x04, 0xa3, 0x76, 0x7b, 0x17, 0x5b, 0x1e, 0x31, 0xb2,
	0x3d, 0x23, 0xcd, 0x52, 0x9f, 0x15, 0xc5, 0x54, 0x69, 0x0b, 0xed, 0x41, 0xd6, 0xff, 0x12, 0x1b,
	0x77, 0x48, 0x41, 0x1d, 0x81, 0x3d, 0xd7, 0x33, 0xae, 0xb1, 0x2b, 0xb3, 0x5a, 0x61, 0xb6, 0xfc,
	0x51, 0xad, 0xa1, 0x97, 0xcf, 0x71, 0xe5, 0xdd, 0x72, 0x65, 0xad, 0x51, 0x79, 0xf1, 0x7e, 0x75,
	0x69, 0xe5, 0xce, 0x87, 0x05, 0xb3, 0xaf, 0x87, 0x36, 0x00, 0x0e, 0x28, 0xe3, 0xa2, 0x11, 0xa8,
	0xa7, 0x46, 0x50, 0x9f, 0xea, 0x19, 0x1a, 0x53, 0x67, 0x53, 0xe6, 0x74, 0x40, 0xdc, 0xf6, 0x55,
	0x0c, 0x98, 0xb6, 0x70, 0x24, 0xa2, 0xfd, 0x89, 0x48, 0xd6, 0xe7, 0x05, 0x1a, 0xeb, 0x90, 0x26,
	0x1d, 0x4c, 0xad, 0x42, 0x7a, 0x64, 0xfe, 0x2b, 0xc5, 0x94, 0x9c, 0xd2, 0x12, 0xe4, 0xc3, 0x52,
	0x73, 0xd7, 0xb1, 0x39, 0x41, 0x97, 0x21, 0xc3, 0x08, 0xf7, 0x2c, 0x11, 0xd4, 0x3b, 0x6b, 0x86,
	0x77, 0xf7, 0x53, 0x3f, 0x0d, 0xa5, 0xf4, 0x35, 0x05, 0xb9, 0xad, 0x01, 0x63, 0x1e, 0x40, 0xda,
	0xa2, 0x1d, 0x2a, 0x12, 0xbd, 0xa9, 0x3f, 0xb6, 0xc5, 0x6a, 0x4d, 0xbe, 0x7a, 0xba, 0x67, 0x64,
	0x16, 0xb5, 0x42, 0xab, 0xac, 0x98, 0x92, 0x85, 0xd6, 0x40, 0x73, 0x71, 0x3b, 0xd9, 0x9b, 0x41,
	0xb6, 0x3f, 0xf1, 0x45, 0xb5, 0xac, 0x98, 0x01, 0x05, 0x2d, 0x83, 0xc6, 0x1d, 0x26, 0x46, 0x29,
	0xbc, 0x19, 0x20, 0x63, 0xcd, 0xa0, 0x9d, 0x6b, 0x33, 0xa4, 0xc7, 0xd1, 0x0c, 0x99, 0x7f, 0x6c,
	0x86, 0xa9, 0xbf, 0x68, 0x86, 0x16, 0xcc, 0x6c, 0x0d, 0xf6, 0xc2, 0x1a, 0x4c, 0x49, 0xf7, 0x79,
	0x41, 0x99, 0x4b, 0x95, 0x73, 0xb5, 0x1b, 0x7a, 0x7c, 0x9d, 0xe8, 0x67, 0xbc, 0xbb, 0xa2, 0xd7,
	0x39, 0x61, 0x66, 0x84, 0x47, 0x17, 0x21, 0x2d, 0x1c, 0x81, 0xad, 0xc0, 0xdb, 0xbc, 0x29, 0x6f,
	0x64, 0x13, 0x7d, 0x53, 0x01, 0x36, 0xc9, 0x24, 0xdc, 0xe7, 0x1f, 0xee, 0x27, 0x90, 0x0b, 0x0a,
	0x1d, 0xda, 0x79, 0x2f, 0x16, 0xed, 0x11, 0xdc, 0x8c, 0x65, 0xff, 0x8b, 0x0a, 0xf9, 0x87, 0x8c,
	0x60, 0x41, 0x22, 0xe7, 0x26, 0x0e, 0x9c, 0xe9, 0xc0, 0x36, 0x5c, 0x88, 0x6a, 0x36, 0x16, 0x13,
	0x7e, 0xa8, 0x90, 0xaf, 0xbb, 0xad, 0x01, 0x13, 0x26, 0xf1, 0x39, 0x4f, 0xf3, 0xa2, 0x5a, 0x8f,
	0xcd, 0xbc, 0x0d, 0x62, 0x91, 0x89, 0x79, 0xff, 0xc7, 0xbc, 0xa8, 0xd6, 0xe3, 0x30, 0xaf, 0xf6,
	0x49, 0x83, 0x9c, 0xff, 0x74, 0x47, 0x9e, 0x9b, 0x91, 0x09, 0xe9, 0xe0, 0xe0, 0x84, 0xe6, 0xf5,
	0xa1, 0x67, 0x67, 0x7d, 0xf0, 0x04, 0x5b, 0x5c, 0x38, 0x1d, 0x14, 0xce, 0xf0, 0x29, 0x68, 0xfe,
	0xfe, 0x8b, 0x4a, 0x09, 0xe8, 0x81, 0xa3, 0x57, 0x71, 0xfe, 0x54, 0x4c, 0x28, 0xb8, 0x05, 0xa9,
	0x4d, 0x22, 0xd0, 0xcd, 0x04, 0xec, 0xc9, 0x2e, 0x5c, 0x2c, 0x9d, 0x06, 0x09, 0xd5, 0xea, 0x90,
	0x91, 0x8b, 0x19, 0x4a, 0xfa, 0x9c, 0xd8, 0xfe, 0x50, 0xbc, 0x75, 0x06, 0xea, 0x44, 0x56, 0xc6,
	0x2c, 0x51, 0x36, 0xb6, 0xe2, 0x25, 0xca, 0xfe, 0x96, 0xd5, 0x3a, 0x64, 0x64, 0x03, 0x24, 0xca,
	0xc6, 0xb2, 0x98, 0x28, 0x1b, 0xef, 0x22, 0x63, 0x17, 0x86, 0xff, 0x2b, 0x3d, 0x53, 0xf6, 0xd6,
	0xdb, 0x54, 0x1c, 0x7a, 0xfb, 0x7a, 0xd3, 0xe9, 0x54, 0x8f, 0x3b, 0xd6, 0x11, 0xb5, 0xdb, 0xd5,
	0x00, 0x5b, 0xe1, 0x02, 0x33, 0x41, 0x58, 0xe5, 0x88, 0x8a, 0xea, 0xf0, 0x9f, 0xb2, 0xfd, 0x4c,
	0x30, 0x5e, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x73, 0x59, 0xee, 0x55, 0xb5, 0x0d, 0x00, 0x00,
}
