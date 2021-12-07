package main

import (
	"testing"
)

func Test_formatString(t *testing.T) {
		tests := []struct {
				in  string
				out string
		}{
				{"無変更", "無変更"},
				{"半角空白 削除", "半角空白削除"},
				{"全角空白　削除", "全角空白削除"},
				{`改行
削除`, "改行削除"},
		}

		for _, tt := range tests {
				if got := formatString(tt.in); got != tt.out {
						t.Errorf("formatString(%v) = %v, want %v", tt.in, got, tt.out)
				}
		}
}
