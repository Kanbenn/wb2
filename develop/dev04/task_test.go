package main

import (
	"reflect"
	"testing"
)

func TestSetOfAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *map[string][]string
	}{
		{
			name: "case1_WildBerriesL2-4",
			args: []string{"Пятак", "пЯтка", "тяпкА", "Листок", "слИток", "стОлик", "одувАн"},
			want: &map[string][]string{
				"Пятак":  {"Пятак", "пЯтка", "тяпкА"},
				"Листок": {"Листок", "слИток", "стОлик"},
			},
		},
		{
			name: "case2_LeetCode242",
			args: []string{"Anagram", "nAgaram", "cAr", "Arc", "Love"},
			want: &map[string][]string{
				"Anagram": {"Anagram", "nAgaram"},
				"cAr":     {"cAr", "Arc"},
			},
		},
		{
			name: "case3_None",
			args: []string{"Anagram", "Arc", "Love"},
			want: &map[string][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetsOfAnagrams(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOfAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
