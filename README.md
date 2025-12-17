# Frogbot Multi-Go Test Repository

This repository contains two separate Go projects to test Frogbot's handling of multiple auto-detected working directories.

## Structure

```
fun-projects/
├── project1/  - Go project using golang.org/x/net v0.35.0
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── project2/  - Go project using golang.org/x/net v0.38.0
    ├── go.mod
    ├── go.sum
    └── main.go
```

Each project has its own `go.mod` and `go.sum` files in their respective directories.

## Purpose

This tests that Frogbot correctly:
1. Auto-detects both Go projects in the `fun-projects/` directory
2. Scans each project separately without flattening results
3. Maps vulnerabilities to the correct working directory
4. Can apply fixes to the correct project

## Testing Instructions

Run Frogbot on this repository without specifying working directories to test auto-detection:

```bash
# Frogbot should auto-detect:
# - fun-projects/project1
# - fun-projects/project2
```
