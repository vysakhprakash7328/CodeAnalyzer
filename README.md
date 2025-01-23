# CodeAnalyze CLI Tool

CodeAnalyze is a powerful and easy-to-use command-line interface (CLI) tool designed to analyze your codebase. It provides actionable suggestions for code improvements, formats code, and generates useful metrics, such as function counts, variable counts, and line counts. Currently, it supports Python, with plans to add support for other programming languages in the future.

## Features

- **Code Suggestions**: Identifies areas in your code that can be improved and provides actionable suggestions.
- **Code Formatting**: Formats your code to ensure it adheres to best practices and coding standards.
- **Metrics Analysis**: 
  - Counts the number of functions in your code.
  - Counts the number of variables.
  - Calculates the total number of lines in the code.
- **Live Watching**: Tracks changes in your codebase and provides real-time suggestions and metrics updates.

---

## Installation

### Prerequisites
- Go 1.19+ installed on your system

### Steps
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd codeanalyze
   ```

2. Build the binary:
   ```bash
   go build -o codeanalyze
   ```

3. Add the binary to your PATH to use it globally (optional):
   ```bash
   mv codeanalyze /usr/local/bin/
   ```

---

## Usage

### Analyze Code
Use the `analyze` command to analyze a file or directory for improvements, formatting, and metrics.

#### Syntax
```bash
codeanalyze analyze [options] <path-to-code>
```

#### Options
- `--metrics`: Generate metrics like function count, variable count, and line count.
- `--suggestions`: Provide actionable suggestions for code improvements.
- `--format`: Format the code according to best practices.
- `--watch`: Enable live watching for real-time analysis and suggestions.

#### Example
To analyze the `main.go` file with all options:
```bash
codeanalyze analyze --metrics --suggestions --format main.py
```

To watch a directory for changes and analyze in real-time:
```bash
codeanalyze analyze --metrics --suggestions --format --watch main.py
```

#### Sample Output
```
Code Suggestions:
- Line 12: Consider renaming variable 'x' to a more descriptive name.
- Line 25: Function 'Calculate' has high complexity; consider refactoring.

Metrics:
- Total Lines: 150
- Functions: 10
- Variables: 25

Code formatted successfully.
```

---

## Roadmap
- **Support for Additional Languages**: Extend support beyond Python to include languages such as JavaScript, Go, Java, and more.
- **Advanced Metrics**: Add cyclomatic complexity and dependency analysis.
- **Integration**: Enable seamless integration with popular IDEs and CI/CD pipelines.

---

## Contributing
Contributions are welcome! If you’d like to contribute to CodeAnalyze, please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.

---

## License
This project is licensed under the MIT License. See the LICENSE file for details.

---

## Support
For issues, suggestions, or questions, please create an issue on the [GitHub repository](<repository-url>) or contact the maintainer.

