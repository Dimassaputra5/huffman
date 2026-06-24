package huffman

import (
	"encoding/binary"
	"errors"
	"os"
)

func Decompress(inputPath, outputPath string) error {
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var magic uint16
	if err := binary.Read(inFile, binary.BigEndian, &magic); err != nil {
		return err
	}
	if magic != Magic {
		return ErrInvalidMagic
	}
	var numSymbols uint16
	if err := binary.Read(inFile, binary.BigEndian, &numSymbols); err != nil {
		return err
	}

	freq := make(map[byte]int)
	for i := 0; i < int(numSymbols); i++ {
		var ch byte
		var f uint32
		if err := binary.Read(inFile, binary.BigEndian, &ch); err != nil {
			return err
		}
		if err := binary.Read(inFile, binary.BigEndian, &f); err != nil {
			return err
		}
		freq[ch] = int(f)
	}

	root := BuildTree(freq)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	br := NewBitReader(inFile)
	totalBytes := 0
	for _, f := range freq {
		totalBytes += f
	}

	for i := 0; i < totalBytes; i++ {
		node := root
		for node.Left != nil || node.Right != nil {
			bit, err := br.ReadBit()
			if err != nil {
				return err
			}
			if bit == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		if _, err := outFile.Write([]byte{node.Char}); err != nil {
			return err
		}
	}
	return nil
}

var ErrInvalidMagic = errors.New("invalid magic number")
