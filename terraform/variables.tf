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

variable "charts.dir" {
  type = string
  default = "../charts"
}

variable "registry.url" {
  type = string
  description = ""
  default = "ghcr.io"
}

variable "registry.username" {
  type = string
  description = ""
  default = "ugurkorkmaz"
}
variable "registry.password" {
  type = string
  description = ""
}