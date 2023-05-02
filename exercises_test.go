package main

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Factorial de 5",
			args: struct{ n int64 }{n: 5},
			want: 120,
		},
		{
			name: "Factorial de 8",
			args: struct{ n int64 }{n: 8},
			want: 40320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorialRecursivo(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Factorial de 5",
			args: struct{ n int64 }{n: 8},
			want: 40320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialRecursivo(tt.args.n); got != tt.want {
				t.Errorf("FactorialRecursivo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorialWithError(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "factorial de 5",
			args:    args{5},
			want:    120,
			wantErr: false,
		},
		{
			name:    "factorial de 8",
			args:    args{8},
			want:    40320,
			wantErr: false,
		},
		{
			name:    "factorial de -5",
			args:    args{-5},
			want:    0,
			wantErr: true,
		},
		{
			name:    "factorial de 16",
			args:    args{16},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FactorialWithError(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("FactorialWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FactorialWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorialRecursiveWithError(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "factorial de 5",
			args:    args{5},
			want:    120,
			wantErr: false,
		},
		{
			name:    "factorial de 8",
			args:    args{8},
			want:    40320,
			wantErr: false,
		},
		{
			name:    "factorial de -5",
			args:    args{-5},
			want:    0,
			wantErr: true,
		},
		{
			name:    "factorial de 16",
			args:    args{16},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FactorialRecursiveWithError(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("FactorialRecursiveWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FactorialRecursiveWithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSliceBubble(t *testing.T) {
	type args struct {
		unordered []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Sort slice asc",
			args: args{[]int{-6, 12, -8, 0, 1, -6, 99}},
			want: []int64{-8, -6, -6, 0, 1, 12, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSliceBubble(tt.args.unordered); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSliceBubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
