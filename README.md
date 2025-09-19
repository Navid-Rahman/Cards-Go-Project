# Cards (Go Learning Project)

This repository contains a simple playing card deck implementation written in Go. It is intentionally small and focused to help you learn core language features:

- Declaring custom types and methods (receiver functions)
- Working with slices (append, slicing, iteration)
- Multiple return values
- String manipulation
- File I/O
- Randomness and shuffling
- Writing tests with the `testing` package

## Project Layout

- `deck.go`: Defines the `deck` type and its helper functions/methods (`newDeck`, `deal`, `toString`, `saveToFile`, `newDeckFromFile`, `shuffle`). All functions include Go doc comments to explain their purpose.
- `main.go`: Small playground to create, shuffle, and print a deck. Contains commented examples for dealing and file operations.
- `deck_test.go`: Tests that verify deck construction and file round-trip behavior.

## Prerequisites

- Go 1.20+ installed (project tested on Go 1.25)

## Getting Started

Initialize the module (already done):

```bash
go mod tidy
```

Run tests:

```bash
go test ./...
```

Run the program:

```bash
go run .
```

You should see a shuffled list of cards printed to stdout.

## Useful Experiments

- In `main.go`, uncomment the sections to:
  - Deal a hand and see the remaining cards
  - Convert a deck to a string
  - Save a deck to a file and read it back

## Notes on Shuffling

The current `shuffle` implementation uses `rand.Intn(len(d) - 1)`, which does not include the last index and can bias the distribution. For learning purposes, this is fine. For a better shuffle, consider seeding and using `rand.Intn(len(d))` or a Fisher–Yates shuffle.

```go
// Example improvement
// import "math/rand"; import "time"
// rand.Seed(time.Now().UnixNano())
// for i := range d {
//     j := rand.Intn(len(d))
//     d[i], d[j] = d[j], d[i]
// }
```

## Cleaning Up Test Artifacts

The test creates and removes a temporary file named `_deckTesting`.

## License

MIT
