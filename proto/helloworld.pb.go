// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: proto/helloworld.proto

package proto

import (
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	WorkloadId string `protobuf:"bytes,2,opt,name=workloadId,proto3" json:"workloadId,omitempty"`
	Filter     string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *JobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobRequest) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

func (x *JobRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type JobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JobReply) Reset() {
	*x = JobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReply) ProtoMessage() {}

func (x *JobReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReply.ProtoReflect.Descriptor instead.
func (*JobReply) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *JobReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_helloworld_proto protoreflect.FileDescriptor

var file_proto_helloworld_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x0a, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xac, 0x01, 0x0a, 0x07,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0f, 0x47, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x42, 0x6c, 0x75, 0x72, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x57, 0x0a, 0x1b, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x73, 0x53,
	0x71, 0x75, 0x61, 0x64, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2d, 0x64, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_helloworld_proto_rawDescOnce sync.Once
	file_proto_helloworld_proto_rawDescData = file_proto_helloworld_proto_rawDesc
)

func file_proto_helloworld_proto_rawDescGZIP() []byte {
	file_proto_helloworld_proto_rawDescOnce.Do(func() {
		file_proto_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_helloworld_proto_rawDescData)
	})
	return file_proto_helloworld_proto_rawDescData
}

var file_proto_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: proto.HelloRequest
	(*HelloReply)(nil),   // 1: proto.HelloReply
	(*JobRequest)(nil),   // 2: proto.JobRequest
	(*JobReply)(nil),     // 3: proto.JobReply
}
var file_proto_helloworld_proto_depIdxs = []int32{
	0, // 0: proto.Greeter.SayHello:input_type -> proto.HelloRequest
	2, // 1: proto.Greeter.GrayscaleFilter:input_type -> proto.JobRequest
	2, // 2: proto.Greeter.BlurFilter:input_type -> proto.JobRequest
	1, // 3: proto.Greeter.SayHello:output_type -> proto.HelloReply
	3, // 4: proto.Greeter.GrayscaleFilter:output_type -> proto.JobReply
	3, // 5: proto.Greeter.BlurFilter:output_type -> proto.JobReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_helloworld_proto_init() }
func file_proto_helloworld_proto_init() {
	if File_proto_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_proto_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_proto_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReply); i {
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
			RawDescriptor: file_proto_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_helloworld_proto_goTypes,
		DependencyIndexes: file_proto_helloworld_proto_depIdxs,
		MessageInfos:      file_proto_helloworld_proto_msgTypes,
	}.Build()
	File_proto_helloworld_proto = out.File
	file_proto_helloworld_proto_rawDesc = nil
	file_proto_helloworld_proto_goTypes = nil
	file_proto_helloworld_proto_depIdxs = nil
}
