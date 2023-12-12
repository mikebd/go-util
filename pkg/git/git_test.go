package git

import (
	"os"
	"testing"
)

func TestCurrentBranchName(t *testing.T) {
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
		{"gitDir=.", args{[]GlobalOptions{{GitDir: "."}}}, "master", true},

		// Should always succeed when the current branch is master:
		{"no global options", args{}, "master", false},
		{"empty global options", args{[]GlobalOptions{}}, "master", false},
		{"gitDir=../../", args{[]GlobalOptions{{GitDir: "../../"}}}, "master", false},
		{"gitDir=../../.git", args{[]GlobalOptions{{GitDir: "../../.git"}}}, "master", false},

		// Exploratory tests - do not commit these uncommented, not portable:
		// {"gitDir=fully qualified ./", args{[]GlobalOptions{{GitDir: "/Users/michael.ben-david/mikebd/go-util"}}}, "master", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CurrentBranchName(tt.args.globalOptions...)
			if (err != nil) != tt.wantErr {
				pwd, _ := os.Getwd()
				t.Errorf("CurrentBranchName() pwd = %v, error = %v, wantErr %v", pwd, err, tt.wantErr)
				return
			}
			if got != tt.want && err == nil {
				pwd, _ := os.Getwd()
				t.Errorf("CurrentBranchName() pwd = %v, got = %v, want %v", pwd, got, tt.want)
			}
		})
	}
}
