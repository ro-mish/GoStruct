package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the input arguments are of the length that we want
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./main <input.csv> <output.jsonl>")
		return
	}

	// First input is input file path (type .csv)
	inputFilePath := os.Args[1]
	// Second input is output file path (type .jsonl)
	outputFilePath := os.Args[2]

	// Open the input file
	// Set csv file 
	csvFile, err := os.Open(inputFilePath)

	// Check if error exists from opening file
	if err != nil {
		fmt.Printf("Error opening CSV file %s: %v\n", inputFilePath, err)
		return
	}

	defer csvFile.Close()

	// Create a new reader for the csv file
	reader := csv.NewReader(csvFile)

	// Read all the records from the csv file
	records, err := reader.ReadAll()

	// Check if error exists from reading file
	if err != nil {
		fmt.Printf("Error reading CSV file %s: %v\n", inputFilePath, err)
		return
	}

	// Create a new file for the output based on args value
	jsonlFile, err := os.Create(outputFilePath)

	// Check if error from creating file
	if err != nil {
		fmt.Printf("Error creating JSONL file %s: %v\n", outputFilePath, err)
		return
	}


}
