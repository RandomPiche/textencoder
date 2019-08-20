package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var charset = []struct {
	input string
	want  string
}{
	{"’", "'"},
	{"”", "\""},
	{"“", "\""},
	{"’", "&#39;"},
	{"”", "&#34;"},
	{"“", "&#34;"},
	{"-", "&#45;"},
	{"&", "&amp;"},
	{"à", "&#224;"},
	{"â", "&#226;"},
	{"è", "&#232;"},
	{"é", "&#233;"},
	{"ê", "&#234;"},
	{"ë", "&#235;"},
}

func main() {
	start := time.Now()
	src, dist := os.Args[1], os.Args[2]
	lines, err := readLines(src)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	if err := writeLines(encodeLines(lines), dist); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	fmt.Printf("Text file encoded to %s in %v\n", dist, time.Since(start).String())
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// encodeLines encodes et trims the given lines.
func encodeLines(lines []string) []string {
	var l []string
	for _, line := range lines {
		for _, char := range charset {
			line = strings.ReplaceAll(line, char.input, char.want)
		}
		l = append(l, strings.TrimSpace(line))
	}
	return l
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
