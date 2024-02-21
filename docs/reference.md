## Reference 


### Project Structure

The key part of the project is to organize all the advent of code events by year. Each year will const of a matrix of days and problems (2 problems per day).

```bash
.env # environment variables file
adventofcode.yaml # configuration file for the project
README.md # 
events/ # root of all the years of advent of code you have worked on
    <year>/
        <day>/
            <solution-files> # Test and code of your solutions
            inputs/ # inputs for testing and submiting the anwser online
            metadata.yaml # information of the current state of your solution
            README.md # Description of the problems (cache from adventofcode.com/year/day)
```       
