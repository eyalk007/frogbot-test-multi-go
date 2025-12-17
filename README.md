# Frogbot Multi-Go Test Repository

This repository contains two separate Go projects to test Frogbot's handling of multiple auto-detected working directories.

## Structure

- **project1/**: Go project using golang.org/x/net v0.35.0
- **project2/**: Go project using golang.org/x/net v0.38.0

Each project has its own `go.mod` and `go.sum` files.

## Purpose

This tests that Frogbot correctly:
1. Auto-detects both Go projects
2. Scans each project separately
3. Maps vulnerabilities to the correct working directory
4. Can apply fixes to the correct project

