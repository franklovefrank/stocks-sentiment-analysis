package export

import (
    "encoding/csv"
    "encoding/json"
    "os"
)

func ExportToCSV(data interface{}, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Convert data to CSV format
    // Write data to CSV file

    return nil
}

func ExportToJSON(data interface{}, filename string) error {
    file, err := os.Create(filename)
   
