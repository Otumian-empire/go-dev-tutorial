package hello

import "fmt"

func SayHello(name string) (string, error) {
	if name == "" || len(name) == 0 {
		return "", fmt.Errorf("provide a value for name")
	}

	return fmt.Sprintln("Hello", name, "!!"), nil
}
