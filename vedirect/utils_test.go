package vedirect

import (
	"testing"
)

func Test_FormatVoltage(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "< 10",
			args: args{
				n: 4,
			},
			want: "0.00 V",
		},
		{
			name: "< 100",
			args: args{
				n: 54,
			},
			want: "0.05 V",
		},
		{
			name: "< 1000",
			args: args{
				n: 354,
			},
			want: "0.35 V",
		},
		{
			name: "< 10000",
			args: args{
				n: 9678,
			},
			want: "9.67 V",
		},
		{
			name: "< 100000",
			args: args{
				n: 12678,
			},
			want: "12.67 V",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatVoltage(tt.args.n); got != tt.want {
				t.Errorf("formatVoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFtoa(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "main",
			args: args{
				f: 12.4565,
			},
			want: "12.45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ftoa(tt.args.f); got != tt.want {
				t.Errorf("Ftoa() = %v, want %v", got, tt.want)
			}
		})
	}
}
