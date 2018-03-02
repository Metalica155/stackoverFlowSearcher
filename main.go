package main

import (
	"bufio"
	"os"
)

func main() {
}

func read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
