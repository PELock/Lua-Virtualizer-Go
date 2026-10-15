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
	// A missing or invalid key is limited to 1000 bytes and the server applies
	// only flat code virtualization (code_virtualization=flat), ignoring every
	// other strategy field.
	client := luavirtualizer.New("ABCD-ABCD-ABCD-ABCD")
	result, err := client.Login(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Demo version status - %v\n", result.Demo)
	fmt.Printf("License expiration - %s\n", result.LicenseExpiration)
	fmt.Printf("Total obfuscations - %d\n", result.UsagesTotal)
	fmt.Printf("Max. source code size - %d\n", result.StringLimit)
}
