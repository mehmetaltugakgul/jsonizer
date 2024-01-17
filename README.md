# Package JSONizer
JSONizer is a Go package that provides utilities for working with JSON data, including JSON
serialization and deserialization, as well as accessing JSON values by key or index.

# Installation
To use this package, you need to import it into your Go project. You can install it using go get:

```bash
go get github.com/mehmetaltugakgul/jsonizer
```

# Usage
# Creating a JSON instance
To create a new JSON instance with your data, you can use the NewJSON function:

```
import "github.com/mehmetaltugakgul/jsonizer"

data := map[string]interface{}{
 "name": "John",
 "age": 30,
}

jsonData := jsonizer.NewJSON(data)

```
# Parsing JSON
To create a JSON instance from a JSON string, you can use the ParseJSON function:

```
jsonStr := `{"name": "Alice", "age": 25}`

parsedJSON, err := jsonizer.ParseJSON(jsonStr)

if err != nil {
 // Handle the error
}

```

## Accessing JSON values
You can retrieve values from the JSON data by key using the Get method:

```
nameJSON, err := jsonData.Get("name")

if err != nil {
 // Handle the error
}

fmt.Println(nameJSON.Stringify())
```

If your JSON data is an array, you can also retrieve values by index using the ArrayGet method:

```
arrayJSON := jsonizer.NewJSON([]interface{}{10, 20, 30})

valueJSON, err := arrayJSON.ArrayGet(1)

if err != nil {
 // Handle the error
}

fmt.Println(valueJSON.Stringify())
```

## Convert ToXML

```
jsonData := &JSON{data: yourJSONObject}

xmlStr, err := jsonData.ToXML()

if err != nil {
    // Handle error
}

fmt.Println(xmlStr)
```

## Convert FromXML

```
xmlStr := "<root><key>value</key></root>"

jsonData, err := FromXML(xmlStr)

if err != nil {
    // Handle error
}

fmt.Println(jsonData)
```

## Convert ToYAML

```
jsonObj := &JSON{data: yourJSONObject}

yamlStr, err := jsonObj.ToYAML()

if err != nil {
    // Handle error
}

fmt.Println(yamlStr)
```
## Convert FromYAML

```
yamlStr := "key: value"

jsonData, err := FromYAML(yamlStr)

if err != nil {
    // Handle error
}

fmt.Println(jsonData)
```

## Example Usage

```
package main

import (
	"fmt"

	"github.com/mehmetaltugakgul/jsonizer"
)

func main() {
	// Create a sample JSON object
	data := map[string]interface{}{
		"name":   "John",
		"age":    30,
		"city":   "New York",
		"status": true,
	}

	jsonObj := jsonizer.NewJSON(data)

	// Stringify the JSON object
	jsonStr, err := jsonObj.Stringify()
	if err != nil {
		fmt.Println("Stringify error:", err)
		return
	}
	fmt.Println("Stringified JSON:", jsonStr)

	// Parse a JSON string
	parsedJSON, err := jsonizer.ParseJSON(jsonStr)
	if err != nil {
		fmt.Println("ParseJSON error:", err)
		return
	}

	// Get a value from the parsed JSON
	name, err := parsedJSON.Get("name")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	fmt.Println("Name:", *name)

	// Try to get a non-existent key
	nonExistent, err := parsedJSON.Get("non_existent_key")

	fmt.Println("nonExistent:", nonExistent)
	if err != nil {
		fmt.Println("Get error for non-existent key:", err)
	}

	// Accessing array elements
	arrayData := []interface{}{1, 2, 3, 4, 5}
	arrayJSON := jsonizer.NewJSON(arrayData)

	// Get an element by index
	index := 2
	element, err := arrayJSON.ArrayGet(index)
	if err != nil {
		fmt.Println("ArrayGet error:", err)
	} else {
		fmt.Printf("Element at index %d: %v\n", index, element)
	}

	// Try to access an out-of-range index
	invalidIndex := 10
	invalidElement, err := arrayJSON.ArrayGet(invalidIndex)
	if err != nil {
		fmt.Println("ArrayGet error for invalid index:", err)
	} else {
		fmt.Println("invalidElement:", invalidElement)
	}

	// Try to access an index on a non-array JSON object
	nonArrayJSON := jsonizer.NewJSON(data)

	nonArrayElement, err := nonArrayJSON.ArrayGet(0)
	if err != nil {
		fmt.Println("ArrayGet error for non-array object:", err)
	} else {
		fmt.Println("nonArrayElement:", nonArrayElement)
	}

	// Try to parse an invalid JSON string
	invalidJSONStr := "invalid_json"
	_, err = jsonizer.ParseJSON(invalidJSONStr)
	if err != nil {
		fmt.Println("ParseJSON error for invalid JSON string:", err)
	}
	
		// Sample JSON data
	jsonData := `{"isbn": "123-456-222",  
	 "author": 
	   {
		 "lastname": "Doe",
		 "firstname": "Jane"
	   },
         "editor": 
	   {
		 "lastname": "Smith",
		 "firstname": "Jane"
	   },
	 "title": "The Ultimate Database Study Guide",  
	 "category": ["Non-Fiction", "Technology"]
	}`

	// Converting JSON data
	jsonObj2, err := jsonizer.ParseJSON(jsonData)
	if err != nil {
		log.Fatalf("JSON parse error: %v", err)
	}

	// Convert JSON data to YAML
	yamlStr, err := jsonObj.ToYAML()
	if err != nil {
		log.Fatalf("YAML conversion error: %v", err)
	}

	fmt.Println("JSON data:")
	fmt.Println(jsonData)

	fmt.Println("\nJSON to YAML conversion:")
	fmt.Println(yamlStr)

	// Converting YAML data to JSON
	jsonObjFromYAML, err := jsonizer.FromYAML(yamlStr)
	if err != nil {
		log.Fatalf("JSON conversion error: %v", err)
	}

	// Checking the result of converting JSON data
	fmt.Println("\nYAML to JSON conversion:")
	if jsonString, err := jsonObjFromYAML.Stringify(); err == nil {
		fmt.Println(jsonString)
	} else {
		log.Fatalf("JSON conversion error: %v", err)
	}
}

```

## License

This package is released under the MIT License. Feel free to use, modify, and distribute it as needed.
