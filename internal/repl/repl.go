package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PeterKWIlliams/pokedexcli-go/internal/commands"
)

func Start(cfg *commands.Config) {
	for {

		fmt.Println("Type 'help' for a list of commands")
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		commandName := strings.Trim(input, " ")
		command, ok := commands.GetCommands()[commandName]

		if ok {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
		continue

	}
}
