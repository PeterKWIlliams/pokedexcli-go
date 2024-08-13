package main

import (
	"time"

	"github.com/PeterKWIlliams/pokedexcli-go/internal/commands"
	"github.com/PeterKWIlliams/pokedexcli-go/internal/pokecache"
	"github.com/PeterKWIlliams/pokedexcli-go/internal/repl"
	"github.com/PeterKWIlliams/pokedexcli-go/internal/thirdparty/api"
)

func main() {
	client := api.NewClient(time.Second * 10)
	pokecache.NewCache(time.Second * 5)
	cfg := &commands.Config{
		Client: client,
	}
	repl.Start(cfg)
}
