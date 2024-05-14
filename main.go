package main

import (
	"bufio"
	"fmt"
	"internal/util"
	"os"
)

type cliCommand struct {
	name        string
	description string
}

func printCommands(commands map[string]cliCommand) {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k := range commands {
		fmt.Printf("%v: %v \n", k, commands[k].description)
	}

	fmt.Println()
}

func commandHelp() error {
	printCommands(cliCommands)
	return nil
}

var cliCommands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Display a help message",
	},
	"exit": {
		name:        "exit",
		description: "Exits the Pokedex",
	},
	"map": {
		name:        "map",
		description: "Loads the location areas",
	},
	"mapb": {
		name:        "mapb",
		description: "Loads the previous page of location areas",
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		if text == "help" {
			commandHelp()
		}
		if text == "map" {
			util.LoadLocationAreas()
		}
		if text == "mapb" {

		}
		if text == "exit" {
			break
		}
	}
}
