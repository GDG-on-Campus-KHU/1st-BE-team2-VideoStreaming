package main

import (
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/kmin1231/proj_grpc/pkg/grpcserver"
)

func main() {
	videoDir, _ := filepath.Abs("./videos")

	grpcServer := grpcserver.NewServer(videoDir)
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/videos", grpcserver.HandleVideoList(videoDir))
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		videoName := r.URL.Query().Get("video")
		log.Printf("Streaming video: %s\n", videoName)
		grpcserver.HandleVideoStream(videoDir)(w, r)
	})

	log.Println("Starting HTTP server on port 9000...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
