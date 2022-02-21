terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "4.19.0"
    }
  }

  required_version = "~> 1.0.11"
}
