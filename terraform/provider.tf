terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.17.0"
    }
    helm = {
      source = "hashicorp/helm"
      version = "2.10.1"
    }
  }
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
  registry {
    url = var.registry_url
    username = var.registry_username
    password = var.registry_password
  }
}