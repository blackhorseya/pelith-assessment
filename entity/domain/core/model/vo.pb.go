// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: entity/domain/core/model/vo.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_entity_domain_core_model_vo_proto protoreflect.FileDescriptor

var file_entity_domain_core_model_vo_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0xf5, 0x01,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07,
	0x56, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x79, 0x61, 0x2f, 0x70, 0x65, 0x6c, 0x69, 0x74, 0x68, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0xa2, 0x02, 0x04, 0x45, 0x44, 0x43, 0x4d, 0xaa, 0x02, 0x18, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0xe2, 0x02, 0x24, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x3a, 0x3a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_entity_domain_core_model_vo_proto_goTypes = []any{}
var file_entity_domain_core_model_vo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_domain_core_model_vo_proto_init() }
func file_entity_domain_core_model_vo_proto_init() {
	if File_entity_domain_core_model_vo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_domain_core_model_vo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_domain_core_model_vo_proto_goTypes,
		DependencyIndexes: file_entity_domain_core_model_vo_proto_depIdxs,
	}.Build()
	File_entity_domain_core_model_vo_proto = out.File
	file_entity_domain_core_model_vo_proto_rawDesc = nil
	file_entity_domain_core_model_vo_proto_goTypes = nil
	file_entity_domain_core_model_vo_proto_depIdxs = nil
}
