# Huffman Compression Tool

## Overview

This project implements a **Huffman coding** compression algorithm in Go. It provides both a command-line interface and an interactive mode for compressing and decompressing files.

## Features

- **Lossless compression** using optimal prefix codes
- **File format** with embedded frequency table (self-contained .huf files)
- **CLI mode** for scripting and automation
- **Interactive mode** for ease of use

## Installation

```bash
go build -o huffman.exe main.go
```

## Usage

### Compress a file

```bash
huffman.exe compress input.txt
```

### Decompress a file

```bash
huffman.exe decompress input.txt.huf
```

### Options

| Flag | Description |
|------|-------------|
| `-o` | Output file path |
| `-f` | Force overwrite |

## Algorithm Details

Huffman coding works by:

1. **Frequency analysis** — Count occurrences of each byte value
2. **Tree construction** — Build a binary tree where frequent bytes get shorter codes
3. **Code generation** — Traverse the tree to assign bit sequences
4. **Encoding** — Replace each byte with its variable-length code

### Example

For the string `"hello world"`:

| Char | Frequency | Code |
|------|-----------|------|
| l    | 3         | 0    |
| o    | 2         | 10   |
| h    | 1         | 110  |
| e    | 1         | 1110 |
| (space) | 1      | 1111 |

## Performance

The ratio depends on the data entropy. Text files typically achieve:

- English text: 40-50% compression
- Source code: 30-40% compression
- Binary data: 0-10% compression (or expansion)

## File Format

The .huf file format is:

```
[MAGIC: 2 bytes]   = 0x4855
[SYMBOLS: 2 bytes] = number of unique symbols
[SYMBOL TABLE]     = (char + frequency) for each symbol
[BITSTREAM]        = Huffman-encoded data
```

## Testing

```bash
go test ./...
```

The test suite covers:

- Roundtrip compression/decompression
- Empty files
- Single character files
- Binary data
- All 256 byte values
- Invalid file detection

## License

This project is for educational purposes.

## Contributing

Contributions are welcome. Please submit pull requests or open issues.

## References

- Huffman, D. A. (1952). "A Method for the Construction of Minimum-Redundancy Codes"
- Wikipedia: Huffman coding
- Go standard library documentation

## Changelog

### v1.0.0

- Initial release
- Basic compression and decompression
- CLI and interactive modes
- Comprehensive test suite

---

*Generated for demonstration purposes*
# Huffman Compression Tool

## Overview

This project implements a **Huffman coding** compression algorithm in Go. It provides both a command-line interface and an interactive mode for compressing and decompressing files.

## Features

- **Lossless compression** using optimal prefix codes
- **File format** with embedded frequency table (self-contained .huf files)
- **CLI mode** for scripting and automation
- **Interactive mode** for ease of use

## Installation

```bash
go build -o huffman.exe main.go
```

## Usage

### Compress a file

```bash
huffman.exe compress input.txt
```

### Decompress a file

```bash
huffman.exe decompress input.txt.huf
```

### Options

| Flag | Description |
|------|-------------|
| `-o` | Output file path |
| `-f` | Force overwrite |

## Algorithm Details

Huffman coding works by:

1. **Frequency analysis** — Count occurrences of each byte value
2. **Tree construction** — Build a binary tree where frequent bytes get shorter codes
3. **Code generation** — Traverse the tree to assign bit sequences
4. **Encoding** — Replace each byte with its variable-length code

### Example

For the string `"hello world"`:

| Char | Frequency | Code |
|------|-----------|------|
| l    | 3         | 0    |
| o    | 2         | 10   |
| h    | 1         | 110  |
| e    | 1         | 1110 |
| (space) | 1      | 1111 |

## Performance

The ratio depends on the data entropy. Text files typically achieve:

- English text: 40-50% compression
- Source code: 30-40% compression
- Binary data: 0-10% compression (or expansion)

## File Format

The .huf file format is:

```
[MAGIC: 2 bytes]   = 0x4855
[SYMBOLS: 2 bytes] = number of unique symbols
[SYMBOL TABLE]     = (char + frequency) for each symbol
[BITSTREAM]        = Huffman-encoded data
```

## Testing

```bash
go test ./...
```

The test suite covers:

- Roundtrip compression/decompression
- Empty files
- Single character files
- Binary data
- All 256 byte values
- Invalid file detection

## License

This project is for educational purposes.

## Contributing

Contributions are welcome. Please submit pull requests or open issues.

## References

- Huffman, D. A. (1952). "A Method for the Construction of Minimum-Redundancy Codes"
- Wikipedia: Huffman coding
- Go standard library documentation

## Changelog

### v1.0.0

- Initial release
- Basic compression and decompression
- CLI and interactive modes
- Comprehensive test suite

---

*Generated for demonstration purposes*
# Huffman Compression Tool

## Overview

