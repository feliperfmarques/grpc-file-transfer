// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GrpcFileTransferServiceClient is the client API for GrpcFileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcFileTransferServiceClient interface {
	Download(ctx context.Context, in *Request, opts ...grpc.CallOption) (GrpcFileTransferService_DownloadClient, error)
}

type grpcFileTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcFileTransferServiceClient(cc grpc.ClientConnInterface) GrpcFileTransferServiceClient {
	return &grpcFileTransferServiceClient{cc}
}

var grpcFileTransferServiceDownloadStreamDesc = &grpc.StreamDesc{
	StreamName:    "Download",
	ServerStreams: true,
}

func (c *grpcFileTransferServiceClient) Download(ctx context.Context, in *Request, opts ...grpc.CallOption) (GrpcFileTransferService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, grpcFileTransferServiceDownloadStreamDesc, "/grpcfiletransfer.GrpcFileTransferService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcFileTransferServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcFileTransferService_DownloadClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type grpcFileTransferServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *grpcFileTransferServiceDownloadClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GrpcFileTransferServiceService is the service API for GrpcFileTransferService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterGrpcFileTransferServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type GrpcFileTransferServiceService struct {
	Download func(*Request, GrpcFileTransferService_DownloadServer) error
}

func (s *GrpcFileTransferServiceService) download(_ interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.Download(m, &grpcFileTransferServiceDownloadServer{stream})
}

type GrpcFileTransferService_DownloadServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type grpcFileTransferServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *grpcFileTransferServiceDownloadServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterGrpcFileTransferServiceService registers a service implementation with a gRPC server.
func RegisterGrpcFileTransferServiceService(s grpc.ServiceRegistrar, srv *GrpcFileTransferServiceService) {
	srvCopy := *srv
	if srvCopy.Download == nil {
		srvCopy.Download = func(*Request, GrpcFileTransferService_DownloadServer) error {
			return status.Errorf(codes.Unimplemented, "method Download not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpcfiletransfer.GrpcFileTransferService",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Download",
				Handler:       srvCopy.download,
				ServerStreams: true,
			},
		},
		Metadata: "proto/grpc_file_transfer.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewGrpcFileTransferServiceService creates a new GrpcFileTransferServiceService containing the
// implemented methods of the GrpcFileTransferService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewGrpcFileTransferServiceService(s interface{}) *GrpcFileTransferServiceService {
	ns := &GrpcFileTransferServiceService{}
	if h, ok := s.(interface {
		Download(*Request, GrpcFileTransferService_DownloadServer) error
	}); ok {
		ns.Download = h.Download
	}
	return ns
}

// UnstableGrpcFileTransferServiceService is the service API for GrpcFileTransferService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableGrpcFileTransferServiceService interface {
	Download(*Request, GrpcFileTransferService_DownloadServer) error
}
