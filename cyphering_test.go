package cyphering

import (
	"testing"
)

const fullAlpha = "abcdefghijklmnopqrstuvwxyz"

func TestReveerseAlphabet(t *testing.T) {
	type args struct {
		std string
	}
	tests := []struct {
		name               string
		args               args
		wantSecretAlphabet string
	}{
		{
			name: "âabcde",
			args: args{
				std: "aâbcde",
			},
			wantSecretAlphabet: "edcbâa",
		},
		{
			name: "µ€∆",
			args: args{
				std: "µ€∆",
			},
			wantSecretAlphabet: "∆€µ",
		},
		{
			name: "StandardAlphabet",
			args: args{
				std: StdLowercaseAlphabet,
			},
			wantSecretAlphabet: "zyxwvutsrqponmlkjihgfedcba",
		},
		{
			name: "Hellenic",
			args: args{
				std: "άσδφγηξκλέ",
			},
			wantSecretAlphabet: "έλκξηγφδσά",
		},
		{
			name: "Empty",
			args: args{
				std: "",
			},
			wantSecretAlphabet: "",
		},
		{
			name: "One Element",
			args: args{
				std: "a",
			},
			wantSecretAlphabet: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSecretAlphabet := ReverseUTF8(tt.args.std); gotSecretAlphabet != tt.wantSecretAlphabet {
				t.Errorf("ReveerseAlphabet() = %v, want %v", gotSecretAlphabet, tt.wantSecretAlphabet)
			}
		})
	}
}

func TestCypher(t *testing.T) {
	type args struct {
		inputAlphabet  string
		secretAlphabet string
		input          string
	}
	tests := []struct {
		name           string
		args           args
		wantCyphertext string
	}{

		{
			name: "SimpleReverse",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: ReverseUTF8(StdLowercaseAlphabet),
				input:          "abcde",
			},
			wantCyphertext: "zyxwv",
		},
		{
			name: "SimpleReverse unkown Char",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: ReverseUTF8(StdLowercaseAlphabet),
				input:          "abc de",
			},
			wantCyphertext: "zyx wv",
		},
		{
			name: "SimpleReverse unkown Char",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: ReverseUTF8(StdLowercaseAlphabet),
				input:          "abcAde",
			},
			wantCyphertext: "zyxAwv",
		},
		{
			name: "SimpleReverse Unicodes",
			args: args{
				inputAlphabet:  "άσδφγηξκλέ",
				secretAlphabet: "έλκξηγφδσά",
				input:          "άσδ",
			},
			wantCyphertext: "έλκ",
		},
		{
			name: "SimpleReverse Mixed",
			args: args{
				inputAlphabet:  "άσδφγηξκλέabcd",
				secretAlphabet: "dcbaέλκξηγφδσά",
				input:          "άcd",
			},
			wantCyphertext: "dσά",
		},
		{
			name: "SimpleReverse empty input",
			args: args{
				inputAlphabet:  "abcd",
				secretAlphabet: "zyxw",
				input:          "",
			},
			wantCyphertext: "",
		},
		{
			name: "SimpleReverse One Element",
			args: args{
				inputAlphabet:  "abcd",
				secretAlphabet: "zyxw",
				input:          "a",
			},
			wantCyphertext: "z",
		},
		{
			name: "SimpleReverse Empty inputAlphabet",
			args: args{
				inputAlphabet:  "",
				secretAlphabet: "zyxw",
				input:          "a",
			},
			wantCyphertext: "",
		},
		{
			name: "SimpleReverse Empty secretAlphabet",
			args: args{
				inputAlphabet:  "abcd",
				secretAlphabet: "",
				input:          "a",
			},
			wantCyphertext: "",
		},
		{
			name: "SimpleReverse Empty secretAlphabet",
			args: args{
				inputAlphabet:  "abcd",
				secretAlphabet: "b",
				input:          "z",
			},
			wantCyphertext: "z",
		},
		// const fullAlpha = "abcdefghijklmnopqrstuvwxyz"
		// const shift3    = "defghijklmnopqrstuvwxyzabc"

		{
			name: "SimpleReverse shiftencoding",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				input:          "hallotheo",
			},
			wantCyphertext: "kdoorwkhr",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown space",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				input:          "hallo theo",
			},
			wantCyphertext: "kdoor wkhr",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown numbers",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				input:          "hallo78theo90",
			},
			wantCyphertext: "kdoor78wkhr90",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCyphertext := Cypher(tt.args.inputAlphabet, tt.args.secretAlphabet, tt.args.input); gotCyphertext != tt.wantCyphertext {
				t.Errorf("Cypher() = %v, want %v", gotCyphertext, tt.wantCyphertext)
			}
		})
	}
}

