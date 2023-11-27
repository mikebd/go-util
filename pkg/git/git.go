package git

import (
	"github.com/mikebd/go-util/pkg/shell"
	"os/exec"
)

// Fetch fetches the specified branch from the specified remote
// for the repository in the current working directory
func Fetch(remote string, branch string) error {
	cmd := exec.Command("git", "fetch", remote, branch)
	return cmd.Run()
}

// IsBehindRemote returns true if the specified branch is behind the specified remote
// for the repository in the current working directory.
func IsBehindRemote(remote string, branch string) (bool, error) {
	fetchErr := Fetch(remote, branch)
	if fetchErr != nil {
		return false, fetchErr
	}

	cmd := exec.Command("git", "rev-list", "--count", branch+".."+remote+"/"+branch)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return shell.IsOutputGreaterThanZero(output), nil
}
