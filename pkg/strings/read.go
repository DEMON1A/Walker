package strings

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isPrintableAscii(c byte) bool {
	return c >= 32 && c <= 126
}

func ReadStringsFromFile(path string, max_length int) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var strings []string
	var currentString []byte
	for _, b := range data {
		if isPrintableAscii(b) {
			currentString = append(currentString, b)
		} else {
			if len(currentString) >= max_length {
				strings = append(strings, string(currentString))
			}
			currentString = nil
		}
	}

	return strings, nil
}

func SplitString(s string) []string {
	slice := strings.Split(s, ",")
	return slice
}