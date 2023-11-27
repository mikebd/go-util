package git

import "os/exec"

func Fetch(remote string, branch string) error {
	cmd := exec.Command("git", "fetch", remote, branch)
	return cmd.Run()
}
