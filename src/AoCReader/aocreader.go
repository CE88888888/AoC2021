package aocreader

import (
	"bufio"
	"os"
)

//open file from same directory
func OpenFile(file string) (scanner *bufio.Scanner) {
	if file, err := os.Open(file); err == nil {
		scanner = bufio.NewScanner(file)
	} else {
		println("Err opening file")
	}
	return
}
