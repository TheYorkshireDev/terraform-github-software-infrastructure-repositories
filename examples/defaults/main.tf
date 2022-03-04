variable "required_name" {
  description = "Name of software repository."
  default     = "default-repo"
}

locals {
  stuff   = ["a", "b", "c"]
  instuff = index(local.stuff, "d")
}

module "default" {
  source = "../../"
  name   = var.required_name
}

output "all_module_outputs" {
  value = module.default
}
