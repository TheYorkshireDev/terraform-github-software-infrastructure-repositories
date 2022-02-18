#tfsec:ignore:github-repositories-private
resource "github_repository" "software_repo" {
  name = var.name
}

resource "github_branch_default" "software_repo_default_branch" {
  repository = github_repository.software_repo.name
  branch     = var.default_branch
}
