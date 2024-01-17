# Package JSONizer
JSONizer is a Go package that provides utilities for working with JSON data, including JSON
serialization and deserialization, as well as accessing JSON values by key or index.

# Installation
To use this package, you need to import it into your Go project. You can install it using go get:

```bash
go get github.com/your-username/jsonizer
```

# Usage
# Creating a JSON instance
To create a new JSON instance with your data, you can use the NewJSON function:

```
import "github.com/your-username/jsonizer"
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