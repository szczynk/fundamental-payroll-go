# Fundamental Payroll Program

## Description

This is a simple CLI payroll program built with Go that allows you to manage employee data and process salaries according to their grade and attendance. The program has the following features:

## Features

The following features are included in the program:

1. Add Employee: Add a new employee to the system with their name, grade, and other relevant details.
2. List Employees: Display a list of all employees currently in the system.
3. Process Salaries: Process employee salaries based on the salary matrix, which takes into account the basic salary, pay cuts, additional salary, and allowance. The program will also calculate additional salary for head of family employees who are married and male.
4. List Salary Matrices: Display a list of all salary matrices currently in the system.

## Getting Started

To run the application, follow the instructions below:

1. Clone this repository and navigate to the project directory.

   ```bash
   git clone https://github.com/szczynk/fundamental-payroll-go.git
   ```

1. Run `go mod tidy && go mod vendor`.

   ```bash
   go mod tidy && go mod vendor
   ```

1. Run `go run main.go`.

   ```bash
   go run main.go
   ```

## Usage

Once the program is running, you can interact with it through the command line interface. The program will display a menu with the available options:

1. Add Employee
2. List Employees
3. Process Salaries
4. List Salary Matrices
5. Exit

To select an option, simply enter the corresponding number and press enter. Follow the prompts to add employee details, process salaries, or view employee and salary matrix data.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
