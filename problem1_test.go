package main

import "testing"

// reports := []string{"SRSSRRR", "RSSRR", "SSSRRRRS", "SSRSSR"}

func TestBossBabyRevenge(t *testing.T) {
	tests := []struct {
		report string
		want   string
	}{
		{
			report: "SRSSRRR",
			want:   "Good boy",
		},
		{
			report: "RSSRR",
			want:   "Bad boy",
		},
		{
			report: "SSSRRRRS",
			want:   "Bad boy",
		},
		{
			report: "SSRR",
			want:   "Good boy",
		},
		{
			report: "SRRSSR",
			want:   "Bad boy",
		},
		{
			report: "SRRSSRR",
			want:   "Good boy",
		},
		{
			report: "SSRSRR",
			want:   "Good boy",
		},
	}
	for _, tt := range tests {
		t.Run("Problem 1: Boss Baby's Revenge", func(t *testing.T) {
			if got := BossBabyRevenge(tt.report); got != tt.want {
				t.Errorf("BossBabyRevenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
