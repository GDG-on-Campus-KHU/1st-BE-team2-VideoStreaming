package video

import (
	"time"

	pb "github.com/kmin1231/proj_grpc/proto"
)

// takes video data as input -> creates a VideoChunk struct
// processes the video data & adds a timestamp
func CreateVideoChunk(data []byte) *pb.VideoChunk {
	// returns a pointer to 'VideoChunk' struct (containing video data & timestamp)
	return &pb.VideoChunk{
		Data:      data,
		Timestamp: time.Now().UnixNano(), // current time in nanoseconds
		// when the video chunk was created
	}
}
