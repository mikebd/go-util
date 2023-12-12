package git

import (
	"os"
	"path/filepath"
)

const gitDirBase = ".git"

// GlobalOptions are options that apply to all git commands.
// An instance of GlobalOptions may optionally be passed to each git command as the final argument.
type GlobalOptions struct {
	// GitDir is the path to the .git directory or its parent directory.
	// i.e. a repository working directory that contains a .git directory may be provided instead
	// of the .git directory.
	GitDir string
}

func (g GlobalOptions) appendGitDir(options []string) []string {
	if len(g.GitDir) == 0 {
		return options
	}

	gitDir := g.GitDir
	if filepath.Base(gitDir) != gitDirBase {
		// gitDir may be the repository working directory, that contains the .git directory
		repoDirGitDir := filepath.Join(gitDir, gitDirBase)
		fileInfo, err := os.Stat(repoDirGitDir)
		if err == nil && fileInfo.IsDir() {
			gitDir = repoDirGitDir
		}
		// otherwise assume gitDir is a .git directory
	}
	return append(options, "--git-dir="+filepath.Clean(gitDir))
}

func (g GlobalOptions) count() int {
	result := 0

	if len(g.GitDir) > 0 {
		result++
	}

	return result
}

func (g GlobalOptions) empty() bool {
	return len(g.GitDir) == 0
}

// options is a helper function that prepends any global options to the command options.
func options(globalOptions []GlobalOptions, commandOptions ...string) []string {
	if len(globalOptions) == 0 {
		return commandOptions
	}

	g := globalOptions[0]
	if g.empty() {
		return commandOptions
	}

	result := make([]string, 0, g.count()+len(commandOptions))
	result = g.appendGitDir(result)

	return append(result, commandOptions...)
}
