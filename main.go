package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	println("Flattening JSON...")

	// Read JSON from file data.json at the root of the project
	jsonData, err := readJSON("data.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Prepare a map to hold the flattened data and a slice for ordered keys
	flatMap := make(map[string]interface{})
	orderedKeys := []string{}

	// Flatten the JSON data
	flattenJSON(jsonData, "", flatMap, &orderedKeys)

	// Write the flattened data to a CSV file
	if err := writeToCSV(orderedKeys, flatMap, "flattened_data.csv"); err != nil {
		log.Fatalf("Error writing to CSV: %v", err)
	}
	fmt.Println("Data written to flattened_data.csv")
}

// readJSON reads JSON data from a file and returns it as a map
func readJSON(filename string) (map[string]interface{}, error) {
	// Read the entire JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func flattenJSON(data map[string]interface{}, parentKey string, flatMap map[string]interface{}, orderedKeys *[]string) {
	for key, value := range data {
		newKey := parentKey
		if parentKey != "" {
			newKey += "."
		}

		newKey += key

		switch v := value.(type) {
		// It is a map
		case map[string]interface{}:
			flattenJSON(v, newKey, flatMap, orderedKeys)

		case []interface{}:
			for i, item := range v {
				itemKey := fmt.Sprintf("%s[%d]", newKey, i)
				if itemMap, ok := item.(map[string]interface{}); ok {
					// The [i] item of the map is a JSON, therefore recursively flatten it.
					flattenJSON(itemMap, itemKey, flatMap, orderedKeys)
				} else {
					// The [i] item of the map is a primitive
					flatMap[itemKey] = item
					*orderedKeys = append(*orderedKeys, itemKey) // Add to ordered keys
				}
			}

		default:
			flatMap[newKey] = v
			*orderedKeys = append(*orderedKeys, newKey) // Add to ordered keys
		}
	}
}

func writeToCSV(orderedKeys []string, flatMap map[string]interface{}, filename string) error {
	// Create a new CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	if err := writer.Write(orderedKeys); err != nil {
		return err
	}

	// Write the data row
	dataRow := make([]string, len(orderedKeys))
	for i, key := range orderedKeys {
		value := flatMap[key]
		// Convert value to string for CSV writing
		dataRow[i] = fmt.Sprintf("%v", value)
	}
	if err := writer.Write(dataRow); err != nil {
		return err
	}

	return nil
}
