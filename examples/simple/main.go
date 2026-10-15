/******************************************************************************
 * Lua Virtualizer WebApi interface usage example.
 *
 * Version        : v1.0.0
 * Language       : Go
 * Author         : Bartosz Wójcik
 * Web page       : https://www.pelock.com
 *
 *****************************************************************************/

package main

import (
	"context"
	"fmt"
	"os"

	luavirtualizer "github.com/PELock/Lua-Virtualizer-Go"
)

func main() {
	client := luavirtualizer.New("ABCD-ABCD-ABCD-ABCD")

	source := `local function greet(name)
  print("Hello, " .. name)
end
greet("Lua")
`
	result, err := client.ObfuscateScriptSource(context.Background(), source)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if result.Error != luavirtualizer.ErrorSuccess {
		fmt.Fprintf(os.Stderr, "An error occurred, error code: %d\n", result.Error)
		os.Exit(1)
	}
	fmt.Println(result.Output)
}
