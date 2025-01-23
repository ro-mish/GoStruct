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

	// Close the file after we're done with it (this is async)
	defer jsonlFile.Close()

	// This assumes first line is csv header
	headers := records[0]

	// Go through records and make json map for each record
	for _, record := range records[1:] {
		// Creates map using built-in make
		jsonMap := make(map[string]interface{})


		for i, value := range record {
			jsonMap[headers[i]] = parseValue(value)
		}

		// Marshal the json map to json
		jsonData, err := json.Marshal(jsonMap)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		
		// Write the json data to the jsonl file
		_, err = jsonlFile.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing to JSONL file:", err)
			return
		}

		// Write a newline character to the end of each record
		jsonlFile.WriteString("\n")
	}

	fmt.Printf("CSV data successfully written to %s\n", outputFilePath)
}

func parseValue(value string) interface{} {
	var intValue int
	var floatValue float64
	var err error

	intValue, err = strconv.Atoi(value)
	if err == nil {
		return intValue
	}

	floatValue, err = strconv.ParseFloat(value, 64)
	if err == nil {
		return floatValue
	}

	return value
}
