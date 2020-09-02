package main

import (
	"fmt"
	pbgo "grpcfiletransfer/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

func download(req *pbgo.Request, responseStream pbgo.GrpcFileTransferService_DownloadServer) error {
	bufferSize := 1024 * 1024
	file, err := os.Open(req.GetFileName())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	buff := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		resp := &pbgo.Response{
			FileChunk: buff[:bytesRead],
		}
		err = responseStream.Send(resp)
		if err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
}

func unimplementedGrpcFileTransferServiceServer() {

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbgo.RegisterGrpcFileTransferServiceService(s, &pbgo.GrpcFileTransferServiceService{
		Download: download})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
