package main

import "strings"

func UpHandler(input string) string {
	wordSlice := strings.Fields(input)
//
	for idx, word := range wordSlice {
		// if word is "(up)"
		if word == "(up)" && idx > 0 {
			// check the if there is a previous word before it
			targetWord := wordSlice[idx-1]
			// transform the previous word to it uppercase form
			capitalisedTargetWord := strings.ToUpper(targetWord)
			//send it back into the slice
			wordSlice[idx-1] = capitalisedTargetWord

		}

		// append the transformed word and others excluding the tag - ("up")
		// return the joined words.
	}

}
