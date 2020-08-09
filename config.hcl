hook "deploy_openflowlabs.com" "github" {
  path = "/deploy-ofl-website"
  secret = "MyGitHubSuperSecretSecrect...?"
  event = "push"
  action = "ofl.DeployWebsite"
  params = {
    local_path = "/var/www/openflowlabs.com/public"
  }
}