package ofl

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/github"
	"os/exec"
)

func DeployWebsite(rawPayload interface{}, params map[string]string) error {
	payload, ok := rawPayload.(github.PushPayload)
	if !ok {
		return fmt.Errorf("gotten the wrong payload type check config expected %T got %T", github.PushPayload{}, payload)
	}

	cmd := exec.Command("/usr/bin/git", "pull")
	cmd.Dir = params["local_path"]
	return cmd.Run()
}
