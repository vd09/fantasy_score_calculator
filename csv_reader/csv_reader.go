package csv_reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// RecordHandler is a type for functions that handle a single CSV record and a map of column indices.
type RecordHandler func(record []string, columnIndex map[string]int) error

// ReadCSV reads a CSV file and processes each record using the provided handler.
func ReadCSV(filePath string, handler RecordHandler) error {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Create a map of column indices
	columnIndex := make(map[string]int)
	for i, header := range headers {
		columnIndex[header] = i
	}

	// Read and process each record
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}

		// Pass the record and column index map to the handler function
		if err := handler(record, columnIndex); err != nil {
			return fmt.Errorf("handler error: %w", err)
		}
	}

	return nil
}

// Example handler function that simply prints the record and uses the column index map
func printRecord(record []string, columnIndex map[string]int) error {
	fmt.Println("Record:", record)
	for columnName, index := range columnIndex {
		fmt.Printf("%s: %s\n", columnName, record[index])
	}
	return nil
}
