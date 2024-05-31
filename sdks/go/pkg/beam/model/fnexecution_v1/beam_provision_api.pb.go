//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Protocol Buffers describing the Provision API, for communicating with a runner
// for job and environment provisioning information over GRPC.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: org/apache/beam/model/fn_execution/v1/beam_provision_api.proto

package fnexecution_v1

import (
	pipeline_v1 "github.com/Beamdust/beam-fork/v3/go/pkg/beam/model/pipeline_v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A request to get the provision info of a SDK harness worker instance.
type GetProvisionInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProvisionInfoRequest) Reset() {
	*x = GetProvisionInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvisionInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvisionInfoRequest) ProtoMessage() {}

func (x *GetProvisionInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvisionInfoRequest.ProtoReflect.Descriptor instead.
func (*GetProvisionInfoRequest) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescGZIP(), []int{0}
}

// A response containing the provision info of a SDK harness worker instance.
type GetProvisionInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ProvisionInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetProvisionInfoResponse) Reset() {
	*x = GetProvisionInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProvisionInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProvisionInfoResponse) ProtoMessage() {}

func (x *GetProvisionInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProvisionInfoResponse.ProtoReflect.Descriptor instead.
func (*GetProvisionInfoResponse) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetProvisionInfoResponse) GetInfo() *ProvisionInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Runtime provisioning information for a SDK harness worker instance,
// such as pipeline options, resource constraints and other job metadata
type ProvisionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (required) Pipeline options. For non-template jobs, the options are
	// identical to what is passed to job submission.
	PipelineOptions *structpb.Struct `protobuf:"bytes,3,opt,name=pipeline_options,json=pipelineOptions,proto3" json:"pipeline_options,omitempty"`
	// (required) The artifact retrieval token produced by
	// LegacyArtifactStagingService.CommitManifestResponse.
	RetrievalToken string `protobuf:"bytes,6,opt,name=retrieval_token,json=retrievalToken,proto3" json:"retrieval_token,omitempty"`
	// (optional) The endpoint that the runner is hosting for the SDK to submit
	// status reports to during pipeline execution. This field will only be
	// populated if the runner supports SDK status reports. For more details see
	// https://s.apache.org/beam-fn-api-harness-status
	StatusEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,7,opt,name=status_endpoint,json=statusEndpoint,proto3" json:"status_endpoint,omitempty"`
	// (optional) The logging endpoint this SDK should use.
	LoggingEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,8,opt,name=logging_endpoint,json=loggingEndpoint,proto3" json:"logging_endpoint,omitempty"`
	// (optional) The artifact retrieval endpoint this SDK should use.
	ArtifactEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,9,opt,name=artifact_endpoint,json=artifactEndpoint,proto3" json:"artifact_endpoint,omitempty"`
	// (optional) The control endpoint this SDK should use.
	ControlEndpoint *pipeline_v1.ApiServiceDescriptor `protobuf:"bytes,10,opt,name=control_endpoint,json=controlEndpoint,proto3" json:"control_endpoint,omitempty"`
	// The set of dependencies that should be staged into this environment.
	Dependencies []*pipeline_v1.ArtifactInformation `protobuf:"bytes,11,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// (optional) A set of capabilities that this SDK is allowed to use in its
	// interactions with this runner.
	RunnerCapabilities []string `protobuf:"bytes,12,rep,name=runner_capabilities,json=runnerCapabilities,proto3" json:"runner_capabilities,omitempty"`
	// (optional) Runtime environment metadata that are static throughout the
	// pipeline execution.
	Metadata map[string]string `protobuf:"bytes,13,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// (optional) If this environment supports SIBLING_WORKERS, used to indicate
	// the ids of sibling workers, if any, that should be started in addition
	// to this worker (which already has its own worker id).
	SiblingWorkerIds []string `protobuf:"bytes,14,rep,name=sibling_worker_ids,json=siblingWorkerIds,proto3" json:"sibling_worker_ids,omitempty"`
}

func (x *ProvisionInfo) Reset() {
	*x = ProvisionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionInfo) ProtoMessage() {}

func (x *ProvisionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionInfo.ProtoReflect.Descriptor instead.
func (*ProvisionInfo) Descriptor() ([]byte, []int) {
	return file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescGZIP(), []int{2}
}

func (x *ProvisionInfo) GetPipelineOptions() *structpb.Struct {
	if x != nil {
		return x.PipelineOptions
	}
	return nil
}

func (x *ProvisionInfo) GetRetrievalToken() string {
	if x != nil {
		return x.RetrievalToken
	}
	return ""
}

func (x *ProvisionInfo) GetStatusEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if x != nil {
		return x.StatusEndpoint
	}
	return nil
}

