package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

func PickFromMenu(prompt string, menu []string) string {
	var v string
	for i, opt := range menu {
		fmt.Printf("%d: %s\n", i, opt)
	}
	fmt.Print(prompt)
	fmt.Scanln(&v)
	vInt, err := strconv.Atoi(v)
	if vInt < 0 || vInt >= len(menu) || err != nil {
		fmt.Println("Invalid choice")
		return PickFromMenu(prompt, menu)
	}
	return menu[vInt]
}
