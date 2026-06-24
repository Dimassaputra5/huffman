package huffman

import (
	"encoding/binary"
	"os"
)

const Magic = 0x4855

func Compress(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}

	root := BuildTree(freq)
	codes := GenerateCodes(root)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := binary.Write(outFile, binary.BigEndian, uint16(Magic)); err != nil {
		return err
	}
	numSymbols := uint16(len(freq))
	if err := binary.Write(outFile, binary.BigEndian, numSymbols); err != nil {
		return err
	}
	for ch, f := range freq {
		if err := binary.Write(outFile, binary.BigEndian, ch); err != nil {
			return err
		}
		if err := binary.Write(outFile, binary.BigEndian, uint32(f)); err != nil {
			return err
		}
	}

	bw := NewBitWriter(outFile)
	for _, b := range data {
		code := codes[b]
		for i := int(code.Bits - 1); i >= 0; i-- {
			bit := int((code.Value >> uint(i)) & 1)
			if err := bw.WriteBit(bit); err != nil {
				return err
			}
		}
	}
	if err := bw.Flush(); err != nil {
		return err
	}
	return nil
}
