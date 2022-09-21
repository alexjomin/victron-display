package effect

import (
	"reflect"
	"testing"
)

// 8X8
var m8x8 = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0}

// 11110000
// 11110000
// 11110000
// 11110000
// 11110000
// 11110000
// 11110000
// 11110000

var m16x8 = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 2, 2, 2, 2, 2, 2, 2, 2}

// 11111111-00000000
// 11111111-11111111
// 11111111-00000000
// 11111111-00000000
// 11111111-00000000
// 11111111-00000000
// 11111111-00000000
// 11111111-00000000

var m8x16 = []byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var m8x32 = []byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xAA, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xAA,
}

// 11111111
// 11111111
// 11111111
// 11111111
// 11111111
// 11111111
// 11111111
// 11111111
// 00000000
// 00000000
// 00000000
// 00000000
// 00000000
// 00000000
// 00000000
// 00000000

func TestSlide(t *testing.T) {
	type args struct {
		buffer []byte
		offset int
		w      int
		h      int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "simple test",
			args: args{
				buffer: m8x8,
				offset: 1,
				w:      8,
				h:      8,
			},
			want: []byte{0, 0xFF, 0xFF, 0xFF, 0xFF, 0, 0, 0},
		},
		{
			name: "16x8 test",
			args: args{
				buffer: m16x8,
				offset: 1,
				w:      16,
				h:      8,
			},
			want: []byte{0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			name: "8x16 test",
			args: args{
				buffer: m8x16,
				offset: 1,
				w:      8,
				h:      16,
			},
			want: []byte{
				0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			name: "8x32",
			args: args{
				buffer: m8x32,
				offset: 1,
				w:      8,
				h:      32,
			},
			want: []byte{
				0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // 0
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 8
				0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // 16
				0x00, 0xAA, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 24
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Translate(tt.args.buffer, tt.args.offset, tt.args.w, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slide() = %v, want %v", got, tt.want)
			}
		})
	}
}
