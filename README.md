# Lua Virtualizer — Obfuscate, Virtualize & Protect Python Code

**[Lua Virtualizer](https://www.pelock.com/products/lua-virtualizer)** is a tool to obfuscate, virtualize & protect Lua `.lua` scripts against reverse engineering with a VM engine (or finite-state automata machine), self-defending integrity checks, polymorphic string encryption, anti-debugging, anti-vm, anti-emulation and 50+ total obfuscation strategies!

Product page:

* https://www.pelock.com/products/lua-virtualizer

Multiple programming APIs available:

* https://www.pelock.com/products/lua-virtualizer/api

## Why Lua scripts need obfuscation?

Scripts are typically distributed as plain `.lua` files or bundled inside packages. That convenience means anyone with file access can read the full logic, hunt for credentials or API keys in strings, and steal your algorithms unless you take extra steps to hide intent.

[Lua](https://www.lua.org/) is one of the most popular general-purpose scripting languages. It is widely used for automation, backend services, data processing, tooling, and licensing/agent code that customers run on their own machines.

## Obfuscation strategies

Lua Virtualizer comes with many advanced obfuscation, virtualization & protection strategies. You can easily tune protection versus size and performance.

### ![Obfuscation](https://www.pelock.com/img/en/icons/obfuscation-32.png) Powerful obfuscation

Polymorphic string encryption is the core of the string pipeline: each build generates a fresh decryptor so literals never share one static algorithm. Integers, floats, and decoy noise join that layer. The result conceals literals and structure while preserving tested runtime behaviour.

### ![Processor](https://www.pelock.com/img/en/icons/processor-32.png) Code virtualization

Code virtualization is the main control-flow defense. Selected statements are lifted into a randomly generated VM engine (or FSA / flattened dispatcher) with shuffled dispatch tables, decoy opcodes, and an obfuscated dispatcher loop. Analysts must interpret the virtual machine instead of reading plain Lua.

### ![Node](https://www.pelock.com/img/en/icons/node-32.png) Finite-state automata (FSA)

Finite-state automata (FSA) obfuscation rewrites linear Lua statement blocks into dual-state automata with opaque schedulers and shuffled dispatch handlers. Instead of reading code top to bottom, analysts must follow numeric states, transition tables, and decoy paths to reconstruct the original order.

### ![Bug](https://www.pelock.com/img/en/icons/bug-32.png) Anti-debugging

Anti-debugging protection inserts polymorphic probes that detect attached debuggers, tracing hooks, and related host signals, together with anti-VM, anti-sandbox and anti-emulator checks. When a check fires, the obfuscated script exits silently instead of revealing protected logic under interactive analysis.

### ![System monitor](https://www.pelock.com/img/en/icons/system-monitor-32.png) Self-integrity checks

Self-defending integrity checks verify that the obfuscated file has not been patched. A bootstrap probe hashes the on-disk script and sets a tamper key when it no longer matches the obfuscated build. String decryptors consume that key, so edited scripts return garbage instead of plaintext.

### ![Brick link](https://www.pelock.com/img/en/icons/brick-link-32.png) Protection linker

The protection linker adds decoy functions and fake calls so a copied fragment still looks like real program code. Hidden traps fire only if someone edits the file or runs a piece of it on its own. When the script starts normally, those extras stay silent and the program runs as usual.

## Before and after obfuscation

Look at this example — the same script becomes harder to read at a glance after obfuscation.

### Sample Lua script before obfuscation

```lua
local function greet(name)
  print("Hello, " .. name)
end
greet("Lua")
```

### After obfuscation

Would you still recognise the original intent if you only had the obfuscated text and no prior copy of the script?

### How does Lua Virtualizer work?

The engine parses Lua source into an AST tree, then applies selectable transforms. Code virtualization (VM, FSA, or flattening) rewrites execution so analysts cannot read the script top to bottom. Polymorphic string encryption hides literals; self-defending integrity probes detect patched files. Also available: identifier renaming, numeric encryption, noise and decoy insertion, the protection linker, and anti-debugging checks. Many techniques are specific to this product; some ideas are shared with our other protection tools.

When all passes finish, the engine emits a new `.lua` file. Edge cases in the Lua grammar and hosting environments mean you should always test the output in your target runtime.

## Lua dialect (`target`)

The client `target` property defaults to `auto`. When `target` is `auto` or empty, the SDK omits the POST field. Sending `target=auto` is also valid. Explicit values are `5.1`, `5.2`, `5.3`, `5.4`, `5.5`, `luajit`, and `luau`.

Omitting `target` or sending `auto` detects the dialect by trying 5.1, then 5.2, 5.3, 5.4, 5.5, then Luau, and keeping the oldest that parses. LuaJIT is never inferred. A file whose only newer syntax is `goto` is detected as 5.2 because `goto` was introduced in Lua 5.2. Pass `luajit` when the script must be treated as LuaJIT. The JSON response field `target` is the dialect actually used.

## Protect your Lua scripts & algorithms

Take no chances, use **Lua Virtualizer** to obfuscate, virtualize and protect your Lua scripts and algorithms.

Our company has a long history in obfuscation technologies and code obfuscators (see our [PowerShell](https://www.pelock.com/products/powershell-pro-obfuscator), [Java](https://www.pelock.com/products/jobfuscator), [AutoIt](https://www.pelock.com/products/autoit-obfuscator), [x86 Assembly](https://www.pelock.com/products/obfuscator) &  obfuscators).

We actively bugfix, research and develop new obfuscation strategies for our tools.

You can count on our expertise and support in this field.

### Installation

The preferred way to install the Web API SDK is via Go modules.

```bash
go get github.com/PELock/Lua-Virtualizer-Go
```

The module path is `github.com/PELock/Lua-Virtualizer-Go`.

## Usage examples

### Example — `login/main.go`

```go
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
```

### Example — `simple/main.go`

```go
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
```

See the `examples/` directory in this repository for complete samples.

## Per-function markers

Markers are source syntax, not a library import. There is no separate decorator package. The marker is the first statement of a function. The engine reads it and strips it.

```lua
local function f()
  LUAVIRTUALIZER(VM(RND), ENABLE(STRINGS), DISABLE(encrypt_strings))
  return 1
end
```

`VM(...)` selects the dispatcher for that function. `ENABLE(...)` replaces the function's strategy set with the named groups or strategies. `DISABLE(...)` then removes the named groups or strategies. Clauses apply as VM, then ENABLE, then DISABLE, whatever order they are written in. Names are case-insensitive. `ENABLE(ALL)` is the same set `--all` turns on (every strategy except the four probes). `DISABLE(ALL)` clears every strategy. `FSA` is accepted as an alias of `FSM`. `self_defending` and `protection_linker` are not decorator names. They stay script-wide.

`decoy_functions`, every `BUCKET` strategy, `callback_registration_stubs`, and the `PROTECT` probes are inserted once for the whole script. A function `ENABLE` or `DISABLE` does not turn those passes on or off. A `VM(...)` nested inside `VM(NONE)` stays native Lua.

| Category | Name | Description |
| --- | --- | --- |
| Virtualization | `VM(VM)` | A word is a permuted opcode plus A, B, C, and Bx. The encrypted blob is a read-only code image. The Lua 5.1 shell fetches `__vll_code[pc]` and calls `__vll_hs[__vll_map[opcode]]`. Four decoy handlers share the table. A jump stores an absolute program counter. Exactly one mode applies to a function. |
| Virtualization | `VM(FSM)` | Same compiler, blob, and handler table as `VM(VM)`. The outer loop is shifted, and an opaque gate reads the virtual program counter before the live handler. `FSA` is an alias of `FSM`. |
| Virtualization | `VM(WEAVE)` | Same compiler and encrypted blob. Registers, constants, and instruction words share one table. Three superoperators cover load-and-arithmetic, compare-and-branch, and the call exit. A back-edge rewrites the words it will fetch again. Host calls are the only exit. |
| Virtualization | `VM(FLAT)` | Same compiler and handler table. The loop form is shifted again, and the opaque gate mixes the virtual program counter with the running checksum. |
| Virtualization | `VM(RND)` | Each run picks `VM(VM)`, `VM(FSM)`, or `VM(WEAVE)` for this function. `VM(FLAT)` is not in the draw. The same seed repeats the choice. Nested functions inherit it unless they set their own `VM(...)`. A `VM(RND)` nested inside `VM(NONE)` stays native Lua. |
| Virtualization | `VM(NONE)` | Leave this function, and every function nested inside it, as ordinary Lua. |
| Group | `NAMES` | `rename_variables`, `rename_parameters`, `rename_functions`, `rename_function_calls`, `resolve_constants`. |
| Group | `STRINGS` | `split_strings`, `modify_strings`, `string_char_array_vault`, `encrypt_strings`, `array_string_crypt`. |
| Group | `NUMBERS` | `integers_to_arrays`, `floats_to_arrays`, `affine_integer_mask`, `encrypt_integers`, `integers_to_floating`, `encrypt_floating`, `mba_binops`, `array_int_crypt`, `array_char_crypt`, `array_double_crypt`. |
| Group | `BOOLEANS` | `complexify_booleans`, `insert_ternary_operators`. |
| Group | `SHAPE` | `decoy_functions`, `lambda_decoys`, `insert_dead_code`, `opaque_branches`, `opaque_mixer_chain`, `try_finally_noise`, `literal_padding`, `fake_import_markers`, `shuffle_functions`. `decoy_functions` is inserted for the whole script. |
| Group | `BUCKET` | `insert_random_value_bucket`, `random_bucket_integers`, `random_bucket_arrays`, `random_bucket_functions`, `random_bucket_characters`, `random_bucket_anti_regex`, `random_bucket_autostart`. Inserted for the whole script. |
| Group | `CALLS` | `dynamic_getattr_calls`, `obfuscate_imports`, `callback_registration_stubs`. `callback_registration_stubs` is inserted for the whole script. |
| Group | `COMMENTS` | `remove_comments`. |
| Group | `PROTECT` | `detect_debugger`, `anti_vm`, `anti_sandbox`, `anti_emulator`. Inserted for the whole script. A positive detection ends the process with no error message. |
| Group | `ALL` | `ENABLE(ALL)` selects every strategy except the four probes. `DISABLE(ALL)` clears every strategy, including the probes. |
| Strategy | `rename_variables` | `NAMES`. Local and global variables become random identifiers. Every reference updates together, and scope rules stay intact. |
| Strategy | `rename_parameters` | `NAMES`. Parameter names and in-body bindings are rewritten to opaque tokens. |
| Strategy | `rename_functions` | `NAMES`. Function definitions are renamed and calls are retargeted. Metamethods such as `__index` and built-in library functions stay untouched. |
| Strategy | `rename_function_calls` | `NAMES`. Direct call sites look up the target through a name table instead of the literal identifier. |
| Strategy | `resolve_constants` | `NAMES`. Module-level constants are inlined at every use, and unused declarations are removed. |
| Strategy | `split_strings` | `STRINGS`. Literals are split into chunks and joined at runtime, so the whole token is no longer in one place. |
| Strategy | `modify_strings` | `STRINGS`. Literals are rewritten as shifted characters and rebuilt by a small runtime transform. |
| Strategy | `string_char_array_vault` | `STRINGS`. Strings are rebuilt from numeric character codes instead of quoted text. |
| Strategy | `encrypt_strings` | `STRINGS`. Each run emits a unique per-string decryptor. Two outputs do not share the same algorithm shape. |
| Strategy | `array_string_crypt` | `STRINGS`. String table elements are encrypted independently and rebuilt on demand. |
| Strategy | `integers_to_arrays` | `NUMBERS`. Integers are stored in tables and read by index, with decoy entries mixed in. |
| Strategy | `floats_to_arrays` | `NUMBERS`. Floats are lifted into table slots, with decoy values mixed in. |
| Strategy | `affine_integer_mask` | `NUMBERS`. Reversible bit masks wrap integers without changing the result. The helper matches the selected dialect. |
| Strategy | `encrypt_integers` | `NUMBERS`. Plain integers become equivalent arithmetic or bitwise expressions. |
| Strategy | `integers_to_floating` | `NUMBERS`. Whole numbers are reconstructed from `math` expressions and cast back with `math.floor`. |
| Strategy | `encrypt_floating` | `NUMBERS`. Floating-point literals become equivalent `math` expressions. |
| Strategy | `mba_binops` | `NUMBERS`. Binary operators are replaced with mixed boolean-arithmetic identities. |
| Strategy | `array_int_crypt` | `NUMBERS`. Integer table contents are stored encrypted and decoded at runtime. |
| Strategy | `array_char_crypt` | `NUMBERS`. Character table literals are stored encrypted and restored by generated decoder code. |
| Strategy | `array_double_crypt` | `NUMBERS`. Float table constants are hidden as encoded bit patterns and reconstructed while the program runs. |
| Strategy | `complexify_booleans` | `BOOLEANS`. Simple true/false tests are rewritten into longer equivalent conditions. |
| Strategy | `insert_ternary_operators` | `BOOLEANS`. Simple branches collapse into `and`/`or` conditional expressions. |
| Strategy | `decoy_functions` | `SHAPE`. Plausible functions that are never called. Inserted for the whole script. |
| Strategy | `lambda_decoys` | `SHAPE`. Unused anonymous functions near the top of the script. They are never called. |
| Strategy | `insert_dead_code` | `SHAPE`. Statements under opaque never-true conditions. The result of the script stays the same. |
| Strategy | `opaque_branches` | `SHAPE`. Always-false branches that look like real control flow. |
| Strategy | `opaque_mixer_chain` | `SHAPE`. Mixer functions and guards that never change the returned value. |
| Strategy | `try_finally_noise` | `SHAPE`. Lua has no try/finally, so the pass emits an inert `pcall` plus cleanup. |
| Strategy | `literal_padding` | `SHAPE`. A bulky string constant that live logic never reads. It shifts line numbers and inflates the file. |
| Strategy | `fake_import_markers` | `SHAPE`. Comments that mimic extra `require` calls. `remove_comments` deletes them. |
| Strategy | `shuffle_functions` | `SHAPE`. Permutes top-level function definitions when the execution order stays valid. |
| Strategy | `insert_random_value_bucket` | `BUCKET`. Enables the module-level decoy pools. Inserted for the whole script. |
| Strategy | `random_bucket_integers` | `BUCKET`. Unused integer decoys in the pool. Inserted for the whole script. |
| Strategy | `random_bucket_arrays` | `BUCKET`. Unused table decoys in the pool. Inserted for the whole script. |
| Strategy | `random_bucket_functions` | `BUCKET`. Unused function decoys in the pool. Inserted for the whole script. |
| Strategy | `random_bucket_characters` | `BUCKET`. Unused character decoys in the pool. Inserted for the whole script. |
| Strategy | `random_bucket_anti_regex` | `BUCKET`. Regex-shaped decoy literals. Inserted for the whole script. |
| Strategy | `random_bucket_autostart` | `BUCKET`. Decoy strings that resemble autostart paths and are never used by live code. Inserted for the whole script. |
| Strategy | `dynamic_getattr_calls` | `CALLS`. Call targets are built from fragments and resolved through a table lookup. |
| Strategy | `obfuscate_imports` | `CALLS`. `require` is rewritten so the module name is built at runtime. Those strings stay available for later string encryption. |
| Strategy | `callback_registration_stubs` | `CALLS`. Snippets that resemble callback registration and never wire an event source. Inserted for the whole script. |
| Strategy | `remove_comments` | `COMMENTS`. Removes `--` comments from the emitted script. |
| Strategy | `detect_debugger` | `PROTECT`. Detects a debugger attached to the process (`debug.gethook()`, and `IsDebuggerPresent` on Windows). Inserted for the whole script. A positive detection ends the process with no error message. |
| Strategy | `anti_vm` | `PROTECT`. Detects VMware, VirtualBox, Parallels, KVM, and Linux DMI product or vendor strings. Inserted for the whole script. |
| Strategy | `anti_sandbox` | `PROTECT`. Detects a very low CPU count, Sandboxie, Joe Sandbox, and analysis-host tokens. Inserted for the whole script. |
| Strategy | `anti_emulator` | `PROTECT`. Detects WINE, Bochs, QEMU, XEN, and a short timing self-check. Inserted for the whole script. |

## Product page

https://www.pelock.com/products/lua-virtualizer

Web API:

https://www.pelock.com/products/lua-virtualizer/api

## Windows GUI client and command line version

You can download it at:

https://www.pelock.com/products/lua-virtualizer

#### Lua Virtualizer comes also with full GUI version for Windows

![Lua Virtualizer Windows Client](https://www.pelock.com/img/en/products/lua-virtualizer/lua-virtualizer-obfuscator-1-0-main.png)

#### Obfuscation options

![Lua Virtualizer obfuscation options](https://www.pelock.com/img/en/products/lua-virtualizer/lua-virtualizer-obfuscator-1-0-options.png)

#### Command line interface aka CLI

![Lua Virtualizer command-line interface](https://www.pelock.com/img/en/products/lua-virtualizer/lua-virtualizer-obfuscator-1-0-command-line.png)

Lua Virtualizer ships with a command-line interface for Windows and Linux automation. Use it to integrate obfuscation into build servers, CI jobs, or batch packaging.

## Visual Studio Code / Cursor extension

![Lua code virtualization in Visual Studio Code](https://www.pelock.com/img/en/products/lua-virtualizer/lua-virtualizer-obfuscator-code-virtualization-vscode-extension.gif)

- [https://marketplace.visualstudio.com/items?itemName=PELock.lua-virtualizer](https://marketplace.visualstudio.com/items?itemName=PELock.lua-virtualizer)
- [https://open-vsx.org/extension/PELock/lua-virtualizer](https://open-vsx.org/extension/PELock/lua-virtualizer)

## Demo mode limitations

A missing or invalid key is limited to 1000 bytes and the server applies only flat code virtualization (`code_virtualization=flat`), ignoring every other strategy field.

Bartosz Wójcik

* Visit my site at — https://www.pelock.com
* X — https://x.com/PELock
* GitHub — https://github.com/PELock
