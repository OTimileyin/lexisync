package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrReadFile error = errors.New("Error: Cannot read input file")

func main() {
	inputArgs := os.Args[1:]

	inputFileName, _, err := ParseArgs(inputArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputFileContent, err := ParseInputContent(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(inputFileContent)

	inputFileContent = HexHandler(inputFileContent)
	fmt.Println(inputFileContent)

}
