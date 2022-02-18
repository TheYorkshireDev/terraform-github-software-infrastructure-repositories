# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "name" {
  type        = string
  description = "Name of the Repository"
  default     = ""
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters match the underlying github_repository provider and defaults
# https://registry.terraform.io/providers/integrations/github/latest/docs/resources/repository
# ---------------------------------------------------------------------------------------------------------------------

variable "default_branch" {
  type        = string
  description = "Name of the Default Branch of the Repository"
  default     = "main"
}
