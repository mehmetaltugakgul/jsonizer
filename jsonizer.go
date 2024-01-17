package jsonizer

import (
	"encoding/json"
	"errors"
)

// JSON is a utility struct for JSON operations.
type JSON struct {
	data interface{}
}

// NewJSON creates a new JSON instance with the provided data.
func NewJSON(data interface{}) *JSON {
	return &JSON{data: data}
}

// Stringify converts the JSON data to a JSON string.
func (j *JSON) Stringify() (string, error) {
	jsonBytes, err := json.Marshal(j.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// ParseJSON creates a new JSON instance from a JSON string.
func ParseJSON(jsonStr string) (*JSON, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, err
	}
	return &JSON{data: data}, nil
}

// Get retrieves a value from the JSON data by key.
func (j *JSON) Get(key string) (*JSON, error) {
	if obj, ok := j.data.(map[string]interface{}); ok {
		if val, exists := obj[key]; exists {
			return &JSON{data: val}, nil
		}
	}
	return nil, errors.New("key not found")
}

// ArrayGet retrieves a value from the JSON data by index if it's an array.
func (j *JSON) ArrayGet(index int) (*JSON, error) {
	if arr, ok := j.data.([]interface{}); ok {
		if index >= 0 && index < len(arr) {
			return &JSON{data: arr[index]}, nil
		}
	}
	return nil, errors.New("index out of range or not an array")

}
