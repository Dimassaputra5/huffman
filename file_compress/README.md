file-compressor/
├── go.mod
├── main.go                 # atau di cmd/compressor/main.go
├── internal/
│   └── huffman/
│       ├── huffman.go      # definisi tree, node, dll.
│       ├── encoder.go      # fungsi kompresi
│       ├── decoder.go      # fungsi dekompresi
│       └── bitio.go        # baca/tulis bit (karena Huffman bekerja bit-level)
└── README.md