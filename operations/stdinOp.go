package operations

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

func CountStdinBytes(content []byte) error {
	if len(content) == 0 {
		return errors.New("no data provided")
	}

	content = bytes.Replace(content, []byte("\n"), []byte("\r\n"), -1)
	fmt.Printf("%d -> stdin\n", len(content))
	return nil
}

func CountStdinLines(content []byte) error {
	var lines int
	if len(content) == 0 {
		return errors.New("no data provided")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines++
	}

	fmt.Printf("%d -> stdin\n", lines)
	return nil

}

func CountStdinWords(content []byte) error {
	var words int
	if len(content) == 0 {
		return errors.New("no data provided")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	fmt.Printf("%d -> stdin\n", words)
	return nil
}

func CountStdinCharacters(content []byte) error {
	var chars int
	if len(content) == 0 {
		return errors.New("no data provided")
	}

	content = bytes.Replace(content, []byte("\n"), []byte("\r\n"), -1)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		chars++
	}

	fmt.Printf("%d -> stdin\n", chars)
	return nil

}

func GetStdin() []byte {
	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return []byte(content)
}
