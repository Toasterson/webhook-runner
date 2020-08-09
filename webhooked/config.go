package webhooked

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
	"gopkg.in/go-playground/webhooks.v5/gogs"
)

type ConfigFile struct {
	HookConfig []HookConfig `hcl:"hook,block"` // Config for the hooks
}

type HookConfig struct {
	Name   string            `hcl:"name,label"` // Human readable name of the hook
	Kind   string            `hcl:"kind,label"` // The kind of handler to use options are github, gitlab and gitea
	Path   string            `hcl:"path"`       // The path where to listen for that webhook
	Secret string            `hcl:"secret"`     // The github secret
	Event  string            `hcl:"event"`      // Event to register action under
	Action string            `hcl:"action"`     // golang action with package.Function
	Params map[string]string `hcl:"params"`     //Parameters the function will have access to
}

var supportedGitHubEventTypes = []github.Event{
	github.CheckRunEvent,
	github.CheckSuiteEvent,
	github.CommitCommentEvent,
	github.CreateEvent,
	github.DeleteEvent,
	github.DeploymentEvent,
	github.DeploymentStatusEvent,
	github.ForkEvent,
	github.GollumEvent,
	github.InstallationEvent,
	github.InstallationRepositoriesEvent,
	github.IntegrationInstallationEvent,
	github.IntegrationInstallationRepositoriesEvent,
	github.IssueCommentEvent,
	github.IssuesEvent,
	github.LabelEvent,
	github.MemberEvent,
	github.MembershipEvent,
	github.MilestoneEvent,
	github.MetaEvent,
	github.OrganizationEvent,
	github.OrgBlockEvent,
	github.PageBuildEvent,
	github.PingEvent,
	github.ProjectCardEvent,
	github.ProjectColumnEvent,
	github.ProjectEvent,
	github.PublicEvent,
	github.PullRequestEvent,
	github.PullRequestReviewEvent,
	github.PullRequestReviewCommentEvent,
	github.PushEvent,
	github.ReleaseEvent,
	github.RepositoryEvent,
	github.RepositoryVulnerabilityAlertEvent,
	github.SecurityAdvisoryEvent,
	github.StatusEvent,
	github.TeamEvent,
	github.TeamAddEvent,
	github.WatchEvent,
}

func getGitHubEventByName(name string) (github.Event, error) {
	for _, e := range supportedGitHubEventTypes {
		if string(e) == name {
			return e, nil
		}
	}

	return "", fmt.Errorf("no event named %s supported for github", name)
}

func getGitLabEventByName(name string) (gitlab.Event, error) {
	return "", fmt.Errorf("no event named %s supported for gitlab", name)
}

func getGogsEventByName(name string) (gogs.Event, error) {
	return "", fmt.Errorf("no event named %s supported for gogs", name)
}
