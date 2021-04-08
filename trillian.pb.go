// Copyright 2016 Google LLC. All Rights Reserved.
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
// 	protoc        v3.12.4
// source: trillian.proto

package trillian

import (
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	keyspb "github.com/google/trillian/crypto/keyspb"
	sigpb "github.com/google/trillian/crypto/sigpb"
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

// LogRootFormat specifies the fields that are covered by the
// SignedLogRoot signature, as well as their ordering and formats.
type LogRootFormat int32

const (
	LogRootFormat_LOG_ROOT_FORMAT_UNKNOWN LogRootFormat = 0
	LogRootFormat_LOG_ROOT_FORMAT_V1      LogRootFormat = 1
)

// Enum value maps for LogRootFormat.
var (
	LogRootFormat_name = map[int32]string{
		0: "LOG_ROOT_FORMAT_UNKNOWN",
		1: "LOG_ROOT_FORMAT_V1",
	}
	LogRootFormat_value = map[string]int32{
		"LOG_ROOT_FORMAT_UNKNOWN": 0,
		"LOG_ROOT_FORMAT_V1":      1,
	}
)

func (x LogRootFormat) Enum() *LogRootFormat {
	p := new(LogRootFormat)
	*p = x
	return p
}

func (x LogRootFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRootFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_trillian_proto_enumTypes[0].Descriptor()
}

func (LogRootFormat) Type() protoreflect.EnumType {
	return &file_trillian_proto_enumTypes[0]
}

func (x LogRootFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRootFormat.Descriptor instead.
func (LogRootFormat) EnumDescriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{0}
}

// Defines the way empty / node / leaf hashes are constructed incorporating
// preimage protection, which can be application specific.
type HashStrategy int32

const (
	// Hash strategy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	// Certificate Transparency strategy: leaf hash prefix = 0x00, node prefix =
	// 0x01, empty hash is digest([]byte{}), as defined in the specification.
	HashStrategy_RFC6962_SHA256 HashStrategy = 1
	// Sparse Merkle Tree strategy:  leaf hash prefix = 0x00, node prefix = 0x01,
	// empty branch is recursively computed from empty leaf nodes.
	// NOT secure in a multi tree environment. For testing only.
	HashStrategy_TEST_MAP_HASHER HashStrategy = 2
	// Append-only log strategy where leaf nodes are defined as the ObjectHash.
	// All other properties are equal to RFC6962_SHA256.
	HashStrategy_OBJECT_RFC6962_SHA256 HashStrategy = 3
	// The CONIKS sparse tree hasher with SHA512_256 as the hash algorithm.
	HashStrategy_CONIKS_SHA512_256 HashStrategy = 4
	// The CONIKS sparse tree hasher with SHA256 as the hash algorithm.
	HashStrategy_CONIKS_SHA256 HashStrategy = 5
)

// Enum value maps for HashStrategy.
var (
	HashStrategy_name = map[int32]string{
		0: "UNKNOWN_HASH_STRATEGY",
		1: "RFC6962_SHA256",
		2: "TEST_MAP_HASHER",
		3: "OBJECT_RFC6962_SHA256",
		4: "CONIKS_SHA512_256",
		5: "CONIKS_SHA256",
	}
	HashStrategy_value = map[string]int32{
		"UNKNOWN_HASH_STRATEGY": 0,
		"RFC6962_SHA256":        1,
		"TEST_MAP_HASHER":       2,
		"OBJECT_RFC6962_SHA256": 3,
		"CONIKS_SHA512_256":     4,
		"CONIKS_SHA256":         5,
	}
)

func (x HashStrategy) Enum() *HashStrategy {
	p := new(HashStrategy)
	*p = x
	return p
}

func (x HashStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_trillian_proto_enumTypes[1].Descriptor()
}

func (HashStrategy) Type() protoreflect.EnumType {
	return &file_trillian_proto_enumTypes[1]
}

func (x HashStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashStrategy.Descriptor instead.
func (HashStrategy) EnumDescriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{1}
}

// State of the tree.
type TreeState int32

const (
	// Tree state cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	// Active trees are able to respond to both read and write requests.
	TreeState_ACTIVE TreeState = 1
	// Frozen trees are only able to respond to read requests, writing to a frozen
	// tree is forbidden. Trees should not be frozen when there are entries
	// in the queue that have not yet been integrated. See the DRAINING
	// state for this case.
	TreeState_FROZEN TreeState = 2
	// Deprecated: now tracked in Tree.deleted.
	//
	// Deprecated: Do not use.
	TreeState_DEPRECATED_SOFT_DELETED TreeState = 3
	// Deprecated: now tracked in Tree.deleted.
	//
	// Deprecated: Do not use.
	TreeState_DEPRECATED_HARD_DELETED TreeState = 4
	// A tree that is draining will continue to integrate queued entries.
	// No new entries should be accepted.
	TreeState_DRAINING TreeState = 5
)

// Enum value maps for TreeState.
var (
	TreeState_name = map[int32]string{
		0: "UNKNOWN_TREE_STATE",
		1: "ACTIVE",
		2: "FROZEN",
		3: "DEPRECATED_SOFT_DELETED",
		4: "DEPRECATED_HARD_DELETED",
		5: "DRAINING",
	}
	TreeState_value = map[string]int32{
		"UNKNOWN_TREE_STATE":      0,
		"ACTIVE":                  1,
		"FROZEN":                  2,
		"DEPRECATED_SOFT_DELETED": 3,
		"DEPRECATED_HARD_DELETED": 4,
		"DRAINING":                5,
	}
)

func (x TreeState) Enum() *TreeState {
	p := new(TreeState)
	*p = x
	return p
}

func (x TreeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TreeState) Descriptor() protoreflect.EnumDescriptor {
	return file_trillian_proto_enumTypes[2].Descriptor()
}

func (TreeState) Type() protoreflect.EnumType {
	return &file_trillian_proto_enumTypes[2]
}

func (x TreeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TreeState.Descriptor instead.
func (TreeState) EnumDescriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{2}
}

// Type of the tree.
type TreeType int32

const (
	// Tree type cannot be determined. Included to enable detection of mismatched
	// proto versions being used. Represents an invalid value.
	TreeType_UNKNOWN_TREE_TYPE TreeType = 0
	// Tree represents a verifiable log.
	TreeType_LOG TreeType = 1
	// Tree represents a verifiable pre-ordered log, i.e., a log whose entries are
	// placed according to sequence numbers assigned outside of Trillian.
	TreeType_PREORDERED_LOG TreeType = 3
)

// Enum value maps for TreeType.
var (
	TreeType_name = map[int32]string{
		0: "UNKNOWN_TREE_TYPE",
		1: "LOG",
		3: "PREORDERED_LOG",
	}
	TreeType_value = map[string]int32{
		"UNKNOWN_TREE_TYPE": 0,
		"LOG":               1,
		"PREORDERED_LOG":    3,
	}
)

func (x TreeType) Enum() *TreeType {
	p := new(TreeType)
	*p = x
	return p
}

func (x TreeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TreeType) Descriptor() protoreflect.EnumDescriptor {
	return file_trillian_proto_enumTypes[3].Descriptor()
}

func (TreeType) Type() protoreflect.EnumType {
	return &file_trillian_proto_enumTypes[3]
}

func (x TreeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TreeType.Descriptor instead.
func (TreeType) EnumDescriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{3}
}

// Represents a tree.
// Readonly attributes are assigned at tree creation, after which they may not
// be modified.
//
// Note: Many APIs within the rest of the code require these objects to
// be provided. For safety they should be obtained via Admin API calls and
// not created dynamically.
type Tree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the tree.
	// Readonly.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	// State of the tree.
	// Trees are ACTIVE after creation. At any point the tree may transition
	// between ACTIVE, DRAINING and FROZEN states.
	TreeState TreeState `protobuf:"varint,2,opt,name=tree_state,json=treeState,proto3,enum=trillian.TreeState" json:"tree_state,omitempty"`
	// Type of the tree.
	// Readonly after Tree creation. Exception: Can be switched from
	// PREORDERED_LOG to LOG if the Tree is and remains in the FROZEN state.
	TreeType TreeType `protobuf:"varint,3,opt,name=tree_type,json=treeType,proto3,enum=trillian.TreeType" json:"tree_type,omitempty"`
	// Hash strategy to be used by the tree.
	// Readonly.
	HashStrategy HashStrategy `protobuf:"varint,4,opt,name=hash_strategy,json=hashStrategy,proto3,enum=trillian.HashStrategy" json:"hash_strategy,omitempty"`
	// Hash algorithm to be used by the tree.
	// Readonly.
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,5,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// Signature algorithm to be used by the tree.
	// Readonly.
	SignatureAlgorithm sigpb.DigitallySigned_SignatureAlgorithm `protobuf:"varint,6,opt,name=signature_algorithm,json=signatureAlgorithm,proto3,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Display name of the tree.
	// Optional.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description of the tree,
	// Optional.
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	// Identifies the private key used for signing tree heads and entry
	// timestamps.
	// This can be any type of message to accommodate different key management
	// systems, e.g. PEM files, HSMs, etc.
	// Private keys are write-only: they're never returned by RPCs.
	// The private_key message can be changed after a tree is created, but the
	// underlying key must remain the same - this is to enable migrating a key
	// from one provider to another.
	PrivateKey *any.Any `protobuf:"bytes,12,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Storage-specific settings.
	// Varies according to the storage implementation backing Trillian.
	StorageSettings *any.Any `protobuf:"bytes,13,opt,name=storage_settings,json=storageSettings,proto3" json:"storage_settings,omitempty"`
	// The public key used for verifying tree heads and entry timestamps.
	// Readonly.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,14,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Interval after which a new signed root is produced even if there have been
	// no submission.  If zero, this behavior is disabled.
	MaxRootDuration *duration.Duration `protobuf:"bytes,15,opt,name=max_root_duration,json=maxRootDuration,proto3" json:"max_root_duration,omitempty"`
	// Time of tree creation.
	// Readonly.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,16,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time of last tree update.
	// Readonly (automatically assigned on updates).
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,17,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// If true, the tree has been deleted.
	// Deleted trees may be undeleted during a certain time window, after which
	// they're permanently deleted (and unrecoverable).
	// Readonly.
	Deleted bool `protobuf:"varint,19,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Time of tree deletion, if any.
	// Readonly.
	DeleteTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func (x *Tree) Reset() {
	*x = Tree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trillian_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tree) ProtoMessage() {}

func (x *Tree) ProtoReflect() protoreflect.Message {
	mi := &file_trillian_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tree.ProtoReflect.Descriptor instead.
func (*Tree) Descriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{0}
}

func (x *Tree) GetTreeId() int64 {
	if x != nil {
		return x.TreeId
	}
	return 0
}

func (x *Tree) GetTreeState() TreeState {
	if x != nil {
		return x.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (x *Tree) GetTreeType() TreeType {
	if x != nil {
		return x.TreeType
	}
	return TreeType_UNKNOWN_TREE_TYPE
}

func (x *Tree) GetHashStrategy() HashStrategy {
	if x != nil {
		return x.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (x *Tree) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if x != nil {
		return x.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (x *Tree) GetSignatureAlgorithm() sigpb.DigitallySigned_SignatureAlgorithm {
	if x != nil {
		return x.SignatureAlgorithm
	}
	return sigpb.DigitallySigned_ANONYMOUS
}

func (x *Tree) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Tree) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tree) GetPrivateKey() *any.Any {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *Tree) GetStorageSettings() *any.Any {
	if x != nil {
		return x.StorageSettings
	}
	return nil
}

func (x *Tree) GetPublicKey() *keyspb.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Tree) GetMaxRootDuration() *duration.Duration {
	if x != nil {
		return x.MaxRootDuration
	}
	return nil
}

func (x *Tree) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tree) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Tree) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Tree) GetDeleteTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

// SignedLogRoot represents a commitment by a Log to a particular tree.
type SignedLogRoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// log_root holds the TLS-serialization of the following structure (described
	// in RFC5246 notation): Clients should validate log_root_signature with
	// VerifySignedLogRoot before deserializing log_root.
	// enum { v1(1), (65535)} Version;
	// struct {
	//   uint64 tree_size;
	//   opaque root_hash<0..128>;
	//   uint64 timestamp_nanos;
	//   uint64 revision;
	//   opaque metadata<0..65535>;
	// } LogRootV1;
	// struct {
	//   Version version;
	//   select(version) {
	//     case v1: LogRootV1;
	//   }
	// } LogRoot;
	//
	// A serialized v1 log root will therefore be laid out as:
	//
	// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+-....--+
	// | ver=1 |          tree_size            |len|    root_hash      |
	// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+-....--+
	//
	// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	// |        timestamp_nanos        |      revision                 |
	// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// +---+---+---+---+---+-....---+
	// |  len  |    metadata        |
	// +---+---+---+---+---+-....---+
	//
	// (with all integers encoded big-endian).
	LogRoot []byte `protobuf:"bytes,8,opt,name=log_root,json=logRoot,proto3" json:"log_root,omitempty"`
	// log_root_signature is the raw signature over log_root.
	LogRootSignature []byte `protobuf:"bytes,9,opt,name=log_root_signature,json=logRootSignature,proto3" json:"log_root_signature,omitempty"`
}

func (x *SignedLogRoot) Reset() {
	*x = SignedLogRoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trillian_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedLogRoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedLogRoot) ProtoMessage() {}

func (x *SignedLogRoot) ProtoReflect() protoreflect.Message {
	mi := &file_trillian_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedLogRoot.ProtoReflect.Descriptor instead.
func (*SignedLogRoot) Descriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{1}
}

func (x *SignedLogRoot) GetLogRoot() []byte {
	if x != nil {
		return x.LogRoot
	}
	return nil
}

func (x *SignedLogRoot) GetLogRootSignature() []byte {
	if x != nil {
		return x.LogRootSignature
	}
	return nil
}

// Proof holds a consistency or inclusion proof for a Merkle tree, as returned
// by the API.
type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// leaf_index indicates the requested leaf index when this message is used for
	// a leaf inclusion proof.  This field is set to zero when this message is
	// used for a consistency proof.
	LeafIndex int64    `protobuf:"varint,1,opt,name=leaf_index,json=leafIndex,proto3" json:"leaf_index,omitempty"`
	Hashes    [][]byte `protobuf:"bytes,3,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trillian_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_trillian_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_trillian_proto_rawDescGZIP(), []int{2}
}

