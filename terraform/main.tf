resource "kubernetes_namespace" "base" {
    metadata {
        name = "base"
    }
}