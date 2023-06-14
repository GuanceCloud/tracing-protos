//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: common/Command.proto

package v3

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

// Command represents an protocol customized data when return.
//
// When the agent communicates with the OAP side using gRPC, the OAP uses Command to return the data content to the Agent.
//
// The available commands are,
// Name: ConfigurationDiscoveryCommand
// Args:
//     SerialNumber: String
//     UUID: String
//     properties: Key-value pairs rely on agent-side implementations
//
// Ref, Java agent supported configurations, https://skywalking.apache.org/docs/skywalking-java/next/en/setup/service-agent/java-agent/configuration-discovery/
//
// Name: ProfileTaskQuery
// Args:
//     SerialNumber: String
//     TaskId: String
//     EndpointName: String
//     Duration: Integer
//     MinDurationThreshold: Integer
//     DumpPeriod: Integer
//     MaxSamplingCount: Integer
//     StartTime: Date Timestamp
//     CreateTime: Date Timestamp
//
// Name: EBPFProfilingTaskQuery
// Args:
//     TaskId: String
//     ProcessId: Integer List
//     TaskUpdateTime: Date timestamp
//     TriggerType: Enum, value = FIXED_TIME
//     TargetType: Enum, value = ON_CPU, OFF_CPU or NETWORK
//     TaskStartTime: Date Timestamp
//     ExtensionConfigJSON: JSON serialization of NetworkSamplings.
//         --- NetworkSamplings ---
//         NetworkSamplings: List
//             URIRegex: String
//             MinDuration: Integer
//             When4xx: Boolean
//             When5xx: Boolean
//             Settings: Object
//                 RequireCompleteRequest: Boolean
//                 MaxRequestSize: Integer
//                 RequireCompleteResponse: Boolean
//                 MaxResponseSize: Integer
//         ------------------------
//     FixedTriggerDuration: Long
//
// Name: ContinuousProfilingPolicyQuery
// Args:
//     ServiceWithPolicyJSON: List JSON serialization of ServiceWithPolicy.
//         --- ServiceWithPolicy ---
//         ServiceName: String
//         UUID: String
//         Profiling: Multiple profiling configuration. Map
//             Key: Profiling type. Enum, value = ON_CPU, OFF_CPU, NETWORK
//             Value: Profiling policies. Map
//                 Key: Monitoring type. Enum, value = PROCESS_CPU, PROCESS_THREAD_COUNT, SYSTEM_LOAD, HTTP_ERROR_RATE, HTTP_AVG_RESPONSE_TIME
//                 Value: Policy configuration. Object.
//                     Threshold: String
//                     Period(s): Integer
//                     Count: Integer
//                     URIList: List<String>
//                     URIRegex: String
//         ---------------------------
//
// Name: ContinuousProfilingReportTask
// Args:
//     TaskId: String
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use command name to distinguish different data type.
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Data content in command.
	// The value of content needs to be serialized as string for transmission.
	//
	// Basic data type: convert as string.
	// The list of basic data: multiple data are split by ",".
	// Complex data: serialize string through json.
	Args []*KeyStringValuePair `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_common_Command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_common_Command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Command) GetArgs() []*KeyStringValuePair {
	if x != nil {
		return x.Args
	}
	return nil
}

// Transferring multiple Command in agent and OAP.
type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_Command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_common_Command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_common_Command_proto_rawDescGZIP(), []int{1}
}

func (x *Commands) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_common_Command_proto protoreflect.FileDescriptor

var file_common_Command_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x35, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x65,
	0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x3e, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x83, 0x01, 0x0a, 0x2b, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x32, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0xaa, 0x02, 0x1d, 0x53,
	0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_Command_proto_rawDescOnce sync.Once
	file_common_Command_proto_rawDescData = file_common_Command_proto_rawDesc
)

func file_common_Command_proto_rawDescGZIP() []byte {
	file_common_Command_proto_rawDescOnce.Do(func() {
		file_common_Command_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_Command_proto_rawDescData)
	})
	return file_common_Command_proto_rawDescData
}

var file_common_Command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_Command_proto_goTypes = []interface{}{
	(*Command)(nil),            // 0: skywalking.v3.Command
	(*Commands)(nil),           // 1: skywalking.v3.Commands
	(*KeyStringValuePair)(nil), // 2: skywalking.v3.KeyStringValuePair
}
var file_common_Command_proto_depIdxs = []int32{
	2, // 0: skywalking.v3.Command.args:type_name -> skywalking.v3.KeyStringValuePair
	0, // 1: skywalking.v3.Commands.commands:type_name -> skywalking.v3.Command
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_Command_proto_init() }
func file_common_Command_proto_init() {
	if File_common_Command_proto != nil {
		return
	}
	file_common_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_Command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_common_Command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
			RawDescriptor: file_common_Command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_Command_proto_goTypes,
		DependencyIndexes: file_common_Command_proto_depIdxs,
		MessageInfos:      file_common_Command_proto_msgTypes,
	}.Build()
	File_common_Command_proto = out.File
	file_common_Command_proto_rawDesc = nil
	file_common_Command_proto_goTypes = nil
	file_common_Command_proto_depIdxs = nil
}