func (x *ProvisionInfo) GetLoggingEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if x != nil {
		return x.LoggingEndpoint
	}
	return nil
}

func (x *ProvisionInfo) GetArtifactEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if x != nil {
		return x.ArtifactEndpoint
	}
	return nil
}

func (x *ProvisionInfo) GetControlEndpoint() *pipeline_v1.ApiServiceDescriptor {
	if x != nil {
		return x.ControlEndpoint
	}
	return nil
}

func (x *ProvisionInfo) GetDependencies() []*pipeline_v1.ArtifactInformation {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *ProvisionInfo) GetRunnerCapabilities() []string {
	if x != nil {
		return x.RunnerCapabilities
	}
	return nil
}

func (x *ProvisionInfo) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ProvisionInfo) GetSiblingWorkerIds() []string {
	if x != nil {
		return x.SiblingWorkerIds
	}
	return nil
}

var File_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto protoreflect.FileDescriptor

var file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x6e, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61,
	0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66, 0x6e, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x37, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x5f,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x65, 0x61,
	0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66,
	0x6e, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0xe4, 0x06, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x60, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x64, 0x0a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x62, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x5a, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x5e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x62, 0x65,
	0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66, 0x6e, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x69, 0x62, 0x6c,
	0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x73, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa8, 0x01, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66, 0x6e, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66, 0x6e, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x84, 0x01, 0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x66,
	0x6e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x5a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62,
	0x65, 0x61, 0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x6e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x3b, 0x66, 0x6e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescOnce sync.Once
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescData = file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDesc
)

func file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescGZIP() []byte {
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescOnce.Do(func() {
		file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescData)
	})
	return file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDescData
}

var file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_goTypes = []interface{}{
	(*GetProvisionInfoRequest)(nil),          // 0: org.apache.beam.model.fn_execution.v1.GetProvisionInfoRequest
	(*GetProvisionInfoResponse)(nil),         // 1: org.apache.beam.model.fn_execution.v1.GetProvisionInfoResponse
	(*ProvisionInfo)(nil),                    // 2: org.apache.beam.model.fn_execution.v1.ProvisionInfo
	nil,                                      // 3: org.apache.beam.model.fn_execution.v1.ProvisionInfo.MetadataEntry
	(*structpb.Struct)(nil),                  // 4: google.protobuf.Struct
	(*pipeline_v1.ApiServiceDescriptor)(nil), // 5: org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	(*pipeline_v1.ArtifactInformation)(nil),  // 6: org.apache.beam.model.pipeline.v1.ArtifactInformation
}
var file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_depIdxs = []int32{
	2, // 0: org.apache.beam.model.fn_execution.v1.GetProvisionInfoResponse.info:type_name -> org.apache.beam.model.fn_execution.v1.ProvisionInfo
	4, // 1: org.apache.beam.model.fn_execution.v1.ProvisionInfo.pipeline_options:type_name -> google.protobuf.Struct
	5, // 2: org.apache.beam.model.fn_execution.v1.ProvisionInfo.status_endpoint:type_name -> org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	5, // 3: org.apache.beam.model.fn_execution.v1.ProvisionInfo.logging_endpoint:type_name -> org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	5, // 4: org.apache.beam.model.fn_execution.v1.ProvisionInfo.artifact_endpoint:type_name -> org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	5, // 5: org.apache.beam.model.fn_execution.v1.ProvisionInfo.control_endpoint:type_name -> org.apache.beam.model.pipeline.v1.ApiServiceDescriptor
	6, // 6: org.apache.beam.model.fn_execution.v1.ProvisionInfo.dependencies:type_name -> org.apache.beam.model.pipeline.v1.ArtifactInformation
	3, // 7: org.apache.beam.model.fn_execution.v1.ProvisionInfo.metadata:type_name -> org.apache.beam.model.fn_execution.v1.ProvisionInfo.MetadataEntry
	0, // 8: org.apache.beam.model.fn_execution.v1.ProvisionService.GetProvisionInfo:input_type -> org.apache.beam.model.fn_execution.v1.GetProvisionInfoRequest
	1, // 9: org.apache.beam.model.fn_execution.v1.ProvisionService.GetProvisionInfo:output_type -> org.apache.beam.model.fn_execution.v1.GetProvisionInfoResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_init() }
func file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_init() {
	if File_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvisionInfoRequest); i {
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
		file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProvisionInfoResponse); i {
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
		file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionInfo); i {
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
			RawDescriptor: file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_goTypes,
		DependencyIndexes: file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_depIdxs,
		MessageInfos:      file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_msgTypes,
	}.Build()
	File_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto = out.File
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_rawDesc = nil
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_goTypes = nil
	file_org_apache_beam_model_fn_execution_v1_beam_provision_api_proto_depIdxs = nil
}
