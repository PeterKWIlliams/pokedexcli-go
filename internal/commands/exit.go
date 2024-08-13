package commands

import (
	"fmt"
	"os"
)

func ExitCommand(config *Config) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}
