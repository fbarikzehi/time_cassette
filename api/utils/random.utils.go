package utils

import "math/rand"

var colorChars = []rune("abcdef0123456789")

func RandomColorString(length int) string {
	color := make([]rune, length)
	for i := range color {
		color[i] = colorChars[rand.Intn(len(colorChars))]
	}
	return string(color)
}

var numericChars = []rune("0123456789")

func NumericRandString(length int) string {
	holder := make([]rune, length)
	for i := range holder {
		holder[i] = numericChars[rand.Intn(len(numericChars))]
	}
	return string(holder)
}

var alphanumericChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomAlphanumericString(length int) string {
	alphanumeric := make([]rune, length)
	for i := range alphanumeric {
		alphanumeric[i] = alphanumericChars[rand.Intn(len(alphanumericChars))]
	}
	return string(alphanumeric)
}
