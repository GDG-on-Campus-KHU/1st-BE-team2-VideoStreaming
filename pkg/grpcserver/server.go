package grpcserver

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	pb "github.com/kmin1231/proj_grpc/proto"
	"google.golang.org/grpc"
)

// gRPC 서버의 동영상 스트리밍을 구현하기 위해 struct 정의
type VideoStreamingServer struct {
	pb.UnimplementedVideoStreamingServer
	VideoDir string
}

// client request에 따라 동영상 데이터를 gRPC stream을 통해 전달
func (s *VideoStreamingServer) StreamVideo(req *pb.VideoRequest, stream pb.VideoStreaming_StreamVideoServer) error {
	filePath := filepath.Join(s.VideoDir, req.VideoName)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.VideoChunk{Data: buffer[:n]}); err != nil {
			return err
		}
	}

	return nil
}

// 새로운 gRPC 서버 생성, VideoStreamingServer 등록 -> 동영상 스트리밍 기능 제공
func NewServer(videoDir string) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterVideoStreamingServer(s, &VideoStreamingServer{VideoDir: videoDir})
	return s
}

// HTTP request에 의해 동영상 목록을 return하는 handler 
func HandleVideoList(videoDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := filepath.Glob(filepath.Join(videoDir, "*.mp4"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var videoNames []string
		for _, file := range files {
			videoNames = append(videoNames, filepath.Base(file))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videoNames)
	}
}

// HTTP request에 의해 동영상을 스트리밍하는 handler
func HandleVideoStream(videoDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoName := r.URL.Query().Get("video")
		if videoName == "" {
			http.Error(w, "Video name is required", http.StatusBadRequest)
			return
		}

		filePath := filepath.Join(videoDir, videoName)
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Failed to open video file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "video/mp4")

		buffer := make([]byte, 1024*1024)
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				http.Error(w, "Error reading video file", http.StatusInternalServerError)
				return
			}
			_, err = w.Write(buffer[:n])
			if err != nil {
				// log.Printf("Error writing to response: %v", err)
				return
			}
		}
	}
}

// HTTP response를 gRPC stream처럼 처리하기 위한 struct 정의
type httpStream struct {
	w http.ResponseWriter
}

// video chunk data를 HTTP response로 전송
func (s *httpStream) Send(chunk *pb.VideoChunk) error {
	_, err := s.w.Write(chunk.Data)
	return err
}