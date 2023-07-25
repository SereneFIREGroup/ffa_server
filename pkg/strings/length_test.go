package strings

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{str: ""}, want: true},
		{args: args{str: "12"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.str); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLenValid(t *testing.T) {
	type args struct {
		str    string
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{str: "", length: 0}, want: true},
		{args: args{str: "1", length: 2}, want: true},
		{args: args{str: "123", length: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLenValid(tt.args.str, tt.args.length); got != tt.want {
				t.Errorf("IsLenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLenValidUTF8(t *testing.T) {
	type args struct {
		str    string
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{str: "", length: 0}, want: true},
		{args: args{str: "12", length: 2}, want: true},
		{args: args{str: "12", length: 1}, want: false},
		{args: args{str: "我的", length: 2}, want: true},
		{args: args{str: "我的", length: 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLenValidUTF8(tt.args.str, tt.args.length); got != tt.want {
				t.Errorf("IsLenValidUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLenUTF8(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{str: ""}, want: 0},
		{args: args{str: "12"}, want: 2},
		{args: args{str: "我的"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenUTF8(tt.args.str); got != tt.want {
				t.Errorf("LenUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
