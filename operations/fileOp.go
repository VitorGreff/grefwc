package operations

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
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

	content = bytes.Replace(content, []byte("\r\n"), []byte("\n"), -1)
	fmt.Printf("%d bytes -> %s\n", len(content), filepath)
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
	fmt.Printf("%d lines -> %s\n", lines, filepath)
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

	fmt.Printf("%d words -> %s\n", words, filepath)
	return nil
}

func CountCharactersNumber(filepath string) error {
	var chars int
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	content = bytes.Replace(content, []byte("\r\n"), []byte("\n"), -1)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		chars++
	}

	fmt.Printf("%d chars -> %s\n", chars, filepath)
	return nil
}
