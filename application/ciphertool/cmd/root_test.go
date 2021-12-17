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

func TestNewRootCmd(t *testing.T) {
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
				text:     "HELLO WORLD",
				decipher: "decipher",
			},
			want: "KHOOR ZRUOG",
		},
		{
			name: "simple small letters ciphering",
			args: args{
				command:  "cipher",
				flags:    []string{},
				text:     "hello world",
				decipher: "",
			},
			want: "KHOOR ZRUOG",
		},
		{
			name: "simple small letters deciphering",
			args: args{
				command:  "decipher",
				flags:    []string{},
				text:     "khoor zruog",
				decipher: "",
			},
			want: "HELLO WORLD",
		},
		{
			name: "simple cipher --raw",
			args: args{
				command: "cipher",
				flags:   []string{"--raw"},
				text:    "hello world",
			},
			want: "hello world",
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
