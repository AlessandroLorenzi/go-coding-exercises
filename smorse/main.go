package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	words := []string{
		"sos",
		"daily",
		"programmer",
		"bits",
		"three",
	}
	for _, word := range words {
		fmt.Printf("smorse(\"%s\") => %s\n", word, smorse(word))
	}
}

func smorse(s string) string {
	var morse bytes.Buffer

	for _, a := range s {
		morse.WriteString(mm[string(a)])
	}
	return morse.String()
}

var mm map[string]string

func init() {
	mm = morseMap()
}

func morseMap() map[string]string {
	dotlines := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	alphabet := "a b c d e f g h i j k l m n o p q r s t u v w x y z"

	m := make(map[string]string)
	for i, a := range strings.Split(alphabet, " ") {
		m[a] = strings.Split(dotlines, " ")[i]
	}

	return m
}
