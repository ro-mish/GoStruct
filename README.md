# GoStruct
GoStruct is a command-line application that converts CSV files to JSON Lines (JSONL) files. 

## Getting Started

### System Requirements and Prerequisites

- **Go**: Ensure you have Go 1.20 or later installed. Find latest version at [golang.org](https://golang.org/dl/).

### Installation

1. **Clone the Repository**

   Open your terminal, and run the following command to clone the repository:

   ```bash
   git clone https://github.com/ro-mish/GoStruct.git
   ```

   Navigate into the project directory:

   ```bash
   cd GoStruct
   ```

2. **Build the Project**

   Run the following command to build the project:

   ```bash
   go build -o GoStruct cmd/GoStruct/main.go
   ```

   This will create an executable named `GoStruct` in your current directory.

### Usage

To run the program, use the following command:
    ```bash
    ./GoStruct <input.csv> <output.jsonl>
    ```

    - `<input.csv>`: Path to the input CSV file.
    - `<output.jsonl>`: Path where the output JSON Lines file will be saved.

    Example:
    ```bash
    ./GoStruct housesInput.csv housesOutput.jsonl
    ```

### Running Tests

To run the tests, execute the following command:

    ```bash
    go test ./tests
    ```

This will run all the unit tests in the `tests` directory.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


### Use of AI Assistants
- I used gpt (o1) to explain to me key concepts in go in relationship to python
    - I learned about how 'marshal' is similar to 'json.dumps()'
    - Go topics like goroutines and how to set them up
    - importing and using built-in packages in Go
- AI was also used in the debugging process in building out the conversion logic
    - Helped me find errors and resolve them when needed
    - Identified best practices when writing the code as well