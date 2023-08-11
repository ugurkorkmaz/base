variable "name" {
  type        = string
  description = "A boilerplate that implements microservices architecture on a monorepo."
  default     = "Base Project"
}
variable "namespace" {
  type        = string
  default     = "base_project"
}
variable "registry_url" {
  type = string
  default = "ghcr.io"
}

variable "registry_username" {
  type = string
  default = "ugurkorkmaz"
}
variable "registry_password" {
  type = string
}