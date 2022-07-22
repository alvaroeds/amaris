package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "TOKYO"
	y := "KYOTO"

	if amigos(x, y) {
		fmt.Printf("Las cadenas '%s' y '%sy' son amigas", x, y)
	} else {
		fmt.Printf("Las cadenas '%s' y '%s' NO son amigas", x, y)
	}
}

func amigos(x, y string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 1; i < len(x); i++ {
		v := x[0:i]
		u := x[i:]
		if strings.EqualFold(y, u+v) {
			return true
		}
	}
	return false
}
