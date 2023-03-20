/*
Package is an encryption tool using a Caesar cipher.

Command line usage should mimic "./main.go [encrypt|decrypt] (-s, --shift #)".
Default shift should be provided. (Caesar used right shift 3)
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

/*
   TODO:
   * Sort out flags
       help[x]
       target[x]
       shift[x]

   * [x]TODO:Help documentation

   * [x]TODO Determine alphabetical shifting- probably through runes (unicode)
   []rune(string)
   string(rune)
   A:65 - z:122 - https://play.golang.org/p/2-16HLapfJk

   * [x] TODO Write to file
   * [] TODO Fix shift

*/

func leadError() {
	fmt.Println("Please provide a directive (encrypt|decrypt) and target to be transformed- string or file.\nUsage should mimic \"./main.go [encrypt/decrypt] filename (-s, --shift #)\".")
	os.Exit(1)
}

// TODO Make this work.
func transform(payload string, directive string, shift int) string {
	var result string
	var cr int
	switch directive {
	case "encrypt":
		r := []rune(payload)
		for _, i := range r {
			if unicode.IsLetter(i) {
				cr = int(i) + shift
				if cr > 122 {
					cr -= 26
				}
			} else {
				cr = int(i)
			}
			result += string(rune(cr))
		}
	case "decrypt":
		r := []rune(payload)
		for _, i := range r {
			if unicode.IsLetter(i) {
				cr = int(i) - shift
				if cr < 65 {
					cr += 26
				}
			} else {
				cr = int(i)
			}
			result += string(rune(cr))
		}
	default:
		log.Fatal("Error transforming payload")
		os.Exit(1)
	}
	return result
}

func leadHelp() {
	fmt.Println(`
dissemble
---
A basic encryption utility, utilizing a Caesar cipher. A pre-determined shift is included, though you are welcome to provide your own.

Usage:
  dissemble [command] target

Available Commands:
  encrypt    Shifts the character down the alphabet.
  decrypt    Shifts the character up the alphabet.

Flags:
  -t, --target   the content to be transformed. Can be a file or string. [Required] -- Is this still relevant?
  -s, --shift    the alphabetical value to shift. Cannot be 0.
`, "")
}

func getFileBody(target string) string {
	content, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func writeNewFile(fileName string, payload []byte, directive string) {
	var newFileName string
	ext := filepath.Ext(fileName)
	basename := strings.TrimSuffix(fileName, ext)
	switch directive {
	case "encrypt":
		newFileName = basename + "_encrypted" + ext
	case "decrypt":
		basename = strings.TrimSuffix(basename, "_encrypted")
		newFileName = basename + ext
	default:
		newFileName = basename + "_modified" + ext
	}
	fmt.Println(newFileName)
	ioutil.WriteFile(newFileName, payload, 0644)
}

func main() {
	help := flag.Bool("help", false, "")
	shift := flag.Int("shift", 3, "Alphabetical value to shift. Cannot be 0. (optional)")
	flag.Parse()

	if *help != false {
		leadHelp()
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		leadError()
	}

	switch os.Args[1] {
	case "encrypt":
		content := getFileBody(os.Args[2])
		result := transform(content, "encrypt", *shift)
		writeNewFile(os.Args[2], []byte(result), "encrypt")
		fmt.Println(*shift)
		fmt.Println(result)
	case "decrypt":
		content := getFileBody(os.Args[2])
		result := transform(content, "decrypt", *shift)
		writeNewFile(os.Args[2], []byte(result), "decrypt")
		fmt.Println(result)
	default:
		leadError()
		flag.PrintDefaults()
		os.Exit(1)
	}
}
