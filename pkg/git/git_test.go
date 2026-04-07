package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBehindRemoteUsesGlobalOptionsForFetch(t *testing.T) {
	tempDir := t.TempDir()
	const branch = "test-branch"

	originDir := filepath.Join(tempDir, "origin.git")
	runGit(t, tempDir, "init", "--bare", originDir)

	seedDir := filepath.Join(tempDir, "seed")
	runGit(t, tempDir, "init", seedDir)
	runGit(t, seedDir, "config", "user.name", "Test User")
	runGit(t, seedDir, "config", "user.email", "test@example.com")
	runGit(t, seedDir, "checkout", "-b", branch)
	writeFile(t, filepath.Join(seedDir, "tracked.txt"), "one\n")
	runGit(t, seedDir, "add", "tracked.txt")
	runGit(t, seedDir, "commit", "-m", "initial")
	runGit(t, seedDir, "remote", "add", "origin", originDir)
	runGit(t, seedDir, "push", "-u", "origin", branch)

	repoB := filepath.Join(tempDir, "repoB")
	runGit(t, tempDir, "clone", originDir, repoB)
	runGit(t, repoB, "checkout", branch)

	writeFile(t, filepath.Join(seedDir, "tracked.txt"), "two\n")
	runGit(t, seedDir, "add", "tracked.txt")
	runGit(t, seedDir, "commit", "-m", "ahead")
	runGit(t, seedDir, "push", "origin", branch)

	repoA := filepath.Join(tempDir, "repoA")
	runGit(t, tempDir, "init", repoA)

	originalWD, err := os.Getwd()
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NoError(t, os.Chdir(repoA)) {
		return
	}
	defer func() {
		_ = os.Chdir(originalWD)
	}()

	behind, err := IsBehindRemote("origin", branch, GlobalOptions{AsIfIn: repoB})
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, behind)
}

func TestCurrentBranchName(t *testing.T) {
	// Change this on other branches, but do not merge, to eliminate noise
	const currentBranch = "master"

	type args struct {
		globalOptions []GlobalOptions
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// Should always fail:
		{"GitDir=.", args{[]GlobalOptions{{GitDir: "."}}}, currentBranch, true},
		{"GitDir=../../", args{[]GlobalOptions{{GitDir: "../../"}}}, currentBranch, true},

		// Should always succeed when the current branch is master:
		{"no global options", args{}, currentBranch, false},
		{"empty global options", args{[]GlobalOptions{}}, currentBranch, false},
		{"AsIfIn=.", args{[]GlobalOptions{{AsIfIn: "."}}}, currentBranch, false},
		{"AsIfIn=..", args{[]GlobalOptions{{AsIfIn: ".."}}}, currentBranch, false},
		{"AsIfIn=../..", args{[]GlobalOptions{{AsIfIn: "../.."}}}, currentBranch, false},
		{"GitDir=../../.git", args{[]GlobalOptions{{GitDir: "../../.git"}}}, currentBranch, false},
		{
			"AsIfIn=.. GitDir=../.git",
			args{[]GlobalOptions{{AsIfIn: "..", GitDir: "../.git"}}},
			currentBranch,
			false,
		},

		// Exploratory tests - do not commit these uncommented, not portable:
		// {"gitDir=fully qualified ./", args{[]GlobalOptions{{GitDir: "/Users/michael.ben-david/mikebd/go-util"}}}, "master", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CurrentBranchName(tt.args.globalOptions...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if err == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()

	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v failed in %s: %v\n%s", args, dir, err, output)
	}
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
