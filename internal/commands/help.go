package commands

import (
	"fmt"
)

func HelpCommand(config *Config) error {
	fmt.Println("Available commands:")
	for name, cmd := range GetCommands() {
		fmt.Printf(" %s %s\n", name, cmd.Description)
	}
	return nil
}
