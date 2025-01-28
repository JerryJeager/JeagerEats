// code generate required for photographer signup
package utils

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lettersArr = strings.Split(letters, "")

func sng() string { //single character(num or alphabet)
	d := rand.IntN(2)
	var c string
	if d == 1 {
		c = string(lettersArr[rand.IntN(len(lettersArr))])
	} else {
		c = string('1' + rand.IntN(9))
	}
	return c
}

func GenerateCode() string {
	return fmt.Sprintf("%s%s%s%s%s", sng(), sng(), sng(), sng(), sng())
}
