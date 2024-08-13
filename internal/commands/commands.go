package commands

import (
	"github.com/PeterKWIlliams/pokedexcli-go/internal/thirdparty/api"
)

type Config struct {
	Client   *api.Client
	Next     *string
	Previous *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: ":displays available commands",
			Callback:    HelpCommand,
		},

		"exit": {
			Name:        "exit",
			Description: "shutdown the program",
			Callback:    ExitCommand,
		},
		"map": {
			Name:        "map-go-forward",
			Description: "Get next list of 20 locations",
			Callback:    mapF,
		},
		"mapB": {
			Name:        "map-go-back",
			Description: "Get list of previous 20 locations",
			Callback:    mapB,
		},
	}
}
