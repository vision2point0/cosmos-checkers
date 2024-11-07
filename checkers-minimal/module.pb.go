package checkers

import (
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// GameData represents the state of a checkers game.
type GameData struct {
	state protoimpl.MessageState
	// SizeCache is not used but required for compatibility with protoimpl.
	sizeCache protoimpl.SizeCache
	// UnknownFields is not used but required for compatibility with protoimpl.
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Player1   string `protobuf:"bytes,2,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2   string `protobuf:"bytes,3,opt,name=player2,proto3" json:"player2,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

// ProtoMessage is required to satisfy the ProtoMessage interface.
func (*GameData) ProtoMessage() {}

// ProtoReflect returns the message descriptor for the message.
func (m *GameData) ProtoReflect() protoreflect.Message {
	return protoimpl.X.MessageOf(m)
}

// Reset resets the message to its zero value.
func (m *GameData) Reset() { *m = GameData{} }

// String returns a string representation of the message.
func (m *GameData) String() string {
	return protoimpl.X.MessageStringOf(m)
}

// ReqCheckersTorram is the request message for creating a checkers game.
type ReqCheckersTorram struct {
	Player1   string `protobuf:"bytes,1,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2   string `protobuf:"bytes,2,opt,name=player2,proto3" json:"player2,omitempty"`
	StartTime int64  `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

// ResCheckersTorram is the response message for the result of creating a checkers game.
type ResCheckersTorram struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

// CheckersTorram is the service definition for the Checkers module.
type CheckersTorram interface {
	CheckersCreateGm(ctx context.Context, req *ReqCheckersTorram) (*ResCheckersTorram, error)
}
