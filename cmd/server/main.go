package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kmin1231/proj_grpc/pkg/grpcserver"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


// Prometheus Counter Initialization
// Prometheus 메트릭을 초기화하여 HTTP 요청을 추적할 수 있도록 함
var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

// Prometheus Metrics Middleware
// 서버 시작 시, Prometheus 메트릭 등록(register)
func init() {
	prometheus.MustRegister(httpRequests)
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		next.ServeHTTP(w, r)
	})
}

// File Upload Handler
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filePath := filepath.Join("videos", header.Filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}


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

	http.Handle("/", http.FileServer(http.Dir("/app/web")))
	// Ensure the videos directory exists
	if _, err := os.Stat("videos"); os.IsNotExist(err) {
		os.Mkdir("videos", os.ModePerm)
	}

	// HTTP Request handlers
	//http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/videos", grpcserver.HandleVideoList(videoDir))
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		videoName := r.URL.Query().Get("video")
		log.Printf("Streaming video: %s\n", videoName)
		grpcserver.HandleVideoStream(videoDir)(w, r)
	})

	// Prometheus /metrics 엔드포인트
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Starting HTTP server on port 9000...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
