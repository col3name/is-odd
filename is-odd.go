package main

import "errors"

var UnsupportedError = errors.New("unsupported error")

func IsOdd(number int64) (bool, error) {
	if number%2 == 0 {
		return true, nil
	} else if number%2 != 0 {
		return false, nil
	}
	return false, UnsupportedError
}
