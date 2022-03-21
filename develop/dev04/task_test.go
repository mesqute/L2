package main

import (
	"reflect"
	"testing"
)

func TestGetAnagramSet(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "пустой входной массив",
			args: args{[]string{}},
			want: map[string][]string{},
		},
		{
			name: "неповторяющиеся значения",
			args: args{[]string{"листок", "пятка", "столик", "тяпка"}},
			want: map[string][]string{
				"листок": {"листок", "столик"},
				"пятка":  {"пятка", "тяпка"},
				"столик": {"листок", "столик"},
				"тяпка":  {"пятка", "тяпка"},
			},
		},
		{
			name: "повторяющиеся значения",
			args: args{[]string{"листок", "пятка", "столик", "тяпка", "листок", "пятка"}},
			want: map[string][]string{
				"листок": {"листок", "столик"},
				"пятка":  {"пятка", "тяпка"},
				"столик": {"листок", "столик"},
				"тяпка":  {"пятка", "тяпка"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAnagramSet(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAnagramSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
