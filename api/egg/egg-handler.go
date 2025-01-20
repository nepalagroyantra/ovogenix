package egg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// IoTData represents the structure of the incoming data
type IoTData struct {
    IoTID       string  `json:"iot_id"`
    Temperature float64 `json:"temperature"`
    Humidity    float64 `json:"humidity"`
    Gas         float64 `json:"gas"`
    Gyro        string  `json:"gyro"`
    Timestamp   string  `json:"timestamp"`
}

// Handler processes incoming IoT data and stores it in a JSON file
func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var data IoTData

    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Add timestamp to the data
    data.Timestamp = time.Now().Format(time.RFC3339)

    // Define the file path
    filePath := "data/data.json"

    // Read existing data from file or create a new slice
    var records []IoTData

    if _, err := os.Stat(filePath); err == nil {
        file, err := os.Open(filePath)
        if err != nil {
            http.Error(w, "Error reading data file", http.StatusInternalServerError)
            return
        }
        defer file.Close()

        err = json.NewDecoder(file).Decode(&records)
        if err != nil && err.Error() != "EOF" {
            http.Error(w, "Error decoding JSON data", http.StatusInternalServerError)
            return
        }
    }

    // Append new data to existing records
    records = append(records, data)

    // Save updated data back to file
    file, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Error saving data", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(records)
    if err != nil {
        http.Error(w, "Error writing to file", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Data received and stored successfully.")
}