func (x *Proof) GetLeafIndex() int64 {
	if x != nil {
		return x.LeafIndex
	}
	return 0
}

func (x *Proof) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

var File_trillian_proto protoreflect.FileDescriptor

var file_trillian_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x1a, 0x1a, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x62, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x73,
	0x69, 0x67, 0x70, 0x62, 0x2f, 0x73, 0x69, 0x67, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x07, 0x0a,
	0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x65, 0x65, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x0a, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2e, 0x54, 0x72,
	0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x72, 0x65, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e,
	0x2e, 0x54, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x72, 0x65, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69,
	0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x4b, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x69, 0x67, 0x70, 0x62,
	0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x0d,
	0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x5a, 0x0a,
	0x13, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x69, 0x67,
	0x70, 0x62, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f,
	0x6d, 0x61, 0x78, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0b,
	0x10, 0x0c, 0x4a, 0x04, 0x08, 0x12, 0x10, 0x13, 0x52, 0x1e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x10, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x16, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69,
	0x74, 0x65, 0x52, 0x1e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x6f, 0x67,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6c, 0x6f, 0x67,
	0x52, 0x6f, 0x6f, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x08, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x64, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x52, 0x0d, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x74, 0x72, 0x65,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x50, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x65, 0x61, 0x66, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2a, 0x44, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x52,
	0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x47,
	0x5f, 0x52, 0x4f, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x52, 0x4f,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x56, 0x31, 0x10, 0x01, 0x2a, 0x97,
	0x01, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x68, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x19, 0x0a, 0x15, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x5f,
	0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x46,
	0x43, 0x36, 0x39, 0x36, 0x32, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x46,
	0x43, 0x36, 0x39, 0x36, 0x32, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x03, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x49, 0x4b, 0x53, 0x5f, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x5f,
	0x32, 0x35, 0x36, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x49, 0x4b, 0x53, 0x5f,
	0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x05, 0x2a, 0x8b, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x65,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52,
	0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x17, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43,
	0x41, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x4f, 0x46, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x1f, 0x0a, 0x17, 0x44, 0x45, 0x50, 0x52, 0x45,
	0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x52, 0x41, 0x49,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x2a, 0x49, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x52,
	0x45, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x47,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x45, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x45, 0x44,
	0x5f, 0x4c, 0x4f, 0x47, 0x10, 0x03, 0x22, 0x04, 0x08, 0x02, 0x10, 0x02, 0x2a, 0x03, 0x4d, 0x41,
	0x50, 0x42, 0x48, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0d,
	0x54, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_trillian_proto_rawDescOnce sync.Once
	file_trillian_proto_rawDescData = file_trillian_proto_rawDesc
)

