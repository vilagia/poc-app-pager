# poc-app-pager

## Overview

`poc-app-pager` is a Go-based application designed to benchmark data retrieval operations with varying activity thresholds.

## Features

- **Benchmarking:** Comprehensive benchmarks to evaluate the performance of data queries with different activity thresholds.

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/vilagia/poc-app-pager.git
    ```
2. **Navigate to the project directory:**
    ```bash
    cd poc-app-pager
    ```
3. **Install dependencies:**
    ```bash
    go mod download
    ```

## Usage

### Running Benchmarks

Execute the benchmark tests to evaluate performance:
```bash
go test -bench=.
```

### Project Structure

- `bench_test.go`: Contains benchmark tests for various data retrieval scenarios.
- `repository/foo.go`: Implements the data retrieval logic with adjustable activity thresholds.
- `go.mod` & `go.sum`: Manage project dependencies.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License.
