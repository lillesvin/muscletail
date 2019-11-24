package main

import (
	"testing"
)

func TestWindow_Full(t *testing.T) {
	tests := []struct {
		name string
		w    *Window
		want bool
	}{
		{
			name: "Full",
			w:    &Window{Length: 3, Data: []float64{1, 2, 3}},
			want: true,
		},
		{
			name: "NotFull",
			w:    &Window{Length: 3, Data: []float64{1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Full(); got != tt.want {
				t.Errorf("Window.Full() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumSlice(t *testing.T) {
	type args struct {
		slice []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "FiftyFive",
			args: args{slice: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSlice(tt.args.slice); got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMean(t *testing.T) {
	type args struct {
		slice []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive",
			args: args{slice: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 5.5,
		},
		{
			name: "Negative",
			args: args{slice: []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}},
			want: -5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMean(tt.args.slice); got != tt.want {
				t.Errorf("calculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
