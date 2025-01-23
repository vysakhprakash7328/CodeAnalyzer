package main

import (
	_ "embed"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"path/filepath"
	"os/exec"
	"regexp"
	"bufio"
	"github.com/fsnotify/fsnotify"
	"strings"

)

var installScript string

func installDependencies() error {
	
	// Write the Python script to a temporary file
	tmpFile, err := os.CreateTemp("", "install_requirements_*.py")
	if err != nil {
		return fmt.Errorf("unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file after execution

	// Write the embedded Python script into the temporary file
	_, err = tmpFile.WriteString(installScript)
	if err != nil {
		return fmt.Errorf("unable to write to temporary file: %v", err)
	}
	
	// Run the Python script to install required dependencies
	cmd := exec.Command("python", tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

const logo = `
	██████╗ ██████╗ ██████╗ ███████╗     ██████╗ ██╗    ██╗    ██╗   ██╗██████╗ ██╗	██╗
	██╔════╝██╔═══██╗██╔══██ ██╔════╝    ██╔═══██╗╚██╗ ██╔╝    ╚██╗ ██╔╝██╔══██╗██║██╔╝
	██║     ██║   ██║██	 ██ ██╗      	 ██║██║ ╚	████╔╝      ╚████╔╝ ██████╔╝████╔╝
	██║     ██║   ██║██╔══██ ██╔══╝      ██║   ██║  ╚██╔╝        ╚██╔╝  ██╔═══╝ ██║██╗
	╚██████╗╚██████╔╝██████ ╔╝██████╗    ╚██████╔╝   ██║          ██║   ██║     ██║ ██╗
	╚═════╝ ╚═════╝ ╚═╝		 ╚══════╝     ╚═════╝    ╚═╝          ╚═╝   ╚═╝     ╚═╝ ╚═╝
																				
				Crafted with ❤️ by VPK 🚀
	`

func displayLogo() {
	fmt.Println(logo)
}

func main() {

	
	
	
	var rootCmd = &cobra.Command{
		Use: "codeanalyze",
		Short: "Analyze the quality of your code",
		Long: "This is a CLI tool to analyze code quality, provide metrics, and suggest improvements.",
		
	}

	
	var analyzeCmd = &cobra.Command{
		Use: "analyze",
		Short: "Analyze the specified code file",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			metrics, _ := cmd.Flags().GetBool("metrics")
			suggestions, _ := cmd.Flags().GetBool("suggestions")
			format, _ := cmd.Flags().GetBool("format")
			continuous, _ := cmd.Flags().GetBool("watch")


			displayLogo()

			if continuous {
				watchFile(file, metrics, suggestions)
			} else {
				analyzeFile(file, metrics, suggestions, format)
			}
		},
	}

	analyzeCmd.Flags().Bool("metrics", false, "Display code quality metrics")
	analyzeCmd.Flags().Bool("suggestions", false, "Provide suggestions for improvements")
	analyzeCmd.Flags().Bool("watch", false, "Continuously monitor the file for changes")
	analyzeCmd.Flags().Bool("format", false, "format the code using autopep8")


	rootCmd.AddCommand(analyzeCmd)

	
	
	

	// Continue with your CLI functionality
	fmt.Println("Dependencies installed successfully.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func analyzeFile(file string, metrics bool, suggestions bool, format bool) {
	// Analyze the code file and provide suggestions for improvements
	fmt.Printf("Analyzing file: %s\n", file)

	ext := filepath.Ext(file)

	

	if metrics {
		if ext == ".py" {

		lines, functions, variables := analyzePythonFile(file)

		fmt.Println("Code quality metrics:")
		fmt.Printf("Lines of code: %d\n", lines)
		fmt.Printf("Number of functions: %d\n", functions)
		fmt.Printf("Number of variables: %d\n", variables)
		}
	}
	
	if suggestions {
		if ext == ".py" {
			suggestSyntax(file)
			suggestRadonIssues(file)
			suggestFlake8Issues(file)
		}	
	}

	if format {
		if ext == ".py" {
			formatPEP8(file)
		}
	}
}


func watchFile(file string, metrics bool, suggestions bool) {
	fmt.Printf("Watching file: %s for changes...\n", file)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating file watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(file)
	if err != nil {
		fmt.Println("Error adding file to watcher:", err)
		return
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File modified, re-analyzing...")
					analyzeFile(file, metrics, suggestions, false)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	// Keep the program running until interrupted
	<-done
}




func analyzePythonFile(file string) (int, int, int) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0, 0
	}
	defer f.Close()

	lines := 0
	functions := 0
	variables := 0

	scanner := bufio.NewScanner(f)
	functionPattern := regexp.MustCompile(`^\s*def\s+\w+\s?\(`) // Regex to match function definitions
	variablePattern := regexp.MustCompile(`^\s*(\w+)\s*=`) // Regex to match variable assignments

	for scanner.Scan() {
		line := scanner.Text()
		lines++

		// Check if it's a function definition
		if functionPattern.MatchString(line) {
			functions++
		}

		// Check if it's a variable assignment
		if variablePattern.MatchString(line) {
			variables++
		}
	}

	// Handle scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lines, functions, variables
}

func suggestSyntax(file string) {
	fmt.Println("Suggestions for syntax improvements:")
	cmd := exec.Command("python", "-m", "py_compile", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Fix the following syntax errors:")
		fmt.Println(string(output)) // Print the exact error lines
	} else {
		fmt.Println("No syntax issues detected.")
	}
}

// Suggest improvements based on Flake8 results
func suggestFlake8Issues(file string) {
	if _, err := exec.LookPath("flake8"); err != nil {
		fmt.Println("radon is not installed, installing...")
		
		// Command to install radon
		cmd := exec.Command("python", "-m", "pip", "install", "flake8")
		output, err := cmd.CombinedOutput()
		
		if err != nil {
			fmt.Printf("Failed to install radon: %v\n", err)
			return
		}
		
		fmt.Println("flake8 installed successfully!")
		fmt.Printf("Command output:\n%s\n", output)
	} else {
		fmt.Println("flake8 is already installed.")
	}
	cmd := exec.Command("flake8", file)
	output, err := cmd.CombinedOutput()
	fmt.Print(err)
	if err != nil {
		// Provide specific suggestions based on the error
		fmt.Println("Style issues detected:")
		fmt.Println(string(output)) // You can parse this output to give more tailored suggestions
		// Example:
		// "Refactor long lines to be within 79 characters."
	} else {
		fmt.Println("No style issues detected by Flake8.")
	}
}

// Suggest improvements based on Radon complexity
func suggestRadonIssues(file string) {
	if _, err := exec.LookPath("radon"); err != nil {
		fmt.Println("radon is not installed, installing...")
		
		// Command to install radon
		cmd := exec.Command("python", "-m", "pip", "install", "radon")
		output, err := cmd.CombinedOutput()
		
		if err != nil {
			fmt.Printf("Failed to install radon: %v\n", err)
			return
		}
		
		fmt.Println("radon installed successfully!")
		fmt.Printf("Command output:\n%s\n", output)
	} else {
		fmt.Println("radon is already installed.")
	}
    fmt.Println("Suggestions for complexity improvements:")
    cmd := exec.Command("radon", "cc", file, "-a")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error running Radon complexity check:")
        fmt.Println(string(output))
        return
    }

    // Display the raw output
    fmt.Println("Raw complexity analysis result:")
    fmt.Println(string(output))

    // Parse the output and provide suggestions
    fmt.Println("\nDetailed complexity suggestions:")
    scanner := bufio.NewScanner(strings.NewReader(string(output)))
    complexityPattern := regexp.MustCompile(`\s+(\w+)\s+-\s+(\w)`) // Match function name and grade

    for scanner.Scan() {
        line := scanner.Text()
        if match := complexityPattern.FindStringSubmatch(line); match != nil {
            functionName := match[1]
            grade := match[2]
            fmt.Printf("Function: %s, Complexity Grade: %s\n", functionName, grade)

            // Provide suggestions based on the grade
            switch grade {
            case "A":
                fmt.Println(" - No action needed. Code is simple and clean.")
            case "B":
                fmt.Println(" - Slightly complex. Consider minor refactoring.")
            case "C":
                fmt.Println(" - Moderately complex. Refactor if possible.")
            case "D":
                fmt.Println(" - Complex. Refactor the code to reduce complexity.")
            case "E":
                fmt.Println(" - Very complex. Refactoring is strongly recommended.")
            case "F":
                fmt.Println(" - Unacceptable complexity. Refactor urgently.")
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error parsing Radon output:", err)
    }
}


// Function to automatically format Python file according to PEP8
func formatPEP8(file string) {
	if _, err := exec.LookPath("autopep8"); err != nil {
		fmt.Println("radon is not installed, installing...")
		
		// Command to install radon
		cmd := exec.Command("python", "-m", "pip", "install", "autopep8")
		output, err := cmd.CombinedOutput()
		
		if err != nil {
			fmt.Printf("Failed to install radon: %v\n", err)
			return
		}
		
		fmt.Println("pep8 installed successfully!")
		fmt.Printf("Command output:\n%s\n", output)
	} else {
		fmt.Println("pep8 is already installed.")
	}

	fmt.Println("Formatting Python code according to PEP8...")

	// Run autopep8 to format the code
	cmd := exec.Command("autopep8", "--in-place", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error formatting the code with autopep8:")
		fmt.Println(string(output))
	} else {
		fmt.Println("Code formatted successfully according to PEP8.")
	}
}