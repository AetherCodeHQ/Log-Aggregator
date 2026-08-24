package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: log-aggregator <logfile>")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	levels := map[string]int{}
	sources := map[string]int{}
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total++
		for _, level := range []string{"INFO", "WARN", "ERROR", "DEBUG", "FATAL"} {
			if strings.Contains(line, level) {
				levels[level]++
			}
		}
		if idx := strings.Index(line, "["); idx >= 0 {
			end := strings.Index(line[idx:], "]")
			if end > 0 {
				src := line[idx+1 : idx+end]
				sources[src]++
			}
		}
	}
	fmt.Println("Log Aggregator Summary")
	fmt.Println("======================")
	fmt.Printf("Total lines: %d\n\n", total)
	fmt.Println("By level:")
	for _, l := range []string{"INFO", "WARN", "ERROR", "DEBUG", "FATAL"} {
		if c, ok := levels[l]; ok {
			pct := float64(c) / float64(total) * 100
			fmt.Printf("  %-6s %d (%.1f%%)\n", l, c, pct)
		}
	}
	if len(sources) > 0 {
		fmt.Println("\nBy source:")
		for s, c := range sources {
			fmt.Printf("  %-20s %d\n", s, c)
		}
	}
}