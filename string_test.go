package cyphering

import (
	"testing"
)

func TestReverseUTF8_Plain(t *testing.T) {
	type args struct {
		std string
	}
	tests := []struct {
		name               string
		args               args
		wantSecretAlphabet string
	}{
		{
			name: "StandardAlphabet",
			args: args{
				std: StdLowercaseAlphabet,
			},
			wantSecretAlphabet: "zyxwvutsrqponmlkjihgfedcba",
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
				t.Errorf("ReverseUTF8() = %v, want %v", gotSecretAlphabet, tt.wantSecretAlphabet)
			}
		})
	}
}

func TestReverseUTF8_UTF8(t *testing.T) {
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
			name: "Hellenic",
			args: args{
				std: "άσδφγηξκλέ",
			},
			wantSecretAlphabet: "έλκξηγφδσά",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSecretAlphabet := ReverseUTF8(tt.args.std); gotSecretAlphabet != tt.wantSecretAlphabet {
				t.Errorf("ReverseUTF8() = %v, want %v", gotSecretAlphabet, tt.wantSecretAlphabet)
			}
		})
	}
}
func TestRotateUTF8(t *testing.T) {
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
			if got := RotateUTF8(tt.args.delta, tt.args.inputAlphabet); got != tt.want {
				t.Errorf("RotateUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotateUTF8Plain(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateUTF8(tt.args.delta, tt.args.inputAlphabet); got != tt.want {
				t.Errorf("RotateUTF8() = %v, want %v", got, tt.want)
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
