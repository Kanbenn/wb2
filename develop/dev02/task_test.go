package main

import "testing"

func TestUnzipStr(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "err case два числа подряд",
			arg:     "a2b45cd",
			want:    "a2b45cd",
			wantErr: ErrIncorrectString,
		},
		{
			name:    "err case число в начале",
			arg:     "1abcd5",
			want:    "1abcd5",
			wantErr: ErrIncorrectString,
		},
		{
			name:    "case1",
			arg:     "a3bc2d5e",
			want:    "aaabccddddde",
			wantErr: nil,
		},
		{
			name:    "edge case число в конце",
			arg:     "abcd3",
			want:    "abcddd",
			wantErr: nil,
		},
		{
			name:    "case только буквы",
			arg:     "abcd",
			want:    "abcd",
			wantErr: nil,
		},
		{
			name:    "case слэш",
			arg:     `qw\\e`,
			want:    `qw\e`,
			wantErr: nil,
		},
		{
			name:    "edge case много слэшей",
			arg:     `\\\\qw\\\\\\e`,
			want:    `\\qw\\\e`,
			wantErr: nil,
		},
		{
			name:    "edge case числа",
			arg:     `\3qwe\5`,
			want:    `3qwe5`,
			wantErr: nil,
		},
		{
			name:    "edge case распаковка числа",
			arg:     `qwe\35`,
			want:    `qwe33333`,
			wantErr: nil,
		},
		{
			name:    "edge case распаковка слэша",
			arg:     `qwe\\5`,
			want:    `qwe\\\\\`,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnzipStr(tt.arg)
			if err != tt.wantErr {
				t.Errorf("UnzipStr() error = %v, want.err %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnzipStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
