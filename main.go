package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	shift      = 36
	expansion  = 0x2c0b02c1
	printable  = 0x21
	multiplier = 0x5d
	filler     = 0x41
	usage      = "Usage:\n go-keygen -login=NAME\nCalculates a password to the given login name for the passcheck application"
)

func main() {
	// stores the generated password
	var password string

	helpPtr := flag.Bool("help", false, "display help")
	hPtr := flag.Bool("h", false, "display help")
	var login string
	flag.StringVar(&login, "login", "", "the login name")
	flag.Parse()
	if *hPtr || *helpPtr {
		fmt.Println(usage)
		os.Exit(0)
	}
	if login == "" {
		fmt.Println(usage)
		os.Exit(1)
	}

	counter := 0
	// at most 8 password characters are generated that way
	for counter < 8 && counter < len(login) {
		if counter == 0 {
			// for the first password character use the first two characters of the username
			x := keygen(int(login[counter]), int(login[counter+1]))
			password += string(x)
			counter++
		} else {
			// all other characters are genereated by using the last character from the password, and a character from the login
			x := keygen((int(login[counter]) << 2), int(password[counter-1]))
			password += string(x)
			counter++
		}
	}
	// the remaining password characters are generated up to the final length of 10
	for 10-counter > 0 {
		x := filler + counter
		y := int(password[counter-1])
		z := 10 - counter
		b := keygen(y, (x << uint(z)))
		password += string(b)
		counter++
	}
	fmt.Println(password)

}

// derived function from the passcheck application
func keygen(l int, r int) (sym int) {
	x := l + r
	sym = printable + x - ((x*expansion)>>shift)*multiplier
	return
}
