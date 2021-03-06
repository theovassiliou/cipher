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
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theovassiliou/cipher"
	exitcodes "github.com/theovassiliou/go-exitcodes"

	"github.com/spf13/viper"
)

var cfgFile string

var defaultCipher = cipher.NewCaesarCipher(cipher.StdUppercaseAlphabet)
var defaultRottation = 3

var ciphers = map[string]cipher.CiphererDecipherer{
	"rotation": cipher.NewStdCipher(cipher.StdUppercaseAlphabet, cipher.RotateUTF8(defaultRottation, cipher.StdUppercaseAlphabet)),
	"caesar":   defaultCipher,
	"reverse":  cipher.NewStdCipher(cipher.StdUppercaseAlphabet, cipher.ReverseUTF8(cipher.StdUppercaseAlphabet)),
	"keyword":  cipher.NewKeywordCipherer("", cipher.StdUppercaseAlphabet),
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ciphertool",
		Short: "A tool to cipher and decipher text",
		Long: `chiphertool is CLI tool to cipher and decipher text using
	a set of different cipher algorithms.`,
		// Run: func(cmd *cobra.Command, args []string) {},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func initC(c *cobra.Command) {
	c.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ciphertool.yaml)")
	c.PersistentFlags().StringP("filename", "f", "", "input filename")
	c.PersistentFlags().Bool("raw", false, "do not preprocess input string or keywords")
	c.PersistentFlags().StringP("cipher", "c", "rotation:3", "name of the cipher and rotation. One of "+printAvailableCiphers())
	c.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func init() {
	cobra.OnInitialize(initConfig)
	initC(rootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ciphertool" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ciphertool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func check(e error, exitcode int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(exitcode)
	}
}

func normalize(cmd *cobra.Command, input string) string {

	raw, err := cmd.Flags().GetBool("raw")

	check(err, exitcodes.CMDLINE_USAGE_ERROR)

	if !raw {
		return strings.ToUpper(input)
	}
	return input
}

func readInputText(cmd *cobra.Command, args []string) string {
	var ct string

	filename, err := cmd.Flags().GetString("filename")
	check(err, exitcodes.CMDLINE_USAGE_ERROR)
	if filename != "" {
		dat, err := os.ReadFile(filename)
		check(err, exitcodes.CANT_OPEN_INPUT)
		ct = string(dat)
		return ct
	}

	if len(args) > 0 {
		ct = args[0]
	} else {
		r := cmd.InOrStdin()
		buf := new(strings.Builder)
		_, err := io.Copy(buf, r)
		check(err, exitcodes.CANT_OPEN_INPUT)
		ct = buf.String()
	}
	return ct

}

func getCipherer(cmd *cobra.Command, args []string) cipher.CiphererDecipherer {
	ciphernameparam, err := cmd.Flags().GetString("cipher")
	check(err, exitcodes.CMDLINE_USAGE_ERROR)
	if ciphernameparam == "" {
		return defaultCipher
	}

	var ciphername = "caesar"
	var param string

	cipherstruct := strings.Split(ciphernameparam, ":")

	switch len(cipherstruct) {
	case 1:
		ciphername = cipherstruct[0]
	case 2:
		ciphername = cipherstruct[0]
		param = cipherstruct[1]
	default:
		ciphername = ciphernameparam
	}

	theCipher := ciphers[ciphername]
	if theCipher == nil {
		fmt.Printf("Unknwon cipher: %s\n", strings.Split(ciphernameparam, ":")[0])
		os.Exit(exitcodes.CMDLINE_USAGE_ERROR)
	}

	switch ciphername {
	// Here some specialization
	case "rotation":
		if param == "" {
			param = strconv.Itoa(defaultRottation)
		}
		rot, err := strconv.Atoi(param)
		check(err, exitcodes.CMDLINE_USAGE_ERROR)
		theCipher = cipher.NewStdCipher(cipher.StdUppercaseAlphabet, cipher.RotateUTF8(rot, cipher.StdUppercaseAlphabet))
		ciphers["rotation"] = theCipher
	case "keyword":
		theCipher = cipher.NewKeywordCipherer(normalize(cmd, param), cipher.StdUppercaseAlphabet)
		ciphers["keyword"] = theCipher
	}
	return theCipher
}

func printAvailableCiphers() string {
	var s = []string{}

	for i := range ciphers {
		s = append(s, i)
	}

	s2 := fmt.Sprintf("%s", s)
	return s2
}
