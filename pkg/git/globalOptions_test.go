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
			"AsIfIn log command",
			args{[]GlobalOptions{{AsIfIn: "some-dir"}}, []string{"log"}},
			[]string{"-C", "some-dir", "log"},
		},
		{
			"GitDir log command",
			args{[]GlobalOptions{{GitDir: "some-dir"}}, []string{"log"}},
			[]string{"--git-dir=some-dir", "log"},
		},
		{
			"AsIfIn GitDir log command",
			args{[]GlobalOptions{{AsIfIn: "some-asifin-dir", GitDir: "some-git-dir"}}, []string{"log"}},
			[]string{"-C", "some-asifin-dir", "--git-dir=some-git-dir", "log"},
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

func TestGlobalOptions_count(t *testing.T) {
	type fields struct {
		AsIfIn string
		GitDir string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"empty", fields{}, 0},
		{"AsIfIn empty", fields{AsIfIn: ""}, 0},
		{"GitDir empty", fields{GitDir: ""}, 0},
		{"AsIfIn", fields{AsIfIn: "some-dir"}, 2},
		{"GitDir", fields{GitDir: "some-dir"}, 1},
		{"AsIfIn & GitDir", fields{AsIfIn: "some-dir", GitDir: "some-dir"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GlobalOptions{
				AsIfIn: tt.fields.AsIfIn,
				GitDir: tt.fields.GitDir,
			}
			if got := g.count(); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlobalOptions_empty(t *testing.T) {
	type fields struct {
		AsIfIn string
		GitDir string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"none", fields{}, true},
		{"AsIfIn empty", fields{AsIfIn: ""}, true},
		{"GitDir empty", fields{GitDir: ""}, true},
		{"AsIfIn", fields{AsIfIn: "some-dir"}, false},
		{"GitDir", fields{GitDir: "some-dir"}, false},
		{"AsIfIn & GitDir", fields{AsIfIn: "some-dir", GitDir: "some-dir"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GlobalOptions{
				AsIfIn: tt.fields.AsIfIn,
				GitDir: tt.fields.GitDir,
			}
			if got := g.empty(); got != tt.want {
				t.Errorf("empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
