package grpcserver

import (
	"io"
	"log"

	pb "github.com/kmin1231/proj_grpc/proto"
	"google.golang.org/grpc"
)

// defines 'VideoStreamingServer' struct
// embeds 'UnimplementedVideoStreamingServer' from the generated gRPC code
type VideoStreamingServer struct {
	pb.UnimplementedVideoStreamingServer
}

// defines 'StreamVideo' method for 'VideoStreamingServer' (bidirectional)
func (s *VideoStreamingServer) StreamVideo(stream pb.VideoStreaming_StreamVideoServer) error {
	// a loop to process video chunks continuously
	for {
		// receives a video chunk from client
		in, err := stream.Recv()

		// end of the stream (EOF) -> normal termination of the function
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		// logs the timestamp of the received video chunk
		log.Printf("Received video chunk with timestamp: %v", in.Timestamp)

		// sends the received video chunk back to client
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}

// defines a function to create a new gRPC server
func NewServer() *grpc.Server {
	// creates a new instance of a gRPC server
	s := grpc.NewServer()

	// registers the 'VideoStreamingServer' with the created gRPC server
	pb.RegisterVideoStreamingServer(s, &VideoStreamingServer{})

	// returns the created gRPC server
	return s
}
