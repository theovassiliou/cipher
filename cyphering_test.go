package cyphering

import (
	"fmt"
	"testing"
)

func TestCypherDecypher(t *testing.T) {
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
		{
			name: "Simple NewAlphabet",
			args: args{
				inputAlphabet:  StdLowercaseAlphabet,
				secretAlphabet: NewAlphabet("weiskopfseeadler", StdLowercaseAlphabet),
				input:          "wirtreffenunsum9uhr",
			},
			wantCyphertext: "vamqmkookctcntb9tfm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCyphertext := Cypher(tt.args.inputAlphabet, tt.args.secretAlphabet, tt.args.input)
			gotDecyphertext := Decypher(tt.args.inputAlphabet, tt.args.secretAlphabet, tt.wantCyphertext)
			if gotCyphertext != tt.wantCyphertext {
				t.Errorf("Cypher() = %v, want %v", gotCyphertext, tt.wantCyphertext)
			}
			if gotDecyphertext != tt.args.input {
				t.Errorf("Decypher() = %v, want %v", gotDecyphertext, tt.args.input)
				t.Errorf("Cypher() = %v, want %v", gotCyphertext, tt.wantCyphertext)

			}

		})
	}
}

func TestCypher_NonReversible(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCyphertext := Cypher(tt.args.inputAlphabet, tt.args.secretAlphabet, tt.args.input)
			if gotCyphertext != tt.wantCyphertext {
				t.Errorf("Cypher() = %v, want %v", gotCyphertext, tt.wantCyphertext)
			}
		})
	}
}

func TestNewAlphabet_Plain(t *testing.T) {
	type args struct {
		keyword      string
		baseAlphabet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				keyword:      "",
				baseAlphabet: "",
			},
			want: "",
		},
		{
			name: "empty keyword",
			args: args{
				keyword:      "",
				baseAlphabet: StdUppercaseAlphabet,
			},
			want: StdUppercaseAlphabet,
		},
		{
			name: "empty base",
			args: args{
				keyword:      StdUppercaseAlphabet,
				baseAlphabet: "",
			},
			want: StdUppercaseAlphabet,
		},
		{
			name: "KEYWORD and Alphabet",
			args: args{
				keyword:      "ASECRETKEYWORD",
				baseAlphabet: StdUppercaseAlphabet,
			},
			want: "ASECRTKYWODBFGHIJLMNPQUVXZ",
		},
		{
			name: "KEYWORD with spaced and Alphabet",
			args: args{
				keyword:      "A SECRET KEYWORD",
				baseAlphabet: StdUppercaseAlphabet,
			},
			want: "A SECRTKYWODBFGHIJLMNPQUVXZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlphabet(tt.args.keyword, tt.args.baseAlphabet); got != tt.want {
				t.Errorf("NewAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAlphabet_UT8(t *testing.T) {
	type args struct {
		keyword      string
		baseAlphabet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "KEYWORD and Alphabet",
			args: args{
				keyword:      "πανμετροναριστον",
				baseAlphabet: StdLowercaseAlphabet,
			},
			want: "πανμετροισabcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "KEYWORD and Alphabet",
			args: args{
				keyword:      "πανmetronαριστον",
				baseAlphabet: StdLowercaseAlphabet,
			},
			want: "πανmetronριστοabcdfghijklpqsuvwxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlphabet(tt.args.keyword, tt.args.baseAlphabet); got != tt.want {
				t.Errorf("NewAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCypher_helloWorld() {
	cleartext := "HELLO WORLD"
	encryptedText := Cypher(StdUppercaseAlphabet, RotateUTF8(3, StdUppercaseAlphabet), cleartext)

	fmt.Println(StdUppercaseAlphabet)
	fmt.Println(RotateUTF8(3, StdUppercaseAlphabet))
	fmt.Println(cleartext)
	fmt.Println(encryptedText)
	// Output:
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// DEFGHIJKLMNOPQRSTUVWXYZABC
	// HELLO WORLD
	// KHOOR ZRUOG
}

func ExampleDecypher_helloWorld() {
	encryptedtext := "KHOOR ZRUOG"
	cleartext := Decypher(StdUppercaseAlphabet, RotateUTF8(3, StdUppercaseAlphabet), encryptedtext)

	fmt.Println(StdUppercaseAlphabet)
	fmt.Println(RotateUTF8(3, StdUppercaseAlphabet))
	fmt.Println(encryptedtext)
	fmt.Println(cleartext)
	// Output:
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// DEFGHIJKLMNOPQRSTUVWXYZABC
	// KHOOR ZRUOG
	// HELLO WORLD
}

func ExampleCypher_dreiFragezeichen() {
	cleartext := "NYT SEITE8 HEUTE 6PM BPPUTHAUS"
	encryptedText := Cypher(StdUppercaseAlphabet, NewAlphabet("WEISKOPFSEEADLER", StdUppercaseAlphabet), cleartext)

	fmt.Println(StdUppercaseAlphabet)
	fmt.Println(NewAlphabet("WEISKOPFSEEADLER", StdUppercaseAlphabet))
	fmt.Println(cleartext)
	fmt.Println(encryptedText)
	// Output:
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// WEISKOPFADLRBCGHJMNQTUVXYZ
	// NYT SEITE8 HEUTE 6PM BPPUTHAUS
	// CYQ NKAQK8 FKTQK 6HB EHHTQFWTN
}
func ExampleDecypher_dreiFragezeichen() {
	encryptedtext := "CYQ NKAQK8 FKTQK 6HB EHHTQFWTN"
	cleartext := Decypher(StdUppercaseAlphabet, NewAlphabet("WEISKOPFSEEADLER", StdUppercaseAlphabet), encryptedtext)

	fmt.Println(StdUppercaseAlphabet)
	fmt.Println(NewAlphabet("WEISKOPFSEEADLER", StdUppercaseAlphabet))
	fmt.Println(encryptedtext)
	fmt.Println(cleartext)
	// Output:
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// WEISKOPFADLRBCGHJMNQTUVXYZ
	// CYQ NKAQK8 FKTQK 6HB EHHTQFWTN
	// NYT SEITE8 HEUTE 6PM BPPUTHAUS
}
