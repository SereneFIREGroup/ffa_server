package slice

import (
	"reflect"
	"testing"
)

func TestStringSetToSlice(t *testing.T) {
	type args struct {
		set map[string]bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSetToSlice(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSetToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceToSet(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceToSet(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInterface(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInterface(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueIntSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name         string
		args         args
		wantNewSlice []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewSlice := UniqueIntSlice(tt.args.slice...); !reflect.DeepEqual(gotNewSlice, tt.wantNewSlice) {
				t.Errorf("UniqueIntSlice() = %v, want %v", gotNewSlice, tt.wantNewSlice)
			}
		})
	}
}

func TestUniqueNoNullSlice(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueNoNullStringSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueNoNullSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueSlice(t *testing.T) {
	type args[T comparable] struct {
		s []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args[string]{
				s: []string{"a", "b", "c", "a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
