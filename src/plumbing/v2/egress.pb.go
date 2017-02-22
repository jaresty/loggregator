// Code generated by protoc-gen-go.
// source: egress.proto
// DO NOT EDIT!

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EgressRequest struct {
	ShardId string  `protobuf:"bytes,1,opt,name=shard_id,json=shardId" json:"shard_id,omitempty"`
	Filter  *Filter `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
}

func (m *EgressRequest) Reset()                    { *m = EgressRequest{} }
func (m *EgressRequest) String() string            { return proto.CompactTextString(m) }
func (*EgressRequest) ProtoMessage()               {}
func (*EgressRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type Filter struct {
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*EgressRequest)(nil), "loggregator.v2.EgressRequest")
	proto.RegisterType((*Filter)(nil), "loggregator.v2.Filter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Egress service

type EgressClient interface {
	Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Egress_serviceDesc.Streams[0], c.cc, "/loggregator.v2.Egress/Receiver", opts...)
	if err != nil {
		return nil, err
	}
	x := &egressReceiverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Egress_ReceiverClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type egressReceiverClient struct {
	grpc.ClientStream
}

func (x *egressReceiverClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Egress service

type EgressServer interface {
	Receiver(*EgressRequest, Egress_ReceiverServer) error
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Receiver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EgressServer).Receiver(m, &egressReceiverServer{stream})
}

type Egress_ReceiverServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type egressReceiverServer struct {
	grpc.ServerStream
}

func (x *egressReceiverServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receiver",
			Handler:       _Egress_Receiver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4d, 0x2f, 0x4a,
	0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xc9, 0x4f, 0x4f, 0x2f, 0x4a,
	0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x33, 0x92, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9,
	0x2f, 0x48, 0x85, 0xc8, 0x2b, 0x45, 0x71, 0xf1, 0xba, 0x82, 0xd5, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x49, 0x72, 0x71, 0x14, 0x67, 0x24, 0x16, 0xa5, 0xc4, 0x67, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x83, 0xf9, 0x9e, 0x29, 0x42, 0x7a, 0x5c, 0x6c, 0x69, 0x99,
	0x39, 0x25, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0xa8, 0x86, 0xeb,
	0xb9, 0x81, 0x65, 0x83, 0xa0, 0xaa, 0x94, 0x54, 0xb9, 0xd8, 0x20, 0x22, 0x42, 0xd2, 0x5c, 0x9c,
	0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x08, 0x53, 0x39, 0x20, 0x02, 0x9e, 0x29, 0x46, 0x81, 0x5c,
	0x6c, 0x10, 0x27, 0x08, 0xb9, 0x73, 0x71, 0x04, 0xa5, 0x26, 0xa7, 0x66, 0x96, 0xa5, 0x16, 0x09,
	0xc9, 0xa2, 0x1b, 0x8e, 0xe2, 0x4c, 0x29, 0x09, 0x0c, 0x69, 0xa8, 0xbf, 0x94, 0x18, 0x0c, 0x18,
	0x93, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xa5, 0x6a, 0x60, 0x0c,
	0x01, 0x00, 0x00,
}