package git

import (
	"github.com/mikebd/go-util/pkg/shell"
	"os/exec"
)

// CurrentBranchName returns the name of the current branch
func CurrentBranchName(globalOptions ...GlobalOptions) (string, error) {
	cmd := git(globalOptions, "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.CombinedOutput()
	if err != nil || len(output) <= 1 {
		return "", err
	}

	return string(output)[:len(output)-1], nil
}

// Fetch fetches the specified branch from the specified remote
// for the repository in the current working directory
func Fetch(remote string, branch string, globalOptions ...GlobalOptions) error {
	cmd := git(globalOptions, "fetch", remote, branch)
	return cmd.Run()
}

// IsBehindRemote returns true if the specified branch is behind the specified remote
// for the repository in the current working directory.
func IsBehindRemote(remote string, branch string, globalOptions ...GlobalOptions) (bool, error) {
	fetchErr := Fetch(remote, branch)
	if fetchErr != nil {
		return false, fetchErr
	}

	cmd := git(globalOptions, "rev-list", "--count", branch+".."+remote+"/"+branch)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return shell.IsOutputGreaterThanZero(output), nil
}

func git(globalOptions []GlobalOptions, commandOptions ...string) *exec.Cmd {
	return exec.Command("git", options(globalOptions, commandOptions...)...)
}
