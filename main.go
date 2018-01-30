package main

import (
	"fmt"
	"os"
)

const (
	shift      = 36
	expansion  = 0x2c0b02c1
	printable  = 0x21
	multiplier = 0x5d
	filler     = 0x41
	usage      = "Usage:\n go-keygen login\nCalculates a password to the given login name for the passcheck application"
)

var password string

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		fmt.Println("Please supply only a login value")
		os.Exit(1)
	}
	login := arguments[0]
	if login == "-h" || login == "--help" {
		fmt.Println(usage)
		os.Exit(0)
	}
	counter := 0
	for counter < 8 && counter < len(login) {
		if counter == 0 {
			x := keygen(int(login[counter]), int(login[counter+1]))
			password += string(x)
			counter++
		} else {
			x := keygen((int(login[counter]) << 2), int(password[counter-1]))
			password += string(x)
			counter++
		}
	}
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
func keygen(l int, r int) (sym int) {
	x := l + r
	sym = printable + x - ((x*expansion)>>shift)*multiplier
	return
}
