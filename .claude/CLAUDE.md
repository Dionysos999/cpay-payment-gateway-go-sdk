# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**go2proto** is a small CLI tool that converts Go structs into Protocol Buffer message definitions. It generates pure proto3 syntax messages (no RPC, no gogo extensions). The entire implementation is a single ~247-line `main.go` file.

## Commands

### Build and Run
```bash
# Build
go build -o go2proto .

# Run example
go run . -p ./example/in -f ./example/out

# With filter
go run . -p ./example/in -f ./example/out -filter User
```

### CLI Flags
- `-p`: Package paths to analyze (required, accepts multiple values, relative paths allowed)
- `-f`: Output folder path (default `.`)
- `-filter`: Case-insensitive struct name filter (optional)

The output filename is always `output.proto` (constant defined in code).

## Architecture

### Single-File Design
All code is in `main.go`. The core data flow:

1. **loadPackages()** - Uses `golang.org/x/tools/go/packages` with `packages.LoadSyntax` mode to parse Go packages
2. **getMessages()** - Iterates `TypesInfo.Defs` to find exported structs, applies optional filter
3. **appendMessage()** - Extracts fields from each struct, converts names and types
4. **writeOutput()** - Renders proto3 using Go's `text/template`

### Type Conversion Logic

Key functions in `main.go`:

- `toProtoFieldTypeName()` - Maps Go types to proto types
  - Calls `normalizeType()` for basic types (`int`→`int64`, `float32`→`float`, `float64`→`double`)
  - Calls `splitNameHelper()` to extract type names from pointers/structs (string-based parsing, noted as TODO in code)
  - Slices are handled separately via `isRepeated()`

- `toProtoFieldName()` - Converts Go field names to proto style
  - 2-char names: all lowercase
  - Longer names: first rune lowercase only

- Field ordering: Preserves struct field position (1-indexed)

### Template
The proto template is hardcoded in `writeOutput()`. Output always uses:
- `syntax = "proto3"`
- `package proto`

### Known Limitations
- No unit tests exist (`go test ./...` is a no-op)
- Complex types (maps, channels, interfaces) are not supported
- Type name extraction uses string parsing (`splitNameHelper`) rather than full type system introspection

## Key Implementation Notes

- Only **exported** structs and fields are processed
- Duplicate message names are prevented via a `seen` map
- Messages are sorted alphabetically in output
- The `splitNameHelper()` function is noted as "ugly" and could be refactored to use proper type system methods
