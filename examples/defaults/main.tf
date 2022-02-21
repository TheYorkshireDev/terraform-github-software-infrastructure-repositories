variable "required_name" {
  description = "Name of software repository."
  default     = "default-repo"
}

module "default" {
  source = "../../"
  name   = var.required_name
}

output "all_module_outputs" {
  value = module.default
}
