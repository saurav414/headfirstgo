// package goutils helps to maintain all custom code at one place
package goutils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat helps to take a float number from console
func GetFloat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		// log.Fatal(err)
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return number, nil

}
