package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readEnvFile reads and parses the given .env file
func readEnvFile(fileName string) (map[string]string, error) {
	envMap := make(map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0 // Initialize line number
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comment lines
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("error at line %d: Invalid format for line: %s", lineNumber, line)
		}

		key := strings.TrimSpace(parts[0])
		if strings.Contains(key, "#") || strings.Contains(key, " ") {
			// Skip keys with # or space
			continue
		}

		value := parseValue(parts[1])
		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %w", err)
	}

	return envMap, nil
}

// parseValue parses the value portion of a key-value pair
func parseValue(value string) string {
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		value = value[1 : len(value)-1]
		value = strings.TrimSpace(value) // Remove spaces before and after the value
	} else {
		// If the value is not quoted, then it ends at the first # character
		value = strings.Split(value, "#")[0]
		value = strings.TrimSpace(value) // Remove spaces before and after the value
	}

	return value
}

func main() {
	// Provide the name of the .env file to read
	envMap, err := readEnvFile(".env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the parsed environment variables
	for key, value := range envMap {
		fmt.Printf("%s=%s\n", key, value)
	}
}
