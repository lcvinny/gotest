package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := []string{"test"}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("go", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "PASS") {
			line = "\033[32m" + line + "\033[0m"
		} else if strings.Contains(line, "FAIL") {
			line = "\033[31m" + line + "\033[0m"
		}
		fmt.Println(line)
	}
}
