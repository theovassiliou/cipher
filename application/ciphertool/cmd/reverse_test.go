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
			testCmd := setup()
			c := bytes.NewBufferString("")
			testCmd.SetOut(c)
			testCmd.SetArgs(buildArgs(tt.args.command, tt.args.flags, tt.args.text))
			testCmd.Execute()
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
				testCmd := setup()
				c := bytes.NewBufferString("")
				testCmd.SetOut(c)
				testCmd.SetArgs(buildArgs(tt.args.decipher, tt.args.flags, tt.want))
				testCmd.Execute()
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
