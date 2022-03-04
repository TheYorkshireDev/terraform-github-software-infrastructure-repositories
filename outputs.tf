output "full_name" {
  value       = github_repository.software_repo.full_name
  description = "A string of the form \"orgname/reponame\""
}

output "html_url" {
  value       = github_repository.software_repo.ssh_clone_url
  description = "URL to the repository on the web"
}

output "ssh_clone_url" {
  value       = github_repository.software_repo.ssh_clone_url
  description = "URL to the repository to clone via SSH"
}

output "http_clone_url" {
  value       = github_repository.software_repo.http_clone_url
  description = "URL to clone the repository via HTTPs"
}

output "git_clone_url" {
  value       = github_repository.software_repo.git_clone_url
  description = "URL to clone the repository via the git protocol"
}

output "node_id" {
  value       = github_repository.software_repo.node_id
  description = "Node ID of the Repository"
}

output "repo_id" {
  value       = github_repository.software_repo.name
  description = "ID of the Repository"
}
