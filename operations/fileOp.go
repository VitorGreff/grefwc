package operations

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountBytesNumber(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Printf("%d -> %s\n", len(content), filepath)
	return nil
}

func CountLinesNumber(filepath string) error {
	var lines int
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines++
	}
	// ignores last \n
	fmt.Printf("%d -> %s\n", lines, filepath)
	return nil
}

func CountWordsNumber(filepath string) error {
	var words int
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}

	fmt.Printf("%d -> %s\n", words, filepath)
	return err
}

func CountCharactersNumber(filepath string) error {
	var chars int
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		chars++
	}

	fmt.Printf("%d -> %s\n", chars, filepath)
	return err
}
