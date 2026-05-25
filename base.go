package main

import (
	"strconv"
	"strings"
)
//
func HexHandler(input string) string {
	wordslice := strings.Fields(input)

	for i := 1; i < len(wordslice); i++ {
		if wordslice[i] == "(hex)" {
			hexVal := wordslice[i-1]
			decVal, _ := strconv.ParseInt(hexVal, 16, 64)
			wordslice[i-1] = strconv.Itoa(int(decVal))
			wordslice = append(wordslice[:i], wordslice[i+1:]...)
		}
	}
	return strings.Join(wordslice, " ")
}

func BinHandler(input string) string {
	wordslice := strings.Fields(input)

	for i := 1; i < len(wordslice); i++ {
		if wordslice[i] == "(bin)" {
			BinVal := wordslice[i-1]
			decVal, _ := strconv.ParseInt(BinVal, 2, 64)
			BinVal = strconv.Itoa(int(decVal))
			wordslice = append(wordslice[:i], wordslice[i+1:]...)
		}
	}
	return strings.Join(wordslice, " ")
}

func OctHandler(input string) string {
	wordslice := strings.Fields(input)

	for i := 1; i < len(wordslice); i++ {
		if wordslice[i] == "(oct)" {
			octVal := wordslice[i-1]
			decVal, _ := strconv.ParseInt(octVal, 8, 64)
			octVal = strconv.Itoa(int(decVal))
			wordslice = append(wordslice[:i], wordslice[i+1:]...)
		}
	}
	return strings.Join(wordslice, " ")
}
