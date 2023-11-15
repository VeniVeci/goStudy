package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/xitongsys/parquet-go/ParquetFile"
	"github.com/xitongsys/parquet-go/ParquetReader"
)

func main() {
	filePath := "D:\\SanyProject\\starfish\\test\\F1111_110_20230719_WindFarmData.parquet"

	// Open Parquet file
	file, err := ParquetFile.NewLocalFileReader(filePath)
	if err != nil {
		log.Fatal("Error opening Parquet file:", err)
	}
	defer file.Close()

	// Create a new Parquet reader
	parquetReader, err := ParquetReader.NewParquetReader(file, nil, 4)
	if err != nil {
		log.Fatal("Error creating Parquet reader:", err)
	}
	defer parquetReader.ReadStop()

	// Variables to track memory usage
	var records []map[string]interface{}
	var memStats runtime.MemStats

	// Read all records into memory
	for {
		// Update memory statistics
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Current memory usage: %d bytes\n", memStats.Alloc)

		record, err := parquetReader.Read()
		if err != nil {
			break
		}
		records = append(records, record)
	}

	// Print the loaded records
	for i, record := range records {
		fmt.Printf("Record %d: %+v\n", i+1, record)
	}
}