This project implements a **Huffman coding** compression algorithm in Go. It provides both a command-line interface and an interactive mode for compressing and decompressing files.

## Features

- **Lossless compression** using optimal prefix codes
- **File format** with embedded frequency table (self-contained .huf files)
- **CLI mode** for scripting and automation
- **Interactive mode** for ease of use

## Installation

```bash
go build -o huffman.exe main.go
```

## Usage

### Compress a file

```bash
huffman.exe compress input.txt
```

### Decompress a file

```bash
huffman.exe decompress input.txt.huf
```

### Options

| Flag | Description |
|------|-------------|
| `-o` | Output file path |
| `-f` | Force overwrite |

## Algorithm Details

Huffman coding works by:

1. **Frequency analysis** — Count occurrences of each byte value
2. **Tree construction** — Build a binary tree where frequent bytes get shorter codes
3. **Code generation** — Traverse the tree to assign bit sequences
4. **Encoding** — Replace each byte with its variable-length code

### Example

For the string `"hello world"`:

| Char | Frequency | Code |
|------|-----------|------|
| l    | 3         | 0    |
| o    | 2         | 10   |
| h    | 1         | 110  |
| e    | 1         | 1110 |
| (space) | 1      | 1111 |

## Performance

The ratio depends on the data entropy. Text files typically achieve:

- English text: 40-50% compression
- Source code: 30-40% compression
- Binary data: 0-10% compression (or expansion)

## File Format

The .huf file format is:

```
[MAGIC: 2 bytes]   = 0x4855
[SYMBOLS: 2 bytes] = number of unique symbols
[SYMBOL TABLE]     = (char + frequency) for each symbol
[BITSTREAM]        = Huffman-encoded data
```

## Testing

```bash
go test ./...
```

The test suite covers:

- Roundtrip compression/decompression
- Empty files
- Single character files
- Binary data
- All 256 byte values
- Invalid file detection

## License

This project is for educational purposes.

## Contributing

Contributions are welcome. Please submit pull requests or open issues.

## References

- Huffman, D. A. (1952). "A Method for the Construction of Minimum-Redundancy Codes"
- Wikipedia: Huffman coding
- Go standard library documentation

## Changelog

### v1.0.0

- Initial release
- Basic compression and decompression
- CLI and interactive modes
- Comprehensive test suite

---

*Generated for demonstration purposes*
# Huffman Compression Tool

## Overview

This project implements a **Huffman coding** compression algorithm in Go. It provides both a command-line interface and an interactive mode for compressing and decompressing files.

## Features

- **Lossless compression** using optimal prefix codes
- **File format** with embedded frequency table (self-contained .huf files)
- **CLI mode** for scripting and automation
- **Interactive mode** for ease of use

## Installation

```bash
go build -o huffman.exe main.go
```

## Usage

### Compress a file

```bash
huffman.exe compress input.txt
```

### Decompress a file

```bash
huffman.exe decompress input.txt.huf
```

### Options

| Flag | Description |
|------|-------------|
| `-o` | Output file path |
| `-f` | Force overwrite |

## Algorithm Details

Huffman coding works by:

1. **Frequency analysis** — Count occurrences of each byte value
2. **Tree construction** — Build a binary tree where frequent bytes get shorter codes
3. **Code generation** — Traverse the tree to assign bit sequences
4. **Encoding** — Replace each byte with its variable-length code

### Example

For the string `"hello world"`:

| Char | Frequency | Code |
|------|-----------|------|
| l    | 3         | 0    |
| o    | 2         | 10   |
| h    | 1         | 110  |
| e    | 1         | 1110 |
| (space) | 1      | 1111 |

## Performance

The ratio depends on the data entropy. Text files typically achieve:

- English text: 40-50% compression
- Source code: 30-40% compression
- Binary data: 0-10% compression (or expansion)

## File Format

The .huf file format is:

```
[MAGIC: 2 bytes]   = 0x4855
[SYMBOLS: 2 bytes] = number of unique symbols
[SYMBOL TABLE]     = (char + frequency) for each symbol
[BITSTREAM]        = Huffman-encoded data
```

## Testing

```bash
go test ./...
```

The test suite covers:

- Roundtrip compression/decompression
- Empty files
- Single character files
- Binary data
- All 256 byte values
- Invalid file detection

## License

This project is for educational purposes.

## Contributing

Contributions are welcome. Please submit pull requests or open issues.

## References

- Huffman, D. A. (1952). "A Method for the Construction of Minimum-Redundancy Codes"
- Wikipedia: Huffman coding
- Go standard library documentation

## Changelog

### v1.0.0

- Initial release
- Basic compression and decompression
- CLI and interactive modes
- Comprehensive test suite

---

*Generated for demonstration purposes*
# Huffman Compression Tool

## Overview

This project implements a **Huffman coding*