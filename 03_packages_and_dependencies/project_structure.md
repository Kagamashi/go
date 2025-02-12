# Project Structure

Summary:
cmd/ → Executables (binaries).
internal/ → Private packages (restricted to the module).
pkg/ → Public reusable packages (utilities, models).
configs/ → Configuration files.
api/ → API definitions.
tests/ → Integration tests & mocks.

'''csharp
myproject/
│── cmd/          # Command-line binaries (main applications)
│   ├── app1/
│   │   └── main.go
│   ├── app2/
│   │   └── main.go
│── internal/     # Private packages (restricted to this module)
│   ├── database/
│   ├── services/
│── pkg/          # Public reusable packages (can be imported by other projects)
│   ├── utils/
│   ├── models/
│── api/          # API definitions (Protobuf, GraphQL, OpenAPI, etc.)
│── configs/      # Configuration files (YAML, JSON, environment variables)
│── docs/         # Documentation files (README, API docs)
│── tests/        # Integration tests, test data, mocks
│── vendor/       # Vendored dependencies (optional)
│── go.mod        # Module definition
│── go.sum        # Dependency checksums
│── README.md     # Project overview
│── Makefile      # Automation tasks
'''

---

## cmd/: Entry Points for Applications

- cmd/ directory contains entry points (executables) for different applications in the project
- each subdirectory under cmd/ contains a main.go file for that specific app
- keeps binaries separate and avoids polluting the root directory 

'''go
package main

import (
	"fmt"
	"myproject/internal/server"
)

func main() {
	fmt.Println("Starting web server...")
	server.Start()
}
'''

---

## internal/: Private Code (Restricted Access)

- code inside internal/ cannot be imported by external projects
- it's enforced by Go module system
- use internal/ for business logic, database connections, utilities that shouldn't be exposed

'''go
package database

import "fmt"

func Connect() {
	fmt.Println("Connecting to database...")
}

'''

---

## pkg/: Publicly Reusable Code

- pkg/ directory contains reusable, public packages that can be imported by other Go projects
- use for utility functions, models, generic components

'''go
package utils

func Sum(a, b int) int {
	return a + b
}
'''

[!] Some teams avoid pkg/ and instead use individual package directories like models/, utils/

--- 

## api/: API Definitions

Store API related files:
- Protobuf files for gRPC (api/proto/)
- OpenAPI specs (api/swagger.yaml)
- GraphQL schemas (api/schema.graphql)

---

## configs/: Configuration Files

Store environment specific configs like:
- config.yaml
- config.json
- .env files

'''go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port int `json:"port"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	return config, err
}
'''

---

## tests/: Integration Tests and Mocks

Store end-to-end tests, integration tests, and test data separately from unit tests

---
