package log

import (
	"encoding/json"
	"fmt"
)

// JSON returns v marshalled into JSON. If an error occurs, an error message is returned instead.
// It can be used to quickly add a JSON object into a log message.
func JSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("<JSON Error: %v>", err)
	}
	return string(data)
}

// JSONIndent returns v marshalled into JSON, with indentation. If an error occurs, an error message is returned instead.
// It can be used to quickly add a JSON object into a log message.
func JSONIndent(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("<JSON Error: %v>", err)
	}
	return string(data)
}
