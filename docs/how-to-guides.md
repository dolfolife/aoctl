## How to Guides

These are easy to follow guides to solve an specific problem you might have.

### Start a new project

In this guide we will create a new project and add your first solution for a [Advent of Code](https://adventofcode.com/) site.

#### Initialize the project

To start a new working folder for advent of code use the command [`init`](https://github.com/dolfolife/aoctl?tab=readme-ov-file#init).

> By default it will create a new folder named `adventofcode` but you can change that with the parameter `--path` to initialize the project with a specific name.
```bash
aoctl init -p aoc-solutions
```

Next, we need to add our user session to the project by editing the `.env` file. To fetch this user session you can check the `session` cookie in the cookies section of the browser inspect console.

> The session cookie is a httpOnly cookie, so it can't be fetch using `document.cookie` object

To add the session to the project use the `session` cmd

```bash
aoctl session -s <SESSION>
```
