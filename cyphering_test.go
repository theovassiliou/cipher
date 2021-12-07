package cyphering

import (
	"testing"
)

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
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
				input:          "hallotheo",
			},
			wantCyphertext: "kdoorwkhr",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown space",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
				input:          "hallo theo",
			},
			wantCyphertext: "kdoor wkhr",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown numbers",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
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
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
				cyphertext:     "kdoorwkhr",
			},
			wantInput: "hallotheo",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown space",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
				cyphertext:     "kdoor wkhr",
			},
			wantInput: "hallo theo",
		},
		{
			name: "SimpleReverse shiftencoding with unknwown numbers",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: RotateUTF8(3, StdLowercaseAlphabet),
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
