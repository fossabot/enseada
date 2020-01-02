// Copyright 2019-2020 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/v1beta1/users_api.proto

package authv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Request body to list all registered OAuth users.
type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{0}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

// Response body with a list of all registered OAuth users.
type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{1}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

// Request body to get a single OAuth user.
type GetUserRequest struct {
	// The username. Required.
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// Response body with a single OAuth user.
type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{3}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Request body to create a new OAuth user.
type CreateUserRequest struct {
	// The user to create. Required.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The user password. Required.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Response body with the newly created OAuth user.
type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Request body to update an OAuth user.
type UpdateUserRequest struct {
	// The user to update. Required.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{6}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Response body with the updated OAuth user.
type UpdateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{7}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// Request body to update an OAuth user password.
// The target user is always the current authenticated user.
type UpdateUserPasswordRequest struct {
	// The new password. Required.
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserPasswordRequest) Reset()         { *m = UpdateUserPasswordRequest{} }
func (m *UpdateUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserPasswordRequest) ProtoMessage()    {}
func (*UpdateUserPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{8}
}

func (m *UpdateUserPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserPasswordRequest.Unmarshal(m, b)
}
func (m *UpdateUserPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserPasswordRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserPasswordRequest.Merge(m, src)
}
func (m *UpdateUserPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserPasswordRequest.Size(m)
}
func (m *UpdateUserPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserPasswordRequest proto.InternalMessageInfo

func (m *UpdateUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateUserPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserPasswordResponse) Reset()         { *m = UpdateUserPasswordResponse{} }
func (m *UpdateUserPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserPasswordResponse) ProtoMessage()    {}
func (*UpdateUserPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{9}
}

func (m *UpdateUserPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserPasswordResponse.Unmarshal(m, b)
}
func (m *UpdateUserPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserPasswordResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserPasswordResponse.Merge(m, src)
}
func (m *UpdateUserPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserPasswordResponse.Size(m)
}
func (m *UpdateUserPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserPasswordResponse proto.InternalMessageInfo

// Request body to delete an OAuth user.
type DeleteUserRequest struct {
	// The username of the user to delete. Required.
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{10}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// Response body with the deleted OAuth user.
type DeleteUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a0450201a6462a0, []int{11}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*ListUsersRequest)(nil), "auth.v1beta1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "auth.v1beta1.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "auth.v1beta1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "auth.v1beta1.GetUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "auth.v1beta1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "auth.v1beta1.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "auth.v1beta1.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "auth.v1beta1.UpdateUserResponse")
	proto.RegisterType((*UpdateUserPasswordRequest)(nil), "auth.v1beta1.UpdateUserPasswordRequest")
	proto.RegisterType((*UpdateUserPasswordResponse)(nil), "auth.v1beta1.UpdateUserPasswordResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "auth.v1beta1.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "auth.v1beta1.DeleteUserResponse")
}

func init() { proto.RegisterFile("auth/v1beta1/users_api.proto", fileDescriptor_7a0450201a6462a0) }

var fileDescriptor_7a0450201a6462a0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x4f, 0xfa, 0x40,
	0x14, 0x4c, 0xe1, 0xf7, 0x07, 0x1e, 0xf8, 0xa7, 0x7b, 0x01, 0x1b, 0x54, 0xd2, 0x83, 0xf6, 0x60,
	0x4a, 0xc0, 0x83, 0x31, 0xe2, 0x01, 0x34, 0x1a, 0x13, 0x12, 0x1b, 0x12, 0x94, 0xa8, 0x89, 0x59,
	0xc2, 0x46, 0x9b, 0x28, 0xad, 0xdd, 0xa2, 0xdf, 0xc7, 0xa3, 0x1f, 0xc5, 0xcf, 0xe4, 0xc1, 0x6c,
	0xbb, 0xb4, 0xdb, 0x2d, 0x55, 0xf1, 0xb8, 0x9d, 0x99, 0x37, 0xf3, 0xf2, 0x26, 0x85, 0x1a, 0x9e,
	0xfa, 0xf7, 0x8d, 0xe7, 0xe6, 0x88, 0xf8, 0xb8, 0xd9, 0x98, 0x52, 0xe2, 0xd1, 0x5b, 0xec, 0xda,
	0xa6, 0xeb, 0x39, 0xbe, 0x83, 0xca, 0x0c, 0x35, 0x39, 0xaa, 0x55, 0x52, 0xdc, 0x90, 0xa6, 0x23,
	0x58, 0xed, 0xd9, 0xd4, 0x1f, 0x30, 0x75, 0x9f, 0x3c, 0x4d, 0x09, 0xf5, 0xf5, 0x43, 0x50, 0x85,
	0x6f, 0xd4, 0x75, 0x26, 0x94, 0x20, 0x03, 0xfe, 0x06, 0x16, 0x55, 0xa5, 0x9e, 0x37, 0x4a, 0x2d,
	0x64, 0x8a, 0xf3, 0x4d, 0xc6, 0xed, 0x87, 0x04, 0x7d, 0x07, 0x96, 0x4f, 0x49, 0xa0, 0xe6, 0x03,
	0x91, 0x06, 0x05, 0x06, 0x4d, 0xf0, 0x23, 0xa9, 0x2a, 0x75, 0xc5, 0x28, 0xf6, 0xa3, 0xb7, 0xbe,
	0x0f, 0x2b, 0x11, 0x9b, 0x5b, 0x6d, 0xc1, 0x1f, 0x06, 0x07, 0xd4, 0xf9, 0x4e, 0x01, 0xae, 0x5f,
	0x82, 0x7a, 0xe4, 0x11, 0xec, 0x13, 0xd1, 0xeb, 0x87, 0x62, 0x96, 0xc9, 0xc5, 0x94, 0xbe, 0x38,
	0xde, 0xb8, 0x9a, 0x0b, 0x33, 0xcd, 0xde, 0x7a, 0x1b, 0x90, 0x38, 0x78, 0xc1, 0x58, 0x07, 0xa0,
	0x0e, 0xdc, 0xf1, 0xef, 0x62, 0x31, 0x6b, 0x51, 0xbc, 0xa0, 0xf5, 0x1e, 0xac, 0xc5, 0x6a, 0x8b,
	0xaf, 0x23, 0x5c, 0x21, 0xda, 0x58, 0x91, 0x36, 0xae, 0x81, 0x36, 0x4f, 0x18, 0xda, 0xeb, 0x0d,
	0x50, 0x8f, 0xc9, 0x03, 0x49, 0x6e, 0xf4, 0xd5, 0x51, 0xdb, 0x80, 0x44, 0xc1, 0x62, 0x5b, 0xb4,
	0x3e, 0xf2, 0x50, 0x08, 0xca, 0xd7, 0xb1, 0xce, 0x50, 0x0f, 0x8a, 0x51, 0x19, 0xd1, 0x46, 0x52,
	0x23, 0x37, 0x57, 0xdb, 0xcc, 0xc4, 0x79, 0x84, 0x13, 0xf8, 0xcf, 0xdb, 0x86, 0x6a, 0x49, 0x6e,
	0xb2, 0xb2, 0xda, 0x7a, 0x06, 0xca, 0xe7, 0x9c, 0x03, 0xc4, 0x0d, 0x41, 0x92, 0x6d, 0xaa, 0x94,
	0x5a, 0x3d, 0x9b, 0x10, 0x0f, 0x8c, 0x0f, 0x20, 0x0f, 0x4c, 0xd5, 0x49, 0x1e, 0x38, 0xa7, 0x32,
	0x77, 0x62, 0x91, 0x66, 0x17, 0x45, 0xdb, 0x59, 0x3a, 0xa9, 0x2c, 0x9a, 0xf1, 0x3d, 0x31, 0x4e,
	0x1e, 0xdf, 0x5a, 0x4e, 0x9e, 0xaa, 0x8d, 0x9c, 0x3c, 0x5d, 0x93, 0xee, 0x35, 0x54, 0x6c, 0xc7,
	0x24, 0x13, 0x4a, 0xf0, 0x18, 0x27, 0xd8, 0xdd, 0xa5, 0xb0, 0x16, 0xae, 0x6d, 0xb1, 0x9f, 0x97,
	0xa5, 0x5c, 0x95, 0x18, 0xcc, 0xd1, 0xd7, 0x5c, 0xbe, 0x33, 0x1c, 0xbe, 0xe5, 0xca, 0x1d, 0x26,
	0xb9, 0x68, 0x76, 0xd9, 0xc7, 0xf7, 0xf0, 0x79, 0xc3, 0x9f, 0xa3, 0x7f, 0xc1, 0x6f, 0x6f, 0xf7,
	0x33, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xfb, 0xda, 0x82, 0x3d, 0x05, 0x00, 0x00,
}