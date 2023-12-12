package git

import (
	"reflect"
	"testing"
)

func Test_options(t *testing.T) {
	type args struct {
		globalOptions  []GlobalOptions
		commandOptions []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"empty global options, empty command options",
			args{[]GlobalOptions{}, []string{}},
			[]string{},
		},
		{
			"empty global options, log command",
			args{[]GlobalOptions{}, []string{"log"}},
			[]string{"log"},
		},
		{
			"empty global options, log -p command",
			args{[]GlobalOptions{}, []string{"log", "-p"}},
			[]string{"log", "-p"},
		},
		{
			"gitDir global option, log command",
			args{[]GlobalOptions{{GitDir: "some-dir"}}, []string{"log"}},
			[]string{"--git-dir=some-dir", "log"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := options(tt.args.globalOptions, tt.args.commandOptions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options() = %v, want %v", got, tt.want)
			}
		})
	}
}
