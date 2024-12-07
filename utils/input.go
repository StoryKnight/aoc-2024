package utils

import (
	"bufio"
	"os"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func ReadFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// clear a map
func Clear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}
