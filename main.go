package main

import (
	"encoding/json"
	"fmt"

	"github.com/neocortical/flexjson/flexjson"
)

func main() {
	fmt.Println()

	for _, exampleJSON := range flexjson.Examples {
		var example flexjson.Example
		err := json.Unmarshal([]byte(exampleJSON), &example)
		if err != nil {
			fmt.Printf("Unmarshal error: %v\n", err)
			continue
		}

		fmt.Printf("Example: %s\n\tSprocket ID: %d\n\tPartial? %t\n\n", example.Desc, example.Sprocket.ID, example.Sprocket.Partial)

		out, err := json.MarshalIndent(example, "", "   ")
		if err != nil {
			fmt.Printf("Marshal error: %v\n", err)
			continue
		}

		fmt.Printf("Output: \n%s\n\n", string(out))
	}
}
