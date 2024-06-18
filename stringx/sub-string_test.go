package stringx

import "testing"

func TestSubString(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `1`,
			args: args{
				s:      `ทดสอบ`,
				start:  0,
				length: 3,
			},
			want: `ทดส`,
		},
		{
			name: `2`,
			args: args{
				s:      `ทดสอบ`,
				start:  2,
				length: 2,
			},
			want: `สอ`,
		},
		{
			name: `3`,
			args: args{
				s:      `ทดสอบ`,
				start:  3,
				length: 10,
			},
			want: `อบ`,
		},
		{
			name: `4`,
			args: args{
				s:      `ทดสอบ`,
				start:  20,
				length: 2,
			},
			want: ``,
		},
		{
			name: `5`,
			args: args{
				s:      `ทดสอบ`,
				start:  -1,
				length: 2,
			},
			want: `ทด`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubString(tt.args.s, tt.args.start, tt.args.length); got != tt.want {
				t.Errorf("SubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
