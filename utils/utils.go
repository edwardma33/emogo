package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetUserInput(prompt string) string {
	var v string
	fmt.Print(prompt)
	fmt.Scanln(&v)
	return v
}

func RunCommand(cmdStr string) error {
	tkns := strings.Split(cmdStr, " ")

	cmd := exec.Command(tkns[0], tkns[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
