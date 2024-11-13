package main

import (
	"errors" // Import the errors package to handle error creation
	"fmt"    // Import the fmt package for formatted I/O operations
)

// commandInspect handles the inspect command to display details of a caught Pokémon
func commandInspect(cfg *config, args ...string) error {
	// Check if exactly one argument (Pokémon name) is provided
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name") // Return error if not exactly one argument
	}

	// Get the Pokémon name from the arguments
	name := args[0]

	// Look up the Pokémon in the caughtPokemon map using the name
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon") // Return error if the Pokémon is not found
	}

	// Print the Pokémon's name
	fmt.Println("Name:", pokemon.Name)
	// Print the Pokémon's height
	fmt.Println("Height:", pokemon.Height)
	// Print the Pokémon's weight
	fmt.Println("Weight:", pokemon.Weight)
	// Print the stats of the Pokémon
	fmt.Println("Stats:")
	// Loop through each stat and print its name and value
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	// Print the types of the Pokémon
	fmt.Println("Types:")
	// Loop through each type and print its name
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	// Return nil indicating no error occurred
	return nil
}
