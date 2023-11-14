package utils

import (
	"github.com/ayoubzulfiqar/Pixify/errors"
)

// Concat concat two string into one
func Concat(first, second string) (result string, err error) {
	//init
	nCopied := 0

	// get lens
	lenFirst := len(first)
	lenSecond := len(second)

	// sum of both lens
	n := lenFirst + lenSecond

	// allocate holder
	holder := make([]byte, n)

	// copy the first holder from pos nCopied:=0
	nCopied = copy(holder[nCopied:], first)

	// if everything is ok copy the second one.
	if nCopied != n-lenSecond {
		return "", errors.ErrorStat{Message: "Can't copy the first string"}
	}

	// coy the second holder from pos nCopied:= previousNCopied
	nCopied = copy(holder[nCopied:], second)

	if nCopied != n-lenFirst {
		return "", errors.ErrorStat{Message: "Can't copy the second string"}
	}

	// TODO find a better way to convert it to string
	return string(holder), nil
}
