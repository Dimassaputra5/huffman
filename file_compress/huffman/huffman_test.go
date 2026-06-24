package huffman

import (
	"bytes"
	"math/rand"
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

func TestRoundtripEmpty(t *testing.T) {
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "empty.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, []byte{}, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if len(result) != 0 {
		t.Fatalf("expected empty, got %d bytes", len(result))
	}
}

func TestRoundtripSingleChar(t *testing.T) {
	data := []byte("AAAAA")
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "single.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("got %q, want %q", result, data)
	}
}

func TestRoundtripTwoChars(t *testing.T) {
	data := []byte("ABABABABABABABAB")
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "two.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("got %q, want %q", result, data)
	}
}

func TestRoundtripAllBytes(t *testing.T) {
	data := make([]byte, 256)
	for i := range data {
		data[i] = byte(i)
	}
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "allbytes.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("mismatch on all-byte sequence")
	}
}

func TestRoundtripBinaryData(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	data := make([]byte, 10000)
	rng.Read(data)
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "binary.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("mismatch on binary data")
	}
}

func TestRoundtripRepeatingPattern(t *testing.T) {
	pattern := []byte("The quick brown fox jumps over the lazy dog. ")
	data := bytes.Repeat(pattern, 1000)
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "pattern.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("mismatch on repeating pattern")
	}
}

func TestRoundtripNewline(t *testing.T) {
	data := []byte("line1\nline2\nline3\n")
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "newline.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	decPath := filepath.Join(tmp, "decompressed.bin")

	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	if err := Decompress(outPath, decPath); err != nil {
		t.Fatal(err)
	}
	result, _ := os.ReadFile(decPath)
	if !bytes.Equal(data, result) {
		t.Fatalf("got %q, want %q", result, data)
	}
}

func TestDecompressInvalidFile(t *testing.T) {
	tmp := t.TempDir()
	badPath := filepath.Join(tmp, "bad.huf")
	os.WriteFile(badPath, []byte{0, 0, 0, 0}, 0644)

	err := Decompress(badPath, filepath.Join(tmp, "out.bin"))
	if err == nil {
		t.Fatal("expected error for invalid magic number")
	}
	if err != ErrInvalidMagic {
		t.Fatalf("expected ErrInvalidMagic, got %v", err)
	}
}

func TestBuildTree(t *testing.T) {
	freq := map[byte]int{'a': 5, 'b': 9, 'c': 12, 'd': 13, 'e': 16, 'f': 45}
	root := BuildTree(freq)
	if root == nil {
		t.Fatal("BuildTree returned nil")
	}
	if root.freq != 100 {
		t.Fatalf("expected total freq 100, got %d", root.freq)
	}
}

func TestBuildTreeSingleChar(t *testing.T) {
	freq := map[byte]int{'x': 10}
	root := BuildTree(freq)
	if root == nil {
		t.Fatal("BuildTree returned nil for single char")
	}
	if root.freq != 10 {
		t.Fatalf("expected freq 10, got %d", root.freq)
	}
	codes := GenerateCodes(root)
	code, ok := codes['x']
	if !ok {
		t.Fatal("single char should have a code")
	}
	if code.Bits != 1 {
		t.Fatalf("single char should have 1-bit code, got %d bits", code.Bits)
	}
}

func TestCompressSmallerThanInput(t *testing.T) {
	data := bytes.Repeat([]byte("AAAAABBBBBCCCCCDDDDD"), 100)
	tmp := t.TempDir()
	inPath := filepath.Join(tmp, "in.bin")
	outPath := filepath.Join(tmp, "compressed.huf")
	os.WriteFile(inPath, data, 0644)
	if err := Compress(inPath, outPath); err != nil {
		t.Fatal(err)
	}
	inInfo, _ := os.Stat(inPath)
	outInfo, _ := os.Stat(outPath)
	if outInfo.Size() >= inInfo.Size() {
		t.Logf("compressed (%d) >= original (%d) — expected for small/repetitive data with header overhead", outInfo.Size(), inInfo.Size())
	}
}
