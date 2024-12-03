terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"  # 최신 버전으로 업데이트
    }
  }
}

provider "google" {
  credentials = file("key.json")  # GCP 서비스 계정 키 경로
  project     = "gdg-side-project"  # 프로젝트 ID
  region      = "us-central1"  # 리전
}

# GKE 클러스터 정의
resource "google_container_cluster" "gke_cluster" {
  name     = "gdg-cluster"
  location = "us-central1"

  network           = "projects/gdg-side-project/global/networks/default"
  subnetwork        = "projects/gdg-side-project/regions/us-central1/subnetworks/default"

  # 변경 무시
    lifecycle {
      ignore_changes = [
        monitoring_config[0].enable_components
      ]
    }

  ip_allocation_policy {
    cluster_ipv4_cidr_block = "10.80.0.0/14"
    services_ipv4_cidr_block = "34.118.224.0/20"
  }

  logging_service    = "logging.googleapis.com/kubernetes"
  monitoring_service = "monitoring.googleapis.com/kubernetes"

  addons_config {
    gce_persistent_disk_csi_driver_config {
      enabled = true
    }
    network_policy_config {
      disabled = true
    }
  }

  release_channel {
    channel = "REGULAR"
  }

  enable_shielded_nodes = true

  logging_config {
    enable_components = [
      "SYSTEM_COMPONENTS",
      "WORKLOADS",
    ]
  }

  monitoring_config {
    enable_components = [
      "SYSTEM_COMPONENTS",
      "APISERVER",
      "CONTROLLER_MANAGER",
      "SCHEDULER",
    ]
    managed_prometheus {
      enabled = true
    }
  }
}

# GKE 노드 풀 정의
resource "google_container_node_pool" "default_pool" {
  cluster = google_container_cluster.gke_cluster.name
  name    = "default-pool"
  location = "us-central1"

  initial_node_count = 2

  node_config {
    machine_type      = "e2-small"
    disk_size_gb      = 20
    disk_type         = "pd-standard"
    image_type        = "COS_CONTAINERD"
    oauth_scopes = [
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/service.management.readonly",
      "https://www.googleapis.com/auth/servicecontrol",
      "https://www.googleapis.com/auth/trace.append",
    ]
    metadata = {
      "disable-legacy-endpoints" = "true"
    }
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }

  upgrade_settings {
    max_surge       = 1
    max_unavailable = 0
    strategy        = "SURGE"
  }
}
