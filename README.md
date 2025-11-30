# Advent of Code CLI

[![Go Reference](https://pkg.go.dev/badge/github.com/dolfolife/aoctl#section-readme.svg)](https://pkg.go.dev/github.com/dolfolife/aoctl#section-readme)
[![codecov](https://codecov.io/github/dolfolife/aoctl/graph/badge.svg?token=GTFZX1J2WX)](https://codecov.io/github/dolfolife/aoctl)

A simple CLI tool to organize and solve [Advent of Code](https://adventofcode.com/) problems. It handles downloading inputs, generating boilerplate code, and running your solutions, so you can focus on the logic.

## Features

- **Project Initialization**: Sets up a clean workspace for your AoC solutions.
- **Synchronization**: Downloads puzzle descriptions (README.md) and inputs for all available days.
- **Code Generation**: Generates simple Go boilerplate for each day.
- **Zero Friction**: No complex interfaces to implement. Just write `Part1` and `Part2` functions.

## Installation

```bash
go install github.com/dolfolife/aoctl@latest
```

## Usage

### 1. Initialize Project

Start by creating a new directory for your Advent of Code solutions and initializing it:

```bash
mkdir my-aoc-solutions
cd my-aoc-solutions
aoctl init .
```

### 2. Configure Session

To download inputs, you need to provide your Advent of Code session cookie. You can find this in your browser's developer tools (Application -> Cookies) when logged into adventofcode.com.

```bash
aoctl session --session <your-session-cookie>
```

### 3. Synchronize Puzzles

Download the puzzles and generate the boilerplate code for the current year (or previous years).

```bash
aoctl synchronize
```

This will create a directory structure like this:

```
events/
  2024/
    01/
      README.md          # The puzzle description
      input/
        input.txt        # The puzzle input
      main.go            # Entry point to run your solution
      solution.go        # Where you write your code
      solution_test.go   # Tests for your solution
```

### 4. Solve the Puzzle

Open `events/2024/01/solution.go` and implement `Part1` and `Part2`.

```go
package main

func Part1(input string) (string, error) {
    // Your logic here
    return "answer", nil
}

func Part2(input string) (string, error) {
    // Your logic here
    return "answer", nil
}
```

### 5. Run and Test

Run your solution:

```bash
cd events/2024/01
go run .
```

Run tests:

```bash
go test .
```

## Helper Functions

The `pkg/puzzle` package provides a simple helper to read the input file:

```go
import "github.com/dolfolife/aoctl/pkg/puzzle"

input, err := puzzle.ReadInput("input/input.txt")
```

## Contributing

PRs are welcome! Please check [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines.
