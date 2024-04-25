package iohandler

import (
	"gfwc/operations"
	"log"
)

func HandleStdin(bytes, lines, words, chars bool) {
	content := operations.GetStdin()
	if bytes {
		if err := operations.CountStdinBytes(content); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if lines {
		if err := operations.CountStdinLines(content); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if words {
		if err := operations.CountStdinWords(content); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if chars {
		if err := operations.CountStdinCharacters(content); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
}

func HandleFile(filepath string, bytes, lines, words, chars bool) {
	if bytes {
		if err := operations.CountBytesNumber(filepath); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if lines {
		if err := operations.CountLinesNumber(filepath); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if words {
		if err := operations.CountWordsNumber(filepath); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
	if chars {
		if err := operations.CountCharactersNumber(filepath); err != nil {
			log.Fatalf("ERROR -> %v", err)
		}
	}
}
