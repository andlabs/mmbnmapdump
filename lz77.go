// 1 september 2013
package main

import (
	"fmt"
	"io"
	"encoding/binary"
)

const (
	// header
	lzFormatShift = 4
	lzFormatMask = 0xF << lzFormatShift
	lzSizeShift = 8
	lzSizeMask = 0xFFFFFF << lzSizeShift

	// copy block spec
	lzDispRightShift = 0
	lzDispRightMask = 0xF << lzDispRightShift
	lznCopyShift = 4
	lznCopyMask = 0xF << lznCopyShift
	lzDispLeftShift = 8
	lzDispLeftMask = 0xFF << lzDispLeftShift

	// combines the two halves of the disp above
	lzDispFinalShift = 4
)

func LZ77Decomp(r io.Reader) (data []byte, err error) {
	readbyte := func() (b byte, err error) {
		err = binary.Read(r, binary.LittleEndian, &b)
		return b, err
	}
	readword := func() (w uint16, err error) {
		err = binary.Read(r, binary.LittleEndian, &w)
		return w, err
	}
	readlong := func() (l uint32, err error) {
		err = binary.Read(r, binary.LittleEndian, &l)
		return l, err
	}

	header, err := readlong()
	if err != nil {
		return nil, err		// TODO wrap with more text?
	}

	format := (header & lzFormatMask) >> lzFormatShift
	size := (header & lzSizeMask) >> lzSizeShift

	if format != 1 {
		return nil, fmt.Error("error: invalid LZ77 data compression format (expected 1, got 0x%X)", format)
	}

	data = make([]byte, 0, size)

	for n := uint32(0); n < size; n++ {
		var b byte

		b, err = readbyte()
		if err != nil {
			return nil, err		// TODO wrap with more text?
		}
		for bits = 0; bits < 8; bits++ {
			if (b & 0x80) == 0 {	// copy next byte
				c, err := readbyte()
				if err != nil {
					return nil, err		// TODO wrap with more text?
				}
				data = append(data, c)
			} else {			// opy n+3 bytes from data
				copyspec, err := readword()
				if err != nil {
					return nil, err		// TODO wrap with more text?
				}
				dispRight := (copyspec & lzDispRightMask) >> lzDispRightShift
				dispLeft := (copyspec & lzDispLeftMask) >> lzDispLeftShift
				disp := (dispLeft << lzDispFinalShift) | lzDispRight
				nCopy := (copyspec & lznCopyMask) >> lznCopyShift
				for i := uint16(0); i < nCopy + 3; i++ {
					data = append(data, data[len(data) - 1 - disp - 1])
				}
			}
			b >>= 1
		}
	}

	return data, nil
}
