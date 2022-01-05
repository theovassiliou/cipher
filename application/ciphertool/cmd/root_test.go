/*
Copyright © 2021 Theo Vassiliou <vassiliou@web.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func setup() {
	rootCmd.AddCommand(cipherCmd)
	rootCmd.AddCommand(decipherCmd)
}

func buildArgs(command string, flags []string, text string) []string {
	args := []string{}
	args = append(args, command)
	args = append(args, flags...)
	args = append(args, text)
	return args
}

func TestStd(t *testing.T) {
	setup()

	type args struct {
		command  string
		flags    []string
		text     string
		decipher string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple cipher",
			args: args{
				command:  "cipher",
				flags:    []string{},
				text:     "ABCDE FGHI",
				decipher: "decipher",
			},
			want: "DEFGH IJKL",
		},
		{
			name: "simple small letters ciphering",
			args: args{
				command:  "cipher",
				flags:    []string{},
				text:     "abcde fghi",
				decipher: "",
			},
			want: "DEFGH IJKL",
		},
		{
			name: "simple small letters deciphering",
			args: args{
				command:  "decipher",
				flags:    []string{},
				text:     "defgh ijkl",
				decipher: "",
			},
			want: "ABCDE FGHI",
		},
		{
			name: "simple cipher out of alphabet --raw",
			args: args{
				command: "cipher",
				flags:   []string{"--raw"},
				text:    "abcd efgh",
			},
			want: "abcd efgh",
		},
		{
			name: "simple cipher within alphabet --raw",
			args: args{
				command: "cipher",
				flags:   []string{"--raw"},
				text:    "ABCD EFGH",
			},
			want: "DEFG HIJK",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bytes.NewBufferString("")
			rootCmd.SetOut(c)
			rootCmd.SetArgs(buildArgs(tt.args.command, tt.args.flags, tt.args.text))
			rootCmd.Execute()
			out, err := ioutil.ReadAll(c)
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tt.want {
				t.Fatalf("expected \"%s\" got \"%s\"", tt.want, string(out))
			}
		})

		if tt.args.decipher == "decipher" ||
			tt.args.decipher == "cipher" {
			t.Run(tt.name+"+"+tt.args.decipher, func(t *testing.T) {
				c := bytes.NewBufferString("")
				rootCmd.SetOut(c)
				rootCmd.SetArgs(buildArgs(tt.args.decipher, tt.args.flags, tt.want))
				rootCmd.Execute()
				out, err := ioutil.ReadAll(c)
				if err != nil {
					t.Fatal(err)
				}
				if string(out) != tt.args.text {
					t.Fatalf("expected \"%s\" got \"%s\"", tt.args.text, string(out))
				}
			})
		}
	}
}

func TestCipherRotationOption(t *testing.T) {
	setup()

	type args struct {
		command  string
		flags    []string
		text     string
		decipher string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rotation cipher",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "rotation"},
				text:     "ABCDE",
				decipher: "decipher",
			},
			want: "DEFGH",
		},
		{
			name: "rotation:5 cipher",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "rotation:5"},
				text:     "ABCDE",
				decipher: "decipher",
			},
			want: "FGHIJ",
		},
		{
			name: "rotation:-5 cipher",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "rotation:-5"},
				text:     "ABCDE",
				decipher: "decipher",
			},
			want: "VWXYZ",
		},
		{
			name: "raw rotation:5 cipher raw",
			args: args{
				command:  "cipher",
				flags:    []string{"--raw", "--cipher", "rotation:5"},
				text:     "ABCDE",
				decipher: "decipher",
			},
			want: "FGHIJ",
		},
		{
			name: "raw rotation:5 cipher raw small",
			args: args{
				command:  "cipher",
				flags:    []string{"--raw", "--cipher", "rotation:5"},
				text:     "abcde",
				decipher: "decipher",
			},
			want: "abcde",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bytes.NewBufferString("")
			rootCmd.SetOut(c)
			rootCmd.SetArgs(buildArgs(tt.args.command, tt.args.flags, tt.args.text))
			rootCmd.Execute()
			out, err := ioutil.ReadAll(c)
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tt.want {
				t.Fatalf("expected \"%s\" got \"%s\"", tt.want, string(out))
			}
		})

		if tt.args.decipher == "decipher" ||
			tt.args.decipher == "cipher" {
			t.Run(tt.name+"+"+tt.args.decipher, func(t *testing.T) {
				c := bytes.NewBufferString("")
				rootCmd.SetOut(c)
				rootCmd.SetArgs(buildArgs(tt.args.decipher, tt.args.flags, tt.want))
				rootCmd.Execute()
				out, err := ioutil.ReadAll(c)
				if err != nil {
					t.Fatal(err)
				}
				if string(out) != tt.args.text {
					t.Fatalf("expected \"%s\" got \"%s\"", tt.args.text, string(out))
				}
			})
		}
	}
}

func TestCipherReverse(t *testing.T) {
	setup()

	type args struct {
		command  string
		flags    []string
		text     string
		decipher string
	}

	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "reverse cipher",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "reverse"},
				text:     "ABCDE",
				decipher: "decipher",
			},
			want: "ZYXWV",
		},
		{
			name: "reverse cipher WS",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "reverse"},
				text:     "ABCD EFGH",
				decipher: "decipher",
			},
			want: "ZYXW VUTS",
		},
		{
			name: "reverse cipher small WS",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "reverse"},
				text:     "abcd efgh",
				decipher: "",
			},
			want: "ZYXW VUTS",
		},
		{
			name: "reverse cipher small raw WS",
			args: args{
				command:  "cipher",
				flags:    []string{"--cipher", "reverse", "--raw"},
				text:     "abcd efgh",
				decipher: "",
			},
			want: "abcd efgh",
		},
		{
			name: "reverse decipher small raw WS",
			args: args{
				command:  "decipher",
				flags:    []string{"--cipher", "reverse", "--raw"},
				text:     "abcd efgh",
				decipher: "",
			},
			want: "abcd efgh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"+"+tt.args.command, func(t *testing.T) {
			c := bytes.NewBufferString("")
			rootCmd.SetOut(c)
			rootCmd.SetArgs(buildArgs(tt.args.command, tt.args.flags, tt.args.text))
			rootCmd.Execute()
			out, err := ioutil.ReadAll(c)
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tt.want {
				t.Fatalf("expected \"%s\" got \"%s\"", tt.want, string(out))
			}
		})

		if tt.args.decipher == "decipher" ||
			tt.args.decipher == "cipher" {
			t.Run(tt.name+"+"+tt.args.decipher, func(t *testing.T) {
				c := bytes.NewBufferString("")
				rootCmd.SetOut(c)
				rootCmd.SetArgs(buildArgs(tt.args.decipher, tt.args.flags, tt.want))
				rootCmd.Execute()
				out, err := ioutil.ReadAll(c)
				if err != nil {
					t.Fatal(err)
				}
				if string(out) != tt.args.text {
					t.Fatalf("expected \"%s\" got \"%s\"", tt.args.text, string(out))
				}
			})
		}
	}
}
