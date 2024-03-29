# Contributing Guide

Before I say anything, thank you for considering putting effort into making this project better.

> Notes: This is a living document that I will use to learn more about the development cycle of a CLI tool.

## Structure of the Advent of Code CLI

There are two main parts `pkg` and `cmd` that together work to satisfy the interface of the `aoc` CLI.

- `pkg`: in the `pkg` part is where the business logic for the CLI lives.
- `cmd`: in the `cmd` part of the project is where the implementation details of the CLI User Experience.

the `examples` folder has a list of languages and the advent of code years I use to test manually. In the future, it will do the integration tests.

## Development Flow

- Automate machine work as much as possible into the GitHub Actions (e.g. build, publish, test).
- Filter all changes proposed by proper checks from GitHub Actions pipelines.
- Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#summary).

### Branch Flow

### Proposing a new change

No matter what type of change we want to propose we need to follow a simple workflow to make sure we are using all the automated processes that are there to help us.

#### Steps

- A good pre-work you can do before starting contributing is to find any issues or open PRs that might be related to your situation.
- If you have not done it, fork this project in a space you have pushing privileges.
- Check the release branch you want to contribute to and create your custom branch.
- Make the changes you find valuable to solve the issue you are fixing.
- Run `go test ./...` and make sure you are not breaking other parts of the codebase.
- Create a PR to the release branch associated with the changes and write all the information the maintainer would need to get context to understand the change.

Thank you!

