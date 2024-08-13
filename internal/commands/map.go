package commands

import (
	"errors"
	"fmt"
)

func mapF(cfg *Config) error {
	locations, err := cfg.Client.GetLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next
	if locations.Next == nil {
		return errors.New("no next page ")
	}
	if locations.Previous != nil {
		fmt.Println(*locations.Previous)
	}
	for _, location := range locations.Results {
		fmt.Println(location, location.Name)
	}
	return nil
}

func mapB(cfg *Config) error {
	if cfg.Previous == nil {
		return errors.New("no previous page")
	}
	locations, err := cfg.Client.GetLocations(cfg.Previous)
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next
	if err != nil {
		return err
	}
	if locations.Previous != nil {
		fmt.Println(*locations.Previous, "this the preious page")
	}
	for _, location := range locations.Results {
		fmt.Println(location, location.Name)
	}

	return nil
}