func file_trillian_proto_rawDescGZIP() []byte {
	file_trillian_proto_rawDescOnce.Do(func() {
		file_trillian_proto_rawDescData = protoimpl.X.CompressGZIP(file_trillian_proto_rawDescData)
	})
	return file_trillian_proto_rawDescData
}

var file_trillian_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_trillian_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trillian_proto_goTypes = []interface{}{
	(LogRootFormat)(0),                       // 0: trillian.LogRootFormat
	(HashStrategy)(0),                        // 1: trillian.HashStrategy
	(TreeState)(0),                           // 2: trillian.TreeState
	(TreeType)(0),                            // 3: trillian.TreeType
	(*Tree)(nil),                             // 4: trillian.Tree
	(*SignedLogRoot)(nil),                    // 5: trillian.SignedLogRoot
	(*Proof)(nil),                            // 6: trillian.Proof
	(sigpb.DigitallySigned_HashAlgorithm)(0), // 7: sigpb.DigitallySigned.HashAlgorithm
	(sigpb.DigitallySigned_SignatureAlgorithm)(0), // 8: sigpb.DigitallySigned.SignatureAlgorithm
	(*any.Any)(nil),             // 9: google.protobuf.Any
	(*keyspb.PublicKey)(nil),    // 10: keyspb.PublicKey
	(*duration.Duration)(nil),   // 11: google.protobuf.Duration
	(*timestamp.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_trillian_proto_depIdxs = []int32{
	2,  // 0: trillian.Tree.tree_state:type_name -> trillian.TreeState
	3,  // 1: trillian.Tree.tree_type:type_name -> trillian.TreeType
	1,  // 2: trillian.Tree.hash_strategy:type_name -> trillian.HashStrategy
	7,  // 3: trillian.Tree.hash_algorithm:type_name -> sigpb.DigitallySigned.HashAlgorithm
	8,  // 4: trillian.Tree.signature_algorithm:type_name -> sigpb.DigitallySigned.SignatureAlgorithm
	9,  // 5: trillian.Tree.private_key:type_name -> google.protobuf.Any
	9,  // 6: trillian.Tree.storage_settings:type_name -> google.protobuf.Any
	10, // 7: trillian.Tree.public_key:type_name -> keyspb.PublicKey
	11, // 8: trillian.Tree.max_root_duration:type_name -> google.protobuf.Duration
	12, // 9: trillian.Tree.create_time:type_name -> google.protobuf.Timestamp
	12, // 10: trillian.Tree.update_time:type_name -> google.protobuf.Timestamp
	12, // 11: trillian.Tree.delete_time:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_trillian_proto_init() }
func file_trillian_proto_init() {
	if File_trillian_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trillian_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tree); i {
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
		file_trillian_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedLogRoot); i {
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
		file_trillian_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
			RawDescriptor: file_trillian_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trillian_proto_goTypes,
		DependencyIndexes: file_trillian_proto_depIdxs,
		EnumInfos:         file_trillian_proto_enumTypes,
		MessageInfos:      file_trillian_proto_msgTypes,
	}.Build()
	File_trillian_proto = out.File
	file_trillian_proto_rawDesc = nil
	file_trillian_proto_goTypes = nil
	file_trillian_proto_depIdxs = nil
}
