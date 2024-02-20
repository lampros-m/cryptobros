package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var (
	JSONPrefix = ""
	JSONIndent = "    "
)

// ReadFile reads a file and returns its content as a byte array
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	return bytes, nil
}

// WriteFileJSON writes to a file in JSON format
func WriteFileJSON(content interface{}, path string) {
	b, _ := json.MarshalIndent(content, JSONPrefix, JSONIndent)
	_ = os.WriteFile(path, b, 0644)
}

// PrettyPrint prints a struct in a pretty format
func PrettyPrintJSON(content interface{}) {
	b, _ := json.MarshalIndent(content, JSONPrefix, JSONIndent)
	fmt.Println(string(b))
}
