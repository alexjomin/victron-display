package tools

import "testing"

func TestWaterProbe(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "main",
			args: args{
				v: 2.0,
			},
			want: 22.1588,
		},
		{
			name: "main",
			args: args{
				v: 19.0,
			},
			want: 197.4067,
		},
		{
			name: "main",
			args: args{
				v: 0.0,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WaterProbe(tt.args.v); got != tt.want {
				t.Errorf("WaterProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}
