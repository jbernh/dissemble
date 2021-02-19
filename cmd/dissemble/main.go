/*
Package dissemble is a basic encryption tool using a classic Caesar cipher.

Command line usage should mimic "./main.go [encrypt|decrypt] (-s, --shift #)".
Default shift should be provided. (Caesar used right shift 3)
*/
package main

import (
    "fmt"
    "flag"
    "os"
)

/*
    TODO:
    * Determine packages-
        fmt
        flags
        os
        (something to shift the runes)
    
    * Sort out flags
        help[] TODO
        target[]
        shift[x]
    
    * TODO:Help documentation
    
    * TODO Determine alphabetical shifting- probably through runes (unicode)
    []rune(string)
    string(rune)
    a:97 - z:122

*/

func leadError() {
    fmt.Println("Please provide a directive (encrypt|decrypt) and target to be transformed- string or file.\nUsage should mimic \"./main.go [encrypt/decrypt] filename (-s, --shift #)\".")
    os.Exit(1)
}

//TODO Make this work.
func transform(payload string, directive string, shift int) string {
    switch directive {
    case "encrypt":
        return "encrypted"
    case "decrypt":
        return "decrypted"
    default:
        return "Nothing"
    }
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
  -t, --target   the content to be transformed. Can be a file or string. [Required]
  -s, --shift    the alphabetical value to shift. Cannot be 0.
`, "")
}

func main() {
    help := flag.Bool("help", false, "")
    target := flag.String("target", "", "Content to be transformed.") 
    shift := flag.Int("shift", 3, "Alphabetical value to shift. Cannot be 0. (optional)")
    flag.Parse()
    
    if *help != false {
        leadHelp()
        os.Exit(0)
    }

//    if *target == "" {
//        leadError()
//        flag.PrintDefaults()
//        os.Exit(1)
//    }

    if len(os.Args) < 2 {
       leadError() 
    }

    switch os.Args[1] {
    case "encrypt":
        transform(os.Args[2], "encrypt", *shift)
        fmt.Printf("This would encrypt %v.", target)
    case "decrypt":
        transform(os.Args[2], "decrypt", *shift)
        fmt.Printf("This would decrypt %v.", target)
    default:
        leadError()
        flag.PrintDefaults()
        os.Exit(1)
    }
    fmt.Printf("%v", flag.Args())
} 
