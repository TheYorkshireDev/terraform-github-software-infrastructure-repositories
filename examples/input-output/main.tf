variable "input_var" {
  description = "Input variable for module."
  default     = "xxx"
}

module "input_output" {
  source    = "../../"
  input_var = var.input_var
}

output "output_var" {
  value = module.input_output.output_var
}