func TestDecypher(t *testing.T) {
	type args struct {
		inputAlphabet  string
		secretAlphabet string
		cyphertext     string
	}
	tests := []struct {
		name      string
		args      args
		wantInput string
	}{
		{
			name: "SimpleReverse",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: ReverseUTF8(StdLowercaseAlphabet),
				cyphertext:     "zyxwv",
			},
			wantInput: "abcde",
		},
		{
			name: "SimpleReverse Unicodes",
			args: args{
				inputAlphabet:  "άσδφγηξκλέ",
				secretAlphabet: "έλκξηγφδσά",
				cyphertext:     "έλκ",
			},
			wantInput: "άσδ",
		},
		{
			name: "SimpleReverse Mixed",
			args: args{
				inputAlphabet:  "άσδφγηξκλέabcd",
				secretAlphabet: "dcbaέλκξηγφδσά",
				cyphertext:     "dσά",
			},
			wantInput: "άcd",
		},
		// const fullAlpha = "abcdefghijklmnopqrstuvwxyz"
		// const shift3    = "defghijklmnopqrstuvwxyzabc"

		{
			name: "SimpleReverse shiftencoding",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				cyphertext:     "kdoorwkhr",
			},
			wantInput: "hallotheo",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown space",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				cyphertext:     "kdoor wkhr",
			},
			wantInput: "hallo theo",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown numbers",
			args: args{
				inputAlphabet:  fullAlpha,
				secretAlphabet: ShiftUTF8(3, fullAlpha),
				cyphertext:     "kdoor78wkhr90",
			},
			wantInput: "hallo78theo90",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := Decypher(tt.args.inputAlphabet, tt.args.secretAlphabet, tt.args.cyphertext); gotInput != tt.wantInput {
				t.Errorf("Decypher() = %v, want %v", gotInput, tt.wantInput)
			}
		})
	}
}

func TestShiftAlphabet(t *testing.T) {
	type args struct {
		delta         int
		inputAlphabet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No shift",
			args: args{
				delta:         0,
				inputAlphabet: "ABCD",
			},
			want: "ABCD",
		},
		{
			name: "ROT-1",
			args: args{
				delta:         1,
				inputAlphabet: "ABCD",
			},
			want: "BCDA",
		},
		{
			name: "ROT-2",
			args: args{
				delta:         2,
				inputAlphabet: "ABCD",
			},
			want: "CDAB",
		},
		{
			name: "ROT-1-UTF8",
			args: args{
				delta:         1,
				inputAlphabet: "αβγδ",
			},
			want: "βγδα",
		},
		{
			name: "ROT-2-UTF8",
			args: args{
				delta:         2,
				inputAlphabet: "αβγδ",
			},
			want: "γδαβ",
		},
		{
			name: "ROT-2-ΜΙΧ",
			args: args{
				delta:         2,
				inputAlphabet: "abγδ",
			},
			want: "γδab",
		},
		{
			name: "ROT--1",
			args: args{
				delta:         -1,
				inputAlphabet: "ABCDE",
			},
			want: "EABCD",
		},
		{
			name: "ROT--2",
			args: args{
				delta:         -2,
				inputAlphabet: "ABCDE",
			},
			want: "DEABC",
		},
		{
			name: "ROT--2 Mixed",
			args: args{
				delta:         -2,
				inputAlphabet: "AβCδE",
			},
			want: "δEAβC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiftUTF8(tt.args.delta, tt.args.inputAlphabet); got != tt.want {
				t.Errorf("ShiftAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRunes(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple",
			args: args{
				in: "abcde",
			},
			want: 5,
		},
		{
			name: "Single",
			args: args{
				in: "a",
			},
			want: 1,
		},
		{
			name: "Empty",
			args: args{
				in: "",
			},
			want: 0,
		},
		{
			name: "Single UTF-8",
			args: args{
				in: "α",
			},
			want: 1,
		},
		{
			name: "Multi UTF-8",
			args: args{
				in: "αβψδ",
			},
			want: 4,
		},

		{
			name: "Mixed",
			args: args{
				in: "aβcδeφ",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRunes(tt.args.in); got != tt.want {
				t.Errorf("countRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
