package main

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "буквенные и числовые символы",
			args:    args{"a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "только буквенные символы",
			args:    args{"abcd"},
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "только числовые символы",
			args:    args{"45"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "пустая строка",
			args:    args{""},
			want:    "",
			wantErr: false,
		},
		{
			name:    "escape-последовательность чисел",
			args:    args{`qwe\4\5`},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    "распаковка числа",
			args:    args{`qwe\45`},
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name:    `распаковка символа \`,
			args:    args{`qwe\\5`},
			want:    `qwe\\\\\`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unpack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
