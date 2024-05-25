# Hierarchical Structure Management Using Composite Pattern in Go

## Overview
This repository demonstrates the application of the Composite design pattern in Go. The project highlights the implementation of composite structures, allowing individual objects and compositions of objects to be treated uniformly. This pattern is ideal for representing part-whole hierarchies in a system.

## Pattern Description
The Composite pattern is used to compose objects into tree structures to represent part-whole hierarchies. It lets clients treat individual objects and compositions of objects uniformly. In this project, we have implemented both Leaf and Composite components to demonstrate this pattern effectively.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and manipulation of the composite structure.
- **pkg/**
  - **composite/**: Houses the component interfaces and their implementations (`Leaf` and `Composite`).

## Component Management
This project features the management of components within a composite structure. It includes functionality for adding, removing, displaying, and counting components, showcasing the flexibility and power of the Composite pattern.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Composite_Pattern.git
cd Go_Composite_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp