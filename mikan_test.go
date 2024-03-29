package mikan

import (
	"reflect"
	"testing"
)

func TestMikan_Split(t *testing.T) {
	type fields struct {
		RuneWidth int
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "runeWidth: 26",
			fields: fields{
				RuneWidth: 26,
			},
			args: args{
				str: "常に最新、最高のモバイル。Androidを開発した同じチームから。",
			},
			want: []string{
				"常に最新、最高のモバイル。", "Androidを開発した同じ", "チームから。",
			},
		},
		{
			name: "runeWidth: 25",
			fields: fields{
				RuneWidth: 25,
			},
			args: args{
				str: "常に最新、最高のモバイル。Androidを開発した同じチームから。",
			},
			want: []string{
				"常に最新、最高の", "モバイル。Androidを", "開発した同じチームから。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mikan{
				RuneWidth: tt.fields.RuneWidth,
			}
			if got := m.Split(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mikan.Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "japanese",
			args: args{
				str: "常に最新、最高のモバイル。Androidを開発した同じチームから。",
			},
			want: []string{
				"常に", "最新、", "最高の", "モバイル。", "Androidを", "開発した", "同じ", "チームから。",
			},
		},
		{
			name: "english",
			args: args{
				str: "Always the latest and best mobile. From the same team that developed Android.",
			},
			want: []string{
				"Always", " ", "the", " ", "latest", " ", "and", " ", "best", " ", "mobile.", " ", "From", " ", "the", " ", "same", " ", "team", " ", "that", " ", "developed", " ", "Android.",
			},
		},
		{
			name: "french",
			args: args{
				str: "Toujours le dernier et le meilleur mobile. De la même équipe qui a développé Android.",
			},
			want: []string{
				"Toujours", " ", "le", " ", "dernier", " ", "et", " ", "le", " ", "meilleur", " ", "mobile.", " ", "De", " ", "la", " ", "même", " ", "équipe", " ", "qui", " ", "a", " ", "développé", " ", "Android.",
			},
		},
		{
			name: "russian",
			args: args{
				str: "Всегда новейший и лучший мобильный. От той же команды, которая разработала Android.",
			},
			want: []string{
				"Всегда", " ", "новейший", " ", "и", " ", "лучший", " ", "мобильный.", " ", "От", " ", "той", " ", "же", " ", "команды,", " ", "которая", " ", "разработала", " ", "Android.",
			},
		},
		{
			name: "hankana",
			args: args{
				str: "ﾊﾛｰﾜｰﾙﾄﾞ",
			},
			want: []string{
				"ﾊﾛｰﾜｰﾙﾄﾞ",
			},
		},
		{
			name: "zenspace",
			args: args{
				str: "ハロー　ワールド",
			},
			want: []string{
				"ハロー", "　", "ワールド",
			},
		},
		{
			name: "include '・'",
			args: args{
				str: "ﾊﾛｰ・ﾜｰﾙﾄﾞ",
			},
			want: []string{
				"ﾊﾛｰ・ﾜｰﾙﾄﾞ",
			},
		},
		{
			name: "include '･'",
			args: args{
				str: "ﾊﾛｰ･ﾜｰﾙﾄﾞ",
			},
			want: []string{
				"ﾊﾛｰ･ﾜｰﾙﾄﾞ",
			},
		},
		{
			name: "include '''",
			args: args{
				str: "mitu's",
			},
			want: []string{
				"mitu's",
			},
		},
		{
			name: "include '’'",
			args: args{
				str: "mitu’s",
			},
			want: []string{
				"mitu’s",
			},
		},
		{
			name: "include '`'",
			args: args{
				str: "mitu`s",
			},
			want: []string{
				"mitu`s",
			},
		},
		{
			name: "support &＆",
			args: args{
				str: "Hello & World",
			},
			want: []string{
				"Hello", " ", "&", " ", "World",
			},
		},
		{
			name: "include -",
			args: args{
				str: "Hello-World",
			},
			want: []string{
				"Hello-World",
			},
		},
		{
			name: "include ー",
			args: args{
				str: "HelloーWorld",
			},
			want: []string{
				"HelloーWorld",
			},
		},
		{
			name: "support (株)",
			args: args{
				str: "(株)",
			},
			want: []string{
				"(株)",
			},
		},
		{
			name: "support ㈱",
			args: args{
				str: "㈱",
			},
			want: []string{
				"㈱",
			},
		},
		{
			name: "include  ́",
			args: args{
				str: "mitu´s",
			},
			want: []string{
				"mitu´s",
			},
		},
		{
			name: "support ~",
			args: args{
				str: "~test~",
			},
			want: []string{
				"~", "test~",
			},
		},
		{
			name: "support 〜",
			args: args{
				str: "〜test〜",
			},
			want: []string{
				"〜", "test", "〜",
			},
		},
		{
			name: "support another ～",
			args: args{
				str: "～test～",
			},
			want: []string{
				"～", "test", "～",
			},
		},
		{
			name: "support ѐ",
			args: args{
				str: "cafѐ",
			},
			want: []string{
				"cafѐ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Analyze(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
