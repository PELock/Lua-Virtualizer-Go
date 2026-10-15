/******************************************************************************
 * Lua Virtualizer WebApi interface usage example.
 *
 * Every public option is set so the client surface is visible.
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
	client.EnableCompression = false
	client.RenameStyle = luavirtualizer.RenameStyleIL
	// auto omits the POST field. Explicit: 5.1, 5.2, 5.3, 5.4, 5.5, luajit, luau.
	client.Target = luavirtualizer.TargetAuto
	client.CodeVirtualization = luavirtualizer.CodeVirtualizationVM

	client.SelfDefending = true
	client.ProtectionLinker = true
	client.DetectDebugger = true
	client.AntiVM = true
	client.AntiSandbox = true
	client.AntiEmulator = true

	client.RenameVariables = true
	client.RenameParameters = true
	client.RenameFunctions = true
	client.RenameFunctionCalls = true
	client.ShuffleFunctions = true
	client.ResolveConstants = true
	client.SplitStrings = true
	client.ModifyStrings = true
	client.EncryptStrings = true
	client.StringCharArrayVault = true
	client.EncryptIntegers = true
	client.EncryptFloating = true
	client.MBABinops = true
	client.IntegersToFloating = true
	client.IntegersToArrays = true
	client.FloatsToArrays = true
	client.AffineIntegerMask = true
	client.ArrayIntCrypt = true
	client.ArrayCharCrypt = true
	client.ArrayDoubleCrypt = true
	client.ArrayStringCrypt = true
	client.InsertRandomValueBucket = true
	client.RandomBucketIntegers = true
	client.RandomBucketArrays = true
	client.RandomBucketFunctions = true
	client.RandomBucketCharacters = true
	client.RandomBucketAntiRegex = true
	client.RandomBucketAutostart = true
	client.InsertTernaryOperators = true
	client.ComplexifyBooleans = true
	client.OpaqueBranches = true
	client.OpaqueMixerChain = true
	client.InsertDeadCode = true
	client.TryFinallyNoise = true
	client.DecoyFunctions = true
	client.LambdaDecoys = true
	client.LiteralPadding = true
	client.FakeImportMarkers = true
	client.DynamicGetattrCalls = true
	client.ObfuscateImports = true
	client.CallbackRegistrationStubs = true
	client.RemoveComments = true

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
	fmt.Println(result.Target)

	// Per-function marker (source syntax, not a library import). The marker is
	// the first statement of a function. The engine reads it and strips it.
	//
	// local function f()
	//   LUAVIRTUALIZER(VM(FSA), OBFUSCATION(encrypt_strings, NUMBERS))
	//   return 1
	// end
	//
	// VM(VM|FSA|FLAT|NONE) selects the dispatcher; NONE stays native Lua.
	// OBFUSCATION(...) is a blacklist: NONE, ALL, group names NAMES, STRINGS,
	// NUMBERS, BOOLEANS, SHAPE, BUCKET, CALLS, COMMENTS, PROTECT, or a specific
	// snake_case strategy flag.
}
