package git

// GlobalOptions are options that apply to all git CLI commands.
// An instance of GlobalOptions may optionally be passed to each git command as the final argument.
type GlobalOptions struct {
	// AsIfIn directs git to operate as if it were run from the specified directory instead of the
	// current working directory.  i.e. the -C <path> option to the git CLI.
	AsIfIn string

	// Log directs the code in this package to log its actions.  It is not passed the git CLI.
	Log bool

	// GitDir is the path to the .git directory or its parent directory.
	// i.e. a repository working directory that contains a .git directory may be provided instead
	// of the .git directory.
	GitDir string
}

func (g GlobalOptions) appendAsIfIn(options []string) []string {
	if len(g.AsIfIn) == 0 {
		return options
	}

	return append(options, "-C", g.AsIfIn)
}

func (g GlobalOptions) appendGitDir(options []string) []string {
	if len(g.GitDir) == 0 {
		return options
	}

	return append(options, "--git-dir="+g.GitDir)
}

// count returns the number of global options that will be added to the git CLI command.
func (g GlobalOptions) count() int {
	result := 0

	if len(g.AsIfIn) > 0 {
		result += 2 // -C <path>
	}

	if len(g.GitDir) > 0 {
		result++ // --git-dir=<path>
	}

	return result
}

// empty returns true if there are no global options to be added to the git CLI command.
func (g GlobalOptions) empty() bool {
	return len(g.AsIfIn)+len(g.GitDir) == 0
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
	result = g.appendAsIfIn(result)
	result = g.appendGitDir(result)

	return append(result, commandOptions...)
}
