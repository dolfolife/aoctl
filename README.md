# Advent of Code CLI

This is a personal project I have to learn more about writing CLI tools in GoLang. The idea behind this project is to help me solve the [Advent of Code](adventofcode.com/) since I have noticed some patterns while solving them.

> Disclaimer: This is yet another CLI that solves a common problem and there are probably N solutions out there. My motivation is not only to provide a CLI but to learn how to write, develop, release, and test them.

I hope you enjoy it and feel free to help me understand the best practices of writing CLI tools or GoLang packages by writing an issue or using the [CONTRIBUTING.md](./CONTRIBUTING.md) to write a PR.

## Advent Of Code Series

Advent of Code is a website made by Eric Wastl. I recommend watching [Advent of Code: Behind the Scenes](https://www.youtube.com/watch?v=CFWuwNDOnIo&ab_channel=CodingTech).

## Quick Note

The inspiration to build a tool around solving comes from laziness of not going to a browser to find puzzles and submit my answers.

I found [Advent Of Code Go](https://github.com/alexchao26/advent-of-code-go/tree/main) and it inspired me to build my own tool and learn more GoLang. 


# AoC CLI Documentation

## Pre-requisite

You need your session ID from the Advent of Code website. I opened [Session ID is a manual process](https://github.com/dolfolife/aoctl/issues/1) issue.

## Puzzles

The Advent of Code puzzles have two parts. Each part has a single string input and a single output.

By that, we have to create an interface that allows us to run each solution against the input and submit the solution to the server of Advent of Code.

## Commands

### Init
Initialize the aoc project.

```bash
aoc init <path> 
```

### Puzzle
Get The puzzle information with the option to save it locally in the repository.

```bash
aoc puzzle [OPTIONS]
```

#### Options

##### -s --session (required)
Advent of Code Session. You can get this in the Cookies of the website. 

##### -d --day (required)
Day of the puzzle you want to get.

##### -y --year (required)
Year of the puzzle you want to get.

##### --sync (optional)
To save the puzzle locally at `<year>/<day>/puzzle.md

### test
Locally test your answers. This relies that each `<year>/<day>` solution has its own tests.

```bash
aoc test [OPTIONS]
```

### submit
Submit your answer to the Advent of Code.

```bash
aoc submit [OPTIONS]
```

#### Options

##### -s --session (required)
Advent of Code Session. You can get this in the Cookies of the website. 

##### -d --day (required)
Day of the puzzle you want to get.

##### -y --year (required)
Year of the puzzle you want to get.

