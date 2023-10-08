# How to work with Puzzles

With a functional programing model we are going to help people working with Advent of Code problems.

## Init

What questions should we ask?
- Name or default
- Language runtime from list of supported languages
- version of the supported language
- want git? 
    - if yes, repository in mind?
- 

## Get
The user will write `aoctl puzzle get/sync` with some informaation


Information needed:
- Year // if year is not provided it will go over 2015-lastDecemberyear
- Day // if day is not provided it will go over 1-25. If we are between dec 1 and dec 25 and YEAR is current, only do until "TODAY's day"
- Part // if part is not provided it will try to get both parts (if first part is not solved it will only get one

> Alert the user if bulk of information is going to be created (in order words, if any information is missing and multiple parts are going to be downlaoded request confirmation from user)

The user should expect a file structure for each year/day

- A puzzle.yaml file with the current state of the puzzle
- For each part (there are two) there should be a file of the selected language of the project that has a main function where the user will write the business logic
- inputs folder. The user will use this folder for testing or submit purposes. [TODO: not full flesh out yet]
- Test files for each part with the description of the part as comments in the file.
- Readme file that explains how to use each file of this puzzle solver


puzzle.yaml
- language this overwrites the root language runtime
- Language version
- depedencies
- Parts
    - solver: filename
      input: inputname
      submitted: bool

## File Structure 

Example of a aocproject files:

/<custom-name-for-project>
     README.md
    .gitignore // the session cookie file should be here
    .git/ // initialized git with provided remote or empty
    .aoc.yaml // configuration like language, and other preferences
    /src
        /<year>
            /<day>
                README.md
                puzzle.yaml
                part1.go
                part1_test.go
                part2.go
                part2_test.go
                /inputs
                    input.txt

For language especific things
- Golang: work directory in a go.mod root file that adds every day of every year?
- Golang: Do we need a main runner? // TODO: Think about the main runner. Example we can replace it with Makefiles

