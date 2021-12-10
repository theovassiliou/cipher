package cipher

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
			name: "No shift",
			args: args{
				delta:         0,
				inputAlphabet: "ABCD",
			},
			want: "ABCD",
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

func TestStripDuplicates_Plain(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "All disjunct",
			args: args{
				input: "cafe",
			},
			want: "cafe",
		},
		{
			name: "One duplicate letter",
			args: args{
				input: "caffe",
			},
			want: "cafe",
		},
		{
			name: "Multiple duplicate letters",
			args: args{
				input: "caffee",
			},
			want: "cafe",
		},
		{
			name: "A somewhat longer text ",
			args: args{
				input: "Asomewhatlongertext",
			},
			want: "Asomewhatlngrx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripDuplicates(tt.args.input); got != tt.want {
				t.Errorf("StripDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripDuplicates_UTF8(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "All disjunct",
			args: args{
				input: "αβψδ",
			},
			want: "αβψδ",
		},
		{
			name: "One duplicate letter",
			args: args{
				input: "αβψδα",
			},
			want: "αβψδ",
		},
		{
			name: "Multiple duplicate letters",
			args: args{
				input: "αββψδδα",
			},
			want: "αβψδ",
		},
		{
			name: "Μιχεδ",
			args: args{
				input: "Πανmetronαριστον",
			},
			want: "Πανmetronριστο",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripDuplicates(tt.args.input); got != tt.want {
				t.Errorf("StripDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
