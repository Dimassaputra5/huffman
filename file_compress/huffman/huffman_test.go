package huffman

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRoundtrip(t *testing.T) {
	original := []byte("Hello, Huffman! This is a test of the compression algorithm.")
	tmp := t.TempDir()

	inPath := filepath.Join(tmp, "input.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	if err := os.WriteFile(inPath, original, 0644); err != nil {
		t.Fatal(err)
	}
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, err := os.ReadFile(decPath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(original, result) {
		t.Fatalf("mismatch:\n  original (%d): % x\n  decoded (%d): % x",
			len(original), original, len(result), result)
	}
}
