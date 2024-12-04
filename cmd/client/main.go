package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Available videos:")

	// 'deployment.yaml'에 정의되어 있음
	// 환경변수 'SERVER_URL'을 읽어 서버 주소를 가져옴
	url := os.Getenv("SERVER_URL")
	if url == "" {
		log.Fatal("SERVER_URL is not set")
	}

	// '/videos' 엔드포인트로부터 Video List를 가져옴
	videos, err := listVideos(url + "/videos")
	if err != nil {
		log.Fatalf("Failed to get video list: %v", err)
	}
	for i, video := range videos {
		fmt.Printf("%d. %s\n", i+1, video)
	}

	fmt.Println("\nPlease open http://localhost:9000 in web browser!\n")

	select {}
}

func listVideos(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var videos []string
	err = json.NewDecoder(resp.Body).Decode(&videos)
	return videos, err
}

// function to open a browser
func openBrowser(url string) error {
	_, err := http.Get(url)
	return err
}