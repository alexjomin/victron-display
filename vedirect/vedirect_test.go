package vedirect

import (
	_ "embed"
	"testing"
)

//go:embed testdata/bluesolar_1.23.dump
var fulldump []byte

func TestParser_ParseByte(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "main",
			args: args{
				b: fulldump,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewParser()

			if err != nil {
				t.Fatalf("Parser.NewParser() error = %v", err)
			}

			s, _ := NewState()

			for _, b := range tt.args.b {
				if p, err = p.ParseByte(b); (err != nil) != tt.wantErr {
					if err == ErrCheckSumNotValid {
						t.Log(err)
						continue
					}
					t.Errorf("Parser.ParseByte() error = %v, wantErr %v", err, tt.wantErr)
				}
				if p.Ready {
					f, err := NewFrame(p.KV)
					if err != nil {
						t.Errorf("NewFrame error = %v", err)
					}
					s = s.Update(f)
					t.Logf("%+v", f)
				}
			}
		})
	}
}
