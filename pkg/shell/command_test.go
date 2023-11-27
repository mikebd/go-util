package shell

import "testing"

func TestIsOutputGreaterThanZero(t *testing.T) {
	tests := []struct {
		name   string
		output []byte
		want   bool
	}{
		{
			name:   "empty output",
			output: []byte{},
			want:   false,
		},
		{
			name:   "output is empty string",
			output: []byte("\n"),
			want:   false,
		},
		{
			name:   "output is zero",
			output: []byte("0\n"),
			want:   false,
		},
		{
			name:   "output is one",
			output: []byte("1\n"),
			want:   true,
		},
		{
			name:   "output is two",
			output: []byte("2\n"),
			want:   true,
		},
		{
			name:   "output is ten",
			output: []byte("10\n"),
			want:   true,
		},
		{
			name:   "output is not a number",
			output: []byte("not a number\n"),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.output
			if got := IsOutputGreaterThanZero(output); got != tt.want {
				t.Errorf("IsOutputGreaterThanZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
