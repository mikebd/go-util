package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
