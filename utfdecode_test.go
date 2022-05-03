package utfdecode

import "testing"

func Test_(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "unicodeエスケープシーケンスなし",
			input: "サンプル",
			want:  "サンプル",
		},
		{
			name:  "サロゲートペア不要のunicodeエスケープシーケンスのみ",
			input: `\u3042\u3044\u3046\u3048\u304a`,
			want:  "あいうえお",
		},
		{
			name:  "サロゲートペア不要のunicodeエスケープシーケンスとエスケープされていない文字が混ざっている",
			input: `\u3042あ\u3044い\u3046う\u3048え\u304aお`,
			want:  "ああいいううええおお",
		},
		{
			name:  "サロゲートペアが必要なunicodeエスケープシーケンスのみ",
			input: `\uD83D\uDE04\uD83D\uDE07\uD83D\uDC7A`,
			want:  "😄😇👺",
		},
		{
			name:  "サロゲートペアが必要なunicodeエスケープシーケンスとエスケープされていない文字が混ざっている",
			input: `\uD83D\uDE04あ\uD83D\uDE07い\uD83D\uDC7Aう`,
			want:  "😄あ😇い👺う",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Decode(tt.input)
			if err != nil {
				t.Fatalf("failed to execute Decode: %v", err)
			}
			if got != tt.want {
				t.Fatalf("want: %v, but got %v", tt.want, got)
			}
		})
	}

}
