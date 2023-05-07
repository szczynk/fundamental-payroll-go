package input

import (
	"bufio"
	"io"
)

// InputReader is a implementation of InputReader that wraps bufio.Scanner
type InputReader struct {
	scanner *bufio.Scanner
}

func NewInputReader(input io.Reader) *InputReader {
	scanner := bufio.NewScanner(input)

	r := new(InputReader)
	r.scanner = scanner

	return r
}

func (r *InputReader) Scan() (string, error) {

	var input string
	if r.scanner.Scan() {
		input = r.scanner.Text()
	}

	// for r.scanner.Scan() {
	// 	input = r.scanner.Text()
	// }

	if err := r.scanner.Err(); err != nil {
		return "", err
	}
	return input, nil
}
