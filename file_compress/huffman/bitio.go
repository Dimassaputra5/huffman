package huffman

import (
	"errors"
	"io"
)

type BitReader struct {
	r        io.Reader
	buffer   byte
	bitsLeft uint8
	err      error
}

func NewBitReader(r io.Reader) *BitReader {
	return &BitReader{r: r}
}

func (br *BitReader) readByte() error {
	if br.err != nil {
		return br.err
	}
	var buf [1]byte
	n, err := br.r.Read(buf[:])
	br.buffer = buf[0]

	if err != nil {
		br.err = err
		return err
	}
	if n == 0 {
		br.err = io.EOF
		return io.EOF
	}
	br.bitsLeft = 8
	return nil
}

func (br *BitReader) ReadBit() (int, error) {
	if br.err != nil {
		return 0, br.err
	}
	if br.bitsLeft == 0 {
		if err := br.readByte(); err != nil {
			return 0, err
		}
	}
	bit := (br.buffer >> (br.bitsLeft - 1)) & 1
	br.bitsLeft--
	return int(bit), nil
}

func (br *BitReader) Readbits(n int) (uint64, error) {
	if n > 64 {
		return 0, errors.New("cannot read more than 64 bits at once")
	}
	var val uint64
	for i := 0; i < n; i++ {
		bit, err := br.ReadBit()
		if err != nil {
			return val, err
		}
		val = (val << 1) | uint64(bit)
	}
	return val, nil
}

type BitWriter struct {
	w        io.Writer
	buffer   byte
	bitsLeft uint8
	err      error
}

func NewBitWriter(w io.Writer) *BitWriter {
	return &BitWriter{w: w, bitsLeft: 8}
}

func (bw *BitWriter) WriteBit(bit int) error {
	if bw.err != nil {
		return bw.err
	}
	if bit != 0 && bit != 1 {
		return errors.New("bit must be 0 or 1")
	}
	bw.bitsLeft--
	bw.buffer |= byte(bit) << bw.bitsLeft
	if bw.bitsLeft == 0 {
		_, err := bw.w.Write([]byte{bw.buffer})
		if err != nil {
			bw.err = err
			return err
		}
		bw.buffer = 0
		bw.bitsLeft = 8
	}
	return nil
}

func (bw *BitWriter) Flush() error {
	if bw.err != nil {
		return bw.err
	}
	if bw.bitsLeft < 8 {
		_, err := bw.w.Write([]byte{bw.buffer})
		if err != nil {
			bw.err = err
			return err
		}
		bw.buffer = 0
		bw.bitsLeft = 0
	}
	return nil
}
