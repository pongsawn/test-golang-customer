package stringx

import "testing"

func TestRand(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `6`,
			args: args{
				n: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rand(tt.args.n)
			if got != tt.want {
				t.Errorf("Rand() = %v, want %v", got, tt.want)
			}
		})
	}
}
