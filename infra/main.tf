# Google Cloud Provider 설정
provider "google" {
  credentials = file("key.json")  # GCP 서비스 계정 키 파일 경로
  project     = "var.project_id" # GCP 프로젝트 ID
  region      = "var.region"     # 원하는 리전
}

# Google Client Config
data "google_client_config" "default" {}

# Kubernetes Provider 설정
provider "kubernetes" {
  host                   = "https://${google_container_cluster.gdg_cluster.endpoint}"
  cluster_ca_certificate = base64decode(google_container_cluster.gdg_cluster.master_auth.0.cluster_ca_certificate)
  token                  = data.google_client_config.default.access_token
}

# GKE 클러스터 선언
resource "google_container_cluster" "gdg_cluster" {
  name     = "gdg-cluster"
  location = "var.region"
}


resource "google_container_node_pool" "default_pool" {
  cluster = google_container_cluster.gdg_cluster.name
  name    = "gdg-node-pool"

  node_config {
    machine_type = "e2-medium"
  }
}

# Kubernetes 네임스페이스 생성
resource "kubernetes_namespace" "monitoring" {
  metadata {
    name = "monitoring"
  }
}


