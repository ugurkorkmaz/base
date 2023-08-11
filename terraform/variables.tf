variable "name" {
  type        = string
  description = "A boilerplate that implements microservices architecture on a monorepo."
  default     = "Base Project"
}
variable "namespace" {
  type        = string
  description = ""
  default     = "base_project"
}
variable "registry_url" {
  type = string
  description = ""
  default = "ghcr.io"
}

variable "registry_username" {
  type = string
  description = ""
  default = "ugurkorkmaz"
}
variable "registry_password" {
  type = string
  description = ""
}