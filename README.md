# gRPC를 활용한 동영상 스트리밍

- 🎬 **Project ID**: <code><b><i>1st-BE-team02-VideoStreaming</i></b></code><br>
- 👥 **Contributors**: [Eden Min Kim](https://github.com/kmin1231), [SeongHo5356](https://github.com/SeongHo5356)

 
### 프로젝트 개요

- **`gRPC`**(gRPC Remote Procedure Call)을 활용하여 로컬 환경의 동영상 파일을 업로드하고 웹 브라우저에서 재생하는 프로그램

- **`Go`** 언어로 작성

- 웹 브라우저에서 동영상 재생을 확인하기 위해 **HTML5**, **CSS**, **JavaScript** 코드 추가로 작성

- 관리 및 배포를 위해 **`Docker`**, **`Prometheus`**, **`Grafana`**, **`GKE(Google Kubernetes Engine)`**, **`Terraform`** 등을 활용

<br>

### 프로젝트 구조

```
.
├── cmd
│   ├── client
│   │   ├── Dockerfile
│   │   └── main.go
│   └── server
│       ├── Dockerfile
│       └── main.go
├── docker-compose.yaml
├── pkg
│   ├── grpcserver
│   │   └── server.go
│   └── video
│       └── video.go
├── proto
│   ├── streaming_grpc.pb.go
│   ├── streaming.pb.go
│   └── streaming.proto
├── web
│   ├── index.html
│   ├── streaming.js
├── infra
│   ├── main.tf
│   ├── outputs.tf
│   └── variables.tf
├── prometheus
│   └── prometheus.yml
├── go.mod
├── go.sum
├── deployment.yaml
└── service.yaml
```

<br>

### 개발 환경
```
$ go version
go version go1.23.1 linux/amd64

$ protoc --version
libprotoc 28.3
```

<br>

### 로컬 환경 실행 방법
[terminal #1]
```$ go run cmd/server/main.go```

[terminal #2]
```$ go run cmd/client/main.go```

<br>


### 프로그램 배포 및 관리
- **`🐳Dockerfile`**, **`🐳docker-compose.yaml`**

- **`☸️deployment.yml`**, **`☸️service.yml`** <br>
→ **`kubectl`** 명령어를 통해 **`GKE`** 에 이식 <br>
→ **`Prometheus`** 와 **`Grafana`** 를 통해 서비스 모니터링 (메트릭 데이터 수집 및 시각화) <br>

- **`Terraform`** 을 통해 서비스 관리
