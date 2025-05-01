package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JoYBoY12/GOPROXY/proxy"
)

func main() {
	fmt.Println("Enter on/off/status")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in taking the input")
		return
	}

	input = strings.TrimSpace(input)

	switch input {
	case "on":
		proxy.Enable()
	case "off":
		proxy.Disable()
	case "status":
		proxy.Status()
	default:
		fmt.Println("Invalid command.Use on/off/status")

	}
}
