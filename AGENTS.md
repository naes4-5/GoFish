# GoFish Agent Guidelines

This document provides essential information for AI agents and developers working on the GoFish project. Adhering to these guidelines ensures consistency, safety, and efficiency across the codebase.

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Version:** 1.25.5 (as specified in `go.mod`)
- **Module Path:** `github.com/naes4-5/GoFish`

## 🚀 Development Workflow

### Build & Run
- **Build the entire project:**
  ```bash
  go build ./...
  ```
- **Run the main application:**
  ```bash
  go run cmd/main.go
  ```

### Testing
- **Run all tests:**
  ```bash
  go test ./...
  ```
- **Run tests with verbose output:**
  ```bash
  go test -v ./...
  ```
- **Run a specific test in a package:**
  ```bash
  go test -v -run <TestName> ./game
  ```
- **Run tests with race detection:**
  ```bash
  go test -race ./...
  ```

### Linting & Verification
- **Run standard Go vet:**
  ```bash
  go vet ./...
  ```
- **Format code:**
  ```bash
  gofmt -s -w .
  ```
- **Tidy dependencies:**
  ```bash
  go mod tidy
  ```

## 📏 Code Style & Conventions

### Formatting
- Always use `gofmt` for indentation (tabs) and spacing.
- Limit line length to approximately 80-100 characters where possible for readability.

### Naming Conventions
- **Exported Identifiers:** Use `PascalCase` (e.g., `MakeDeck`, `Card`).
- **Unexported Identifiers:** Use `camelCase` (e.g., `suitName`, `drawCard`).
- **Package Names:** Use short, lowercase, single-word names (e.g., `game`, `deck`).
- **Receiver Names:** Use short, 1-3 letter abbreviations of the type name (e.g., `d *Deck`, `c *Card`).
- **Interfaces:** Usually end in `-er` (e.g., `Player`, `Dealer`).

### Imports
- Group imports into three sections separated by a newline:
  1. Standard library imports
  2. Third-party library imports
  3. Internal project imports
- Example:
  ```go
  import (
      "errors"
      "fmt"

      "github.com/google/uuid"

      "github.com/naes4-5/GoFish/game"
  )
  ```

### Error Handling
- Errors should be the last return value of a function.
- Check errors immediately after the function call.
- Use `fmt.Errorf` with the `%w` verb to wrap errors for context.
- Avoid using `panic` for expected error conditions; use it only for truly unrecoverable states.
- Error messages should be lowercase and not end in punctuation.

### Structs & Interfaces
- Favor composition over inheritance.
- Use pointer receivers for methods that modify the receiver or for large structs to avoid copying.
- Define interfaces where they are used (consumer-side) rather than where they are implemented.

### Documentation
- Use standard Go doc comments: `// FunctionName ...` for exported functions.
- Comments should explain the *why* rather than the *what* for complex logic.

## 🏗 Project Patterns

### Constructors
- Use the `Make<Type>` pattern for initializing structs that require setup (e.g., `MakeDeck`).
- Prefer returning a pointer to the created struct.

### Slice Operations
- When removing elements from a slice, use the standard `append` trick:
  ```go
  slice = append(slice[:i], slice[i+1:]...)
  ```
- If the order doesn't matter, consider swapping with the last element and truncating for $O(1)$ removal.

### Game Logic
- Keep the `game` package focused on pure logic.
- Input/Output and CLI interactions should reside in the `cmd/` directory.

## 🧪 Testing Strategy
- Place tests in the same package as the code they test, using the `_test.go` suffix.
- Use table-driven tests for multiple test cases of the same function.
- Example:
  ```go
  func TestDrawCard(t *testing.T) {
      tests := []struct {
          name    string
          deck    *Deck
          wantErr bool
      }{
          {"empty deck", &Deck{Cards: []Card{}}, true},
          {"populated deck", MakeDeck(), false},
      }
      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              _, err := tt.deck.DrawCard()
              if (err != nil) != tt.wantErr {
                  t.Errorf("DrawCard() error = %v, wantErr %v", err, tt.wantErr)
              }
          })
      }
  }
  ```

## 🔒 Safety & Security
- Never hardcode secrets or API keys.
- Be cautious with `math/rand/v2` usage; ensure seeds are appropriately handled if determinism is required (though `rand/v2` handles this better than `v1`).
- Avoid `unsafe` package usage unless strictly necessary for performance and thoroughly documented.
- Validate all external inputs, especially if they influence slice indices or memory allocation.

## 🤖 Agent Instructions
- When adding new features, first check for existing tests that might be affected.
- If no tests exist for a module, proactively create them before making significant changes.
- Ensure `go mod tidy` is run after adding or removing dependencies.
- Adhere to the `gofmt` standard strictly.
