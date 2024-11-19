package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/kmin1231/proj_grpc/pkg/video"
	pb "github.com/kmin1231/proj_grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// initiates an insecure connection to connect to a local server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// creates a new gRPC client for the video streaming service
	client := pb.NewVideoStreamingClient(conn)

	// creates a bidirectional stream for video streaming
	// context.Background(): default empty context
	stream, err := client.StreamVideo(context.Background())
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	// opens the test video file
	file, err := os.Open("cmd/client/test_video.mp4")
	if err != nil {
		log.Fatalf("error opening video file: %v", err)
	}
	defer file.Close()

	// reads & sends video chunks
	// 'buffer' array -> reads data from the video file in chunks of 1024 bytes
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // terminates
		}
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}

		// processes the read data using 'video.CreateVideoChunk'
		chunk := video.CreateVideoChunk(buffer[:n])
		// sends to the server
		if err := stream.Send(chunk); err != nil {
			log.Fatalf("error sending chunk: %v", err)
		}

		// simulates real-time streaming by introducing a 50-millisecond delay
		time.Sleep(time.Millisecond * 50)
	}

	// closes send direction of the stream
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("error closing send: %v", err)
	}

	for {
		// receives video chunks from the server
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving chunk: %v", err)
		}

		// logs timestamp of each video chunk received
		log.Printf("Received video chunk with timestamp: %v", chunk.Timestamp)
	}

	// final log
	log.Println("Streaming completed!")
}
