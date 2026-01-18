# CLI Calculator

<div align="center">

<img alt="GitHub Created At" src="https://img.shields.io/github/created-at/KieranPritchard/CLI-Calculator">

<img alt="GitHub License" src="https://img.shields.io/github/license/KieranPritchard/CLI-Calculator">

<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/t/KieranPritchard/CLI-Calculator">

<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/KieranPritchard/CLI-Calculator">

<img alt="GitHub language count" src="https://img.shields.io/github/languages/count/KieranPritchard/CLI-Calculator">

<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/KieranPritchard/CLI-Calculator">

</div>

## Project Description

### Objective

To build a CLI calculator will allow me to learn some of the basic elements of the Go langauge, to build up skill with the langauge so I can use it in my ethical hacking hobby and start building custom tools

### Technology and Tools Used

* **Language:** Go
* **Tools:** Git, VS Code.

### Challenges Faced

I had a trouble converting the arguements from the command line from strings to integers. I fixed this by using a package called `strconv`, which allows strings to be turned into different data types like integers. I also made a rookie error of not pointing to the correct pointers in the if statements which was a easy fix.

### Outcome

By completing this project, I gained a much better understanding of some of the core concepts in Go. This included parsing command-line arguments using the flag package, converting data types with strconv, and handling errors properly instead of letting the program crash. I also learned the importance of writing reusable functions to avoid repeating logic, as well as how pointers returned by command-line flags work.
The final result is a simple but functional CLI calculator that supports addition, subtraction, multiplication, and division through command-line flags. Overall, this project gave me a solid foundation to build on and will help me move on to more advanced Go projects, especially ones focused on networking, automation, and building custom tools for ethical hacking.

## How to Use the Project

**1. Clone The Repository:**
* Use Git to clone the project:
```
git clone https://github.com/KieranPritchard/Remote-Command-Execution.git
```
**2. Build the Program:**
* Navigate to the project directory.
* Build the binary:
```
go build
```
**3. Run the Calculator:**
* Use one of the following flags with two numbers inside quotes:
* Addition:
```
./Remote-Command-Execution -add "5 3"
```
* Subtraction:
```
./Remote-Command-Execution -sub "10 4"
```
* Multiplication:
```
./Remote-Command-Execution -multi "6 7"
```
* Division:
```
./Remote-Command-Execution -div "20 5"
```
* The program will output the result directly to the terminal.

## Licenses
License is found in the root of the repository.
