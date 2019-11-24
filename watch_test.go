package main

import (
	"testing"
)

func TestWatch_conditionsMet(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		w    *Watch
		args args
		want bool
	}{
		{
			name: "Default_False",
			w:    &Watch{Threshold: 5},
			args: args{val: 8},
			want: false,
		},
		{
			name: "Default_True",
			w:    &Watch{Threshold: 5},
			args: args{val: 3},
			want: true,
		},
		{
			name: "LessThan_True",
			w:    &Watch{Threshold: 5, Comparison: "LessThan"},
			args: args{val: 3},
			want: true,
		},
		{
			name: "GreaterThan_False",
			w:    &Watch{Threshold: 5, Comparison: "GreaterThan"},
			args: args{val: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.conditionsMet(tt.args.val); got != tt.want {
				t.Errorf("Watch.conditionsMet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatch_stringSearch(t *testing.T) {
	type args struct {
		haystack string
	}
	tests := []struct {
		name string
		w    *Watch
		args args
		want bool
	}{
		{
			name: "Match",
			w:    &Watch{Matches: []string{"qwerty"}},
			args: args{haystack: "qwertyuiop"},
			want: true,
		},
		{
			name: "Match_UTF8",
			w:    &Watch{Matches: []string{"æ—€"}},
			args: args{haystack: "mæ—€lk"},
			want: true,
		},
		{
			name: "NoMatch",
			w:    &Watch{Matches: []string{"qwerty"}},
			args: args{haystack: "azerty"},
			want: false,
		},
		{
			name: "NoMatch_UTF8",
			w:    &Watch{Matches: []string{"k–lk"}},
			args: args{haystack: "k—lk"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.stringSearch(tt.args.haystack); got != tt.want {
				t.Errorf("Watch.stringSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
