// Code generated by protoc-gen-go.
// source: snapshot_recovery_response.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/coreos/etcd/third_party/code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtoSnapshotRecoveryResponse struct {
	Term			*uint64	`protobuf:"varint,1,req" json:"Term,omitempty"`
	Success			*bool	`protobuf:"varint,2,req" json:"Success,omitempty"`
	CommitIndex		*uint64	`protobuf:"varint,3,req" json:"CommitIndex,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (m *ProtoSnapshotRecoveryResponse) Reset()		{ *m = ProtoSnapshotRecoveryResponse{} }
func (m *ProtoSnapshotRecoveryResponse) String() string	{ return proto.CompactTextString(m) }
func (*ProtoSnapshotRecoveryResponse) ProtoMessage()	{}

func (m *ProtoSnapshotRecoveryResponse) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *ProtoSnapshotRecoveryResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *ProtoSnapshotRecoveryResponse) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func init() {
}