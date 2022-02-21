variable "name" {
  type        = string
  description = "The name of the software repository (e.g. hello-world). This variable is used in all repositories related to project."
}

variable "default_branch" {
  type        = string
  description = "The default branch of the repository."
  default     = "main"
}
