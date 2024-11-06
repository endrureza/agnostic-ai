package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// DD prints the given data to the console in JSON format and exits the program.
func DD(data ...interface{}) {
	for _, item := range data {
		// Marshal the data to JSON with indentation for readability.
		output, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling data:", err)
			fmt.Printf("%+v\n", item) // Print raw data if JSON conversion fails
		} else {
			fmt.Println(string(output))
		}
	}

	// Terminate the program
	os.Exit(1)
}
