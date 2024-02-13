# README

## Description

Prototype application to display hierarchical structure as HTML page.

This application is treated as a learning project to get familiar with the Go language and its libraries.

The model is hard coded into the application with the structure of Group > Service > Component.
The data is stored in YAML format and can be dynamically changed. 
The created HTML page uses Go templates for easy rendering and can also be dynamically changed.

## Local Development in Go

### Running the Application Locally

To run the application, use the following command:

    go run .

Then, open the following URL in a browser: 

    http://localhost:8080


### Updating Dependencies

To update the dependencies, use the following command:

    go mod tidy

### Formatting the Source Code

To format the source code, use the following command:

    go fmt ./...

## Testing

### How to Run Tests

To run the tests, use the following command:

    go test ./...

### Running Tests with Coverage

To run the tests with coverage, use the following command:

    go test -coverpkg=./... ./...
